<script>
    import NavBar from "../Components/NavBar.svelte";
</script>

<div class="welcome-screen">
    <NavBar />
    <div class="welcome-content">
        <h1>Bienvenido al Sistema de Gestión de Vehículos</h1>
        <p>Sistema integral para la administración de reservas, contratos, pagos y mantenimientos de vehículos.</p>
        <div class="features">
            <div class="feature">
                <h3>Gestión de Reservas</h3>
                <p>Control completo de reservaciones de vehículos</p>
            </div>
            <div class="feature">
                <h3>Mantenimientos</h3>
                <p>Seguimiento de mantenimientos vehiculares</p>
            </div>
            <div class="feature">
                <h3>Reportes</h3>
                <p>Informes detallados y estadísticas</p>
            </div>
        </div>
    </div>
</div>

<style>
    .welcome-screen {
        min-height: 100vh;
        background-color: #b8b8ba;
    }

    .welcome-content {
        padding: 2rem;
        text-align: center;
        max-width: 1200px;
        margin: 0 auto;
    }

    h1 {
        font-size: 2.5rem;
        color: #333;
        margin-bottom: 1rem;
    }

    .features {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
        gap: 2rem;
        margin-top: 3rem;
    }

    .feature {
        background-color: white;
        padding: 1.5rem;
        border-radius: 8px;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    }

    .feature h3 {
        color: #111111;
        margin-bottom: 0.5rem;
    }
</style>