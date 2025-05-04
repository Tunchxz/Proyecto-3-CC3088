-- =============================
--  EJERCICIO 5 (ddl)
-- =============================

-- -----------------------------
--  ELIMINAR TABLAS
-- -----------------------------

DROP TABLE IF EXISTS Refund CASCADE;
DROP TABLE IF EXISTS Payment CASCADE;
DROP TABLE IF EXISTS Contract_Fine CASCADE;
DROP TABLE IF EXISTS Fine CASCADE;
DROP TABLE IF EXISTS RentalContract CASCADE;
DROP TABLE IF EXISTS Reservation CASCADE;
DROP TABLE IF EXISTS Vehicle_Maintenance CASCADE;
DROP TABLE IF EXISTS Maintenance CASCADE;
DROP TABLE IF EXISTS Vehicle CASCADE;
DROP TABLE IF EXISTS Rates CASCADE;
DROP TABLE IF EXISTS Model_Color CASCADE;
DROP TABLE IF EXISTS Color CASCADE;
DROP TABLE IF EXISTS Model CASCADE;
DROP TABLE IF EXISTS VehicleType CASCADE;
DROP TABLE IF EXISTS Manufacturer CASCADE;
DROP TABLE IF EXISTS Facility CASCADE;
DROP TABLE IF EXISTS Customer_Address CASCADE;
DROP TABLE IF EXISTS Customer CASCADE;
DROP TABLE IF EXISTS Address CASCADE;
DROP TABLE IF EXISTS Country CASCADE;
DROP TABLE IF EXISTS OperationStatus CASCADE;

-- -----------------------------
--  CREAR TABLAS (PRINCIPALES)
-- -----------------------------

-- Tabla: Países
CREATE TABLE Country (
    id SERIAL PRIMARY KEY,
    country_name VARCHAR(64) NOT NULL UNIQUE
);

-- Tabla: Estados
CREATE TABLE OperationStatus (
    id SERIAL PRIMARY KEY,
    status_name VARCHAR(64) NOT NULL UNIQUE
);

-- Tabla: Direcciones
CREATE TABLE Address (
    id SERIAL PRIMARY KEY,
    unit_number VARCHAR(8),
    street_number VARCHAR(8),
    address_line_1 VARCHAR(64) NOT NULL,
    address_line_2 VARCHAR(64),
    city VARCHAR(64) NOT NULL,
    region VARCHAR(64),
    country_id INTEGER NOT NULL REFERENCES Country(id)
);

-- Tabla: Clientes
CREATE TABLE Customer (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    date_of_birth DATE NOT NULL,
    driver_license_number VARCHAR(64) NOT NULL UNIQUE,
    email VARCHAR(128) NOT NULL UNIQUE,
    phone_number VARCHAR(16) NOT NULL,
    registration_date DATE NOT NULL DEFAULT CURRENT_DATE
);

-- Tabla: Cliente-Dirección
CREATE TABLE Customer_Address (
    customer_id INTEGER NOT NULL REFERENCES Customer(id),
    address_id INTEGER NOT NULL REFERENCES Address(id),
    PRIMARY KEY (customer_id, address_id)
);

-- Tabla: Fabricantes
CREATE TABLE Manufacturer (
    id SERIAL PRIMARY KEY,
    manufacturer_name VARCHAR(128) NOT NULL UNIQUE
);

-- Tabla: Tipos de Vehículo
CREATE TABLE VehicleType (
    id SERIAL PRIMARY KEY,
    vehicle_type_name VARCHAR(32) NOT NULL UNIQUE
);

-- Tabla: Modelos
CREATE TABLE Model (
    id SERIAL PRIMARY KEY,
    manufacturer_id INTEGER NOT NULL REFERENCES Manufacturer(id),
    model_name VARCHAR(128) NOT NULL,
    transmission_type VARCHAR NOT NULL CHECK (transmission_type IN ('Manual', 'Automatic')),
    number_of_seats INTEGER NOT NULL CHECK (number_of_seats > 0),
    vehicle_type_id INTEGER NOT NULL REFERENCES VehicleType(id),
    UNIQUE (manufacturer_id, model_name)
);

-- Tabla: Colores
CREATE TABLE Color (
    id SERIAL PRIMARY KEY,
    color_name VARCHAR(16) NOT NULL UNIQUE
);

-- Tabla: Modelo-Color
CREATE TABLE Model_Color (
    model_id INTEGER NOT NULL REFERENCES Model(id),
    color_id INTEGER NOT NULL REFERENCES Color(id),
    PRIMARY KEY (model_id, color_id)
);

-- Tabla: Sucursales
CREATE TABLE Facility (
    id SERIAL PRIMARY KEY,
    facility_name VARCHAR(128) NOT NULL UNIQUE,
    facility_phone_number VARCHAR(16) NOT NULL UNIQUE,
    address_id INTEGER NOT NULL REFERENCES Address(id)
);

-- Tabla: Precios
CREATE TABLE Rates (
    id SERIAL PRIMARY KEY,
    daily_rate DECIMAL(10,2) NOT NULL CHECK (daily_rate >= 0),
    weekly_rate DECIMAL(10,2) NOT NULL CHECK (weekly_rate >= 0),
    monthly_rate DECIMAL(10,2) NOT NULL CHECK (monthly_rate >= 0)
);

-- Tabla: Vehículos
CREATE TABLE Vehicle (
    id SERIAL PRIMARY KEY,
    model_id INTEGER NOT NULL REFERENCES Model(id),
    facility_id INTEGER NOT NULL REFERENCES Facility(id),
    car_plate VARCHAR(64) NOT NULL UNIQUE,
    mileage INTEGER NOT NULL CHECK (mileage >= 0),
    status_id INTEGER NOT NULL REFERENCES OperationStatus(id),
    rates_id INTEGER NOT NULL REFERENCES Rates(id)
);

-- Tabla: Mantenimientos
CREATE TABLE Maintenance (
    id SERIAL PRIMARY KEY,
    maintenance_date DATE NOT NULL DEFAULT CURRENT_DATE,
    description TEXT NOT NULL,
    cost DECIMAL(10,2) NOT NULL CHECK (cost >= 0)
);

-- Tabla: Mantenimiento-Vehículo
CREATE TABLE Vehicle_Maintenance (
    vehicle_id INTEGER NOT NULL REFERENCES Vehicle(id),
    maintenance_id INTEGER NOT NULL REFERENCES Maintenance(id),
    PRIMARY KEY (vehicle_id, maintenance_id)
);

-- Tabla: Reservaciones
CREATE TABLE Reservation (
    id SERIAL PRIMARY KEY,
    customer_id INTEGER NOT NULL REFERENCES Customer(id),
    vehicle_id INTEGER NOT NULL REFERENCES Vehicle(id),
    reservation_date DATE NOT NULL DEFAULT CURRENT_DATE,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL CHECK (end_date >= start_date),
    status_id INTEGER NOT NULL REFERENCES OperationStatus(id)
);

-- Tabla: Contratos de Renta
CREATE TABLE RentalContract (
    id SERIAL PRIMARY KEY,
    reservation_id INTEGER NOT NULL UNIQUE REFERENCES Reservation(id),
    start_date DATE NOT NULL,
    end_date DATE NOT NULL CHECK (end_date >= start_date),
    status_id INTEGER NOT NULL REFERENCES OperationStatus(id)
);

-- Tabla: Multas
CREATE TABLE Fine (
    id SERIAL PRIMARY KEY,
    rental_contract_id INTEGER NOT NULL REFERENCES RentalContract(id),
    fine_date DATE NOT NULL DEFAULT CURRENT_DATE,
    amount DECIMAL(10,2) NOT NULL CHECK (amount >= 0),
    reason TEXT NOT NULL,
    status_id INTEGER NOT NULL REFERENCES OperationStatus(id)
);

-- Tabla: Contratos-Multas
CREATE TABLE Contract_Fine (
    rental_contract_id INTEGER NOT NULL REFERENCES RentalContract(id),
    fine_id INTEGER NOT NULL REFERENCES Fine(id),
    PRIMARY KEY (rental_contract_id, fine_id)
);

-- Tabla: Pagos
CREATE TABLE Payment (
    id SERIAL PRIMARY KEY,
    rental_contract_id INTEGER REFERENCES RentalContract(id),
    fine_id INTEGER REFERENCES Fine(id),
    payment_date DATE NOT NULL DEFAULT CURRENT_DATE,
    amount DECIMAL(10,2) NOT NULL CHECK (amount >= 0),
    payment_method VARCHAR NOT NULL CHECK (payment_method IN ('Cash', 'Card', 'Transfer')),
    status_id INTEGER NOT NULL REFERENCES OperationStatus(id)
);

-- Tabla: Rembolsos
CREATE TABLE Refund (
    id SERIAL PRIMARY KEY,
    payment_id INTEGER NOT NULL REFERENCES Payment(id),
    refund_date DATE NOT NULL DEFAULT CURRENT_DATE,
    amount DECIMAL(10,2) NOT NULL CHECK (amount >= 0),
    reason TEXT NOT NULL,
    status_id INTEGER NOT NULL REFERENCES OperationStatus(id)
);