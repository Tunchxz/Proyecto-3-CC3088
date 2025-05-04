-- =============================
--  EJERCICIO 5 (triggers)
-- =============================

-- ----------------------------------------------------------------------------
--  TRIGGER 1: ACTUALIZAR ESTADO DEL VEHÍCULO
-- ----------------------------------------------------------------------------
/*
Este trigger actualiza automáticamente el estado de un vehículo cuando cambia el estado de una reserva o 
un contrato de alquiler. Por ejemplo, cuando se confirma una reserva, el estado del vehículo cambia a 
"Reservado". Cuando se activa un contrato de alquiler, el estado del vehículo cambia a "Alquilado". Cuando 
se completa un contrato de alquiler, el estado del vehículo cambia a "Disponible".
*/

CREATE OR REPLACE FUNCTION update_vehicle_status()
RETURNS TRIGGER AS $$
DECLARE
    status_id_available INT;
    status_id_reserved INT;
    status_id_rented INT;
    reservation_status VARCHAR;
    contract_status VARCHAR;
BEGIN
    -- Obtener identificadores de estado
    SELECT id INTO status_id_available FROM OperationStatus WHERE status_name = 'Available';
    SELECT id INTO status_id_reserved FROM OperationStatus WHERE status_name = 'Reserved';
    SELECT id INTO status_id_rented FROM OperationStatus WHERE status_name = 'Rented';
    
    IF TG_TABLE_NAME = 'reservation' THEN
        -- Cuando cambia una reserva
        SELECT status_name INTO reservation_status 
        FROM OperationStatus 
        WHERE id = NEW.status_id;
        
        IF reservation_status = 'Confirmed' THEN
            -- Actualizar vehículo a Reservado
            UPDATE Vehicle SET status_id = status_id_reserved
            WHERE id = NEW.vehicle_id AND status_id = status_id_available;
        ELSIF reservation_status = 'Cancelled' OR reservation_status = 'Completed' THEN
            -- Verificar si no hay un contrato de alquiler activo para este vehículo antes de marcar como Disponible
            IF NOT EXISTS (
                SELECT 1 FROM RentalContract rc 
                JOIN OperationStatus s ON rc.status_id = s.id 
                WHERE rc.reservation_id = NEW.id AND s.status_name = 'Active'
            ) THEN
                UPDATE Vehicle SET status_id = status_id_available
                WHERE id = NEW.vehicle_id AND status_id = status_id_reserved;
            END IF;
        END IF;
    ELSIF TG_TABLE_NAME = 'rentalcontract' THEN
        -- Cuando un contrato cambia
        SELECT status_name INTO contract_status 
        FROM OperationStatus 
        WHERE id = NEW.status_id;
        
        IF contract_status = 'Active' THEN
            -- Establecer vehículo como Alquilado
            UPDATE Vehicle SET status_id = status_id_rented
            WHERE id = (SELECT vehicle_id FROM Reservation WHERE id = NEW.reservation_id);
        ELSIF contract_status = 'Completed' OR contract_status = 'Cancelled' THEN
            -- Volver a poner el vehículo en Disponible
            UPDATE Vehicle SET status_id = status_id_available
            WHERE id = (SELECT vehicle_id FROM Reservation WHERE id = NEW.reservation_id);
        END IF;
    END IF;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Crear Trigger AFTER INSERT/UPDATE
CREATE TRIGGER after_reservation_status_change
AFTER INSERT OR UPDATE OF status_id ON Reservation
FOR EACH ROW
EXECUTE FUNCTION update_vehicle_status();

-- Crear Trigger AFTER INSERT/UPDATE
CREATE TRIGGER after_contract_status_change
AFTER INSERT OR UPDATE OF status_id ON RentalContract
FOR EACH ROW
EXECUTE FUNCTION update_vehicle_status();


-- ----------------------------------------------------------------------------
--  TRIGGER 2: CÁLCULO AUTOMÁTICO DEL IMPORTE DE PAGO
-- ----------------------------------------------------------------------------
/*
Este trigger calcula automáticamente el importe del pago basado en la duración del alquiler y las tarifas.
Cuando se crea un pago para un contrato de alquiler, calcula el importe basado en:
- La duración del alquiler (en días)
- La tarifa aplicable (diaria, semanal o mensual)
- Cualquier multa aplicable
*/

CREATE OR REPLACE FUNCTION calculate_payment_amount()
RETURNS TRIGGER AS $$
DECLARE
    v_start_date DATE;
    v_end_date DATE;
    v_days INTEGER;
    v_daily_rate DECIMAL(10,2);
    v_weekly_rate DECIMAL(10,2);
    v_monthly_rate DECIMAL(10,2);
    v_vehicle_id INTEGER;
    v_total_amount DECIMAL(10,2) := 0;
    v_fine_amount DECIMAL(10,2) := 0;
BEGIN
    -- Si es un pago de contrato
    IF NEW.rental_contract_id IS NOT NULL THEN
        -- Obtener fechas del contrato
        SELECT rc.start_date, rc.end_date, r.vehicle_id
        INTO v_start_date, v_end_date, v_vehicle_id
        FROM RentalContract rc
        JOIN Reservation r ON rc.reservation_id = r.id
        WHERE rc.id = NEW.rental_contract_id;
        
        -- Calcular duración del alquiler en días
        v_days := v_end_date - v_start_date + 1;
        
        -- Obtener tarifas aplicables
        SELECT daily_rate, weekly_rate, monthly_rate
        INTO v_daily_rate, v_weekly_rate, v_monthly_rate
        FROM Rates r
        JOIN Vehicle v ON r.id = v.rates_id
        WHERE v.id = v_vehicle_id;
        
        -- Calcular importe de alquiler basado en precio óptimo
        IF v_days >= 30 THEN
            -- Usar tarifa mensual para meses completos y tarifa diaria para días restantes
            v_total_amount := (v_days / 30) * v_monthly_rate + 
                             (v_days % 30) * v_daily_rate;
        ELSIF v_days >= 7 THEN
            -- Usar tarifa semanal para semanas completas y tarifa diaria para días restantes
            v_total_amount := (v_days / 7) * v_weekly_rate + 
                             (v_days % 7) * v_daily_rate;
        ELSE
            -- Usar tarifa diaria
            v_total_amount := v_days * v_daily_rate;
        END IF;
        
        -- Añadir multas asociadas
        SELECT COALESCE(SUM(amount), 0)
        INTO v_fine_amount
        FROM Fine
        WHERE rental_contract_id = NEW.rental_contract_id;
        
        v_total_amount := v_total_amount + v_fine_amount;
        
        -- Actualizar el importe del pago
        NEW.amount := v_total_amount;
    
    -- Si es un pago de multa, establecer importe igual al de la multa
    ELSIF NEW.fine_id IS NOT NULL THEN
        SELECT amount INTO v_fine_amount FROM Fine WHERE id = NEW.fine_id;
        NEW.amount := v_fine_amount;
    END IF;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Crear Trigger BEFORE INSERT
CREATE TRIGGER before_payment_insert
BEFORE INSERT ON Payment
FOR EACH ROW
EXECUTE FUNCTION calculate_payment_amount();


-- ----------------------------------------------------------------------------
--  TRIGGER 3: SEGUIMIENTO DE MANTENIMIENTO DE VEHÍCULOS
-- ----------------------------------------------------------------------------
/*
Este trigger crea automáticamente registros de mantenimiento cuando el kilometraje del vehículo alcanza
ciertos umbrales. Por ejemplo, cambios de aceite cada 5.000 millas, servicio mayor cada 30.000 millas.
*/

CREATE OR REPLACE FUNCTION check_vehicle_maintenance()
RETURNS TRIGGER AS $$
DECLARE
    last_oil_change_mileage INTEGER := 0;
    last_major_service_mileage INTEGER := 0;
    maintenance_id INTEGER;
    status_id_maintenance INTEGER;
BEGIN
    -- Solo verificar en actualizaciones de kilometraje
    IF NEW.mileage <= OLD.mileage THEN
        RETURN NEW;
    END IF;
    
    -- Obtener el ID del estado de mantenimiento
    SELECT id INTO status_id_maintenance FROM OperationStatus WHERE status_name = 'Under Maintenance';
    
    -- Obtener el kilometraje del último cambio de aceite
    SELECT COALESCE(MAX(v.mileage), 0)
    INTO last_oil_change_mileage
    FROM Vehicle_Maintenance vm
    JOIN Maintenance m ON vm.maintenance_id = m.id
    JOIN Vehicle v ON vm.vehicle_id = v.id
    WHERE vm.vehicle_id = NEW.id
    AND m.description LIKE '%Oil Change%';
    
    -- Obtener el kilometraje del último servicio mayor
    SELECT COALESCE(MAX(v.mileage), 0)
    INTO last_major_service_mileage
    FROM Vehicle_Maintenance vm
    JOIN Maintenance m ON vm.maintenance_id = m.id
    JOIN Vehicle v ON vm.vehicle_id = v.id
    WHERE vm.vehicle_id = NEW.id
    AND m.description LIKE '%Major Service%';
    
    -- Comprobar si toca cambio de aceite (cada 5.000 millas)
    IF (NEW.mileage - last_oil_change_mileage) >= 5000 THEN
        -- Crear registro de mantenimiento
        INSERT INTO Maintenance (maintenance_date, description, cost)
        VALUES (CURRENT_DATE, 'Oil Change due at ' || NEW.mileage || ' miles', 50.00)
        RETURNING id INTO maintenance_id;
        
        -- Vincular al vehículo
        INSERT INTO Vehicle_Maintenance (vehicle_id, maintenance_id)
        VALUES (NEW.id, maintenance_id);
        
        -- Actualizar estado del vehículo si es necesario
        IF NEW.status_id <> status_id_maintenance THEN
            NEW.status_id := status_id_maintenance;
        END IF;
    END IF;
    
    -- Comprobar si toca servicio mayor (cada 30.000 millas)
    IF (NEW.mileage - last_major_service_mileage) >= 30000 THEN
        -- Crear registro de mantenimiento
        INSERT INTO Maintenance (maintenance_date, description, cost)
        VALUES (CURRENT_DATE, 'Major Service due at ' || NEW.mileage || ' miles', 300.00)
        RETURNING id INTO maintenance_id;
        
        -- Vincular al vehículo
        INSERT INTO Vehicle_Maintenance (vehicle_id, maintenance_id)
        VALUES (NEW.id, maintenance_id);
        
        -- Actualizar estado del vehículo
        IF NEW.status_id <> status_id_maintenance THEN
            NEW.status_id := status_id_maintenance;
        END IF;
    END IF;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Crear Trigger BEFORE UPDATE
CREATE TRIGGER before_vehicle_mileage_update
BEFORE UPDATE OF mileage ON Vehicle
FOR EACH ROW
EXECUTE FUNCTION check_vehicle_maintenance();


-- ----------------------------------------------------------------------------
--  TRIGGER 4: ASEGURAR CONSISTENCIA ENTRE RESERVAS Y CONTRATOS
-- ----------------------------------------------------------------------------
/*
Este trigger garantiza que no se puedan crear contratos de alquiler para vehículos que ya están alquilados
o para rangos de fechas superpuestos. Esto mantiene la integridad de los datos y evita la doble reserva.
*/

CREATE OR REPLACE FUNCTION check_vehicle_availability()
RETURNS TRIGGER AS $$
DECLARE
    v_vehicle_id INTEGER;
    status_id_available INTEGER;
    status_id_reserved INTEGER;
BEGIN
    -- Obtener ID del vehículo de la reserva
    SELECT vehicle_id INTO v_vehicle_id
    FROM Reservation
    WHERE id = NEW.reservation_id;
    
    -- Obtener IDs de estado
    SELECT id INTO status_id_available FROM OperationStatus WHERE status_name = 'Available';
    SELECT id INTO status_id_reserved FROM OperationStatus WHERE status_name = 'Reserved';
    
    -- Verificar si el vehículo ya está alquilado para el período dado
    IF EXISTS (
        SELECT 1
        FROM RentalContract rc
        JOIN Reservation r ON rc.reservation_id = r.id
        JOIN OperationStatus s ON rc.status_id = s.id
        WHERE r.vehicle_id = v_vehicle_id
        AND s.status_name = 'Active'
        AND rc.id <> COALESCE(NEW.id, 0)  -- Excluir el contrato actual en caso de actualización
        AND (
            (rc.start_date <= NEW.end_date AND rc.end_date >= NEW.start_date)
            OR
            (NEW.start_date <= rc.end_date AND NEW.end_date >= rc.start_date)
        )
    ) THEN
        RAISE EXCEPTION 'El vehículo ya está alquilado para el período especificado';
    END IF;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Crear Trigger BEFORE INSERT/UPDATE
CREATE TRIGGER before_rental_contract_insert_update
BEFORE INSERT OR UPDATE ON RentalContract
FOR EACH ROW
EXECUTE FUNCTION check_vehicle_availability();


-- ----------------------------------------------------------------------------
--  TRIGGER 5: PREVENIR DOBLE RESERVA
-- ----------------------------------------------------------------------------
/*
Este trigger impide la creación de reservas que causarían una doble reserva de un vehículo. 
Verifica si hay reservas superpuestas para el mismo vehículo y evita la creación de nuevas
reservas que entrarían en conflicto.
*/

CREATE OR REPLACE FUNCTION prevent_double_booking()
RETURNS TRIGGER AS $$
BEGIN
    -- Verificar si hay reservas superpuestas para el mismo vehículo
    IF EXISTS (
        SELECT 1
        FROM Reservation r
        JOIN OperationStatus s ON r.status_id = s.id
        WHERE r.vehicle_id = NEW.vehicle_id
        AND r.id <> COALESCE(NEW.id, 0)  -- Excluir la reserva actual en caso de actualización
        AND s.status_name IN ('Confirmed', 'Active')
        AND (
            (r.start_date <= NEW.end_date AND r.end_date >= NEW.start_date)
            OR
            (NEW.start_date <= r.end_date AND NEW.end_date >= r.start_date)
        )
    ) THEN
        RAISE EXCEPTION 'Este vehículo ya está reservado para el período especificado';
    END IF;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Crear Trigger BEFORE INSERT/UPDATE
CREATE TRIGGER before_reservation_insert_update
BEFORE INSERT OR UPDATE ON Reservation
FOR EACH ROW
EXECUTE FUNCTION prevent_double_booking();