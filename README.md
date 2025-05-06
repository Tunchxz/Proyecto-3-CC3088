# Proyecto 3 - Sistema de Alquiler de Vehículos

Proyecto 3 – CC3088 Bases de Datos 1  
Universidad del Valle de Guatemala

## Descripción del Proyecto

Este sistema permite generar reportes en tiempo real a partir de una base de datos relacional para un **sistema de alquiler de vehículos**. Incluye funcionalidades como:

- Registro de reservas y contratos de alquiler.
- Gestión de pagos, multas y mantenimientos.
- Generación de reportes con múltiples filtros.
- Interfaz web desarrollada para facilitar la visualización.

El sistema fue construido para demostrar conceptos avanzados de bases de datos, incluyendo normalización, restricciones, funciones, triggers, vistas y consultas complejas.

---

## Estructura del Proyecto

Proyecto-3-CC3088/

├── backend/ # API en Go (Golang)

├── frontend/ # Interfaz de usuario en Svelte

├── database/ # Scripts DDL, datos y funciones SQL

├── database/ # Carpeta con las reflexiones por integrante

├── docker-compose.yml  # Orquestación de servicios

├── README.md # Instrucciones del proyecto

## Instrucciones para ejecutar el proyecto

### 1. Requisitos previos

Asegúrate de tener instalado lo siguiente:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### 2. Clonar el repositorio

```
git clone https://github.com/Tunchxz/Proyecto-3-CC3088.git
cd Proyecto-3-CC3088
```

### 3. Levantar los servicios
Dentro del directorio raíz del proyecto, ejecuta:

```
docker-compose up --build
```

Este comando construirá las imágenes necesarias y levantará los contenedores para backend, frontend y base de datos.

### 4. Acceder a la aplicación
Una vez que los contenedores estén corriendo correctamente, abre tu navegador en:

http://localhost:4000

## ▶️ Video demostrativo del funcionamiento del Proyecto

- Enlace al video: https://youtu.be/PyqFkSsrxlE

