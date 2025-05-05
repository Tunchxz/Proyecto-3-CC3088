<script>
    import { onMount } from 'svelte';
    import Table from '../Components/Table.svelte';
    import { exportToPDF, exportToCSV, exportToExcel } from '../utils/exportUtils.js';

    let maintenances = [];

    // Filtros
    let start_date = '';
    let end_date = '';
    let vehicle_id = '';
    let min_cost = '';
    let max_cost = '';
    let status_id = '';
    let keyword = '';

    const headers = ['ID', 'Vehículo ID', 'Fecha Mantenimiento', 'Descripción', 'Costo', 'Estado ID'];
    const keys = ['id', 'vehicle_id', 'maintenance_date', 'description', 'cost', 'status_id'];
    const formatters = {
        maintenance_date: (date) => new Date(date).toLocaleDateString(),
        cost: (cost) => `Q${parseFloat(cost).toFixed(2)}`
    };

    function buildUrl() {
        const params = new URLSearchParams();

        if (start_date) params.append('start_date', start_date);
        if (end_date) params.append('end_date', end_date);
        if (vehicle_id) params.append('vehicle_id', vehicle_id);
        if (min_cost) params.append('min_cost', min_cost);
        if (max_cost) params.append('max_cost', max_cost);
        if (status_id) params.append('status_id', status_id);
        if (keyword) params.append('keyword', keyword);

        return `http://localhost:9000/report/maintenance?${params.toString()}`;
    }

    function fetchMaintenances() {
        fetch(buildUrl())
            .then(response => response.json())
            .then(data => {
                maintenances = data;
            })
            .catch(error => {
                console.error('Error fetching data:', error);
            });
    }

    onMount(fetchMaintenances);
</script>

<main>
    <h1>Reporte de Mantenimientos</h1>

    <!-- Filtros -->
    <div class="filtros">
        <label>Fecha Inicio: <input type="date" bind:value={start_date} /></label>
        <label>Fecha Fin: <input type="date" bind:value={end_date} /></label>
        <label>Vehículo ID: <input type="number" bind:value={vehicle_id} min="1" /></label>
        <label>Costo Mínimo: <input type="number" bind:value={min_cost} min="0" step="0.01" /></label>
        <label>Costo Máximo: <input type="number" bind:value={max_cost} min="0" step="0.01" /></label>
        <label>Estado ID: <input type="number" bind:value={status_id} min="1" /></label>
        <label>Buscar Descripción: <input type="text" bind:value={keyword} placeholder="aceite, frenos, etc." /></label>
        <button on:click={fetchMaintenances}>Aplicar filtros</button>
    </div>

    <!-- Botones de exportación -->
    <div class="export-buttons">
        <button on:click={() => exportToPDF(maintenances, headers, keys)}>Exportar a PDF</button>
        <button on:click={() => exportToCSV(maintenances)}>Exportar a CSV</button>
        <button on:click={() => exportToExcel(maintenances)}>Exportar a Excel</button>
    </div>

    <!-- Tabla -->
    <div id="maintenance-table" class="raport-container">
        <Table 
            data={maintenances}
            headers={headers}
            keys={keys}
            formatters={formatters}
        />
    </div>
</main>

<style>
    .raport-container {
        margin-top: 1rem;
    }

    .filtros {
        display: flex;
        flex-wrap: wrap;
        gap: 1rem;
        margin-bottom: 1rem;
        justify-content: center;
        align-items: center;
    }

    label {
        display: flex;
        flex-direction: column;
        font-weight: 500;
        color: white;
    }

    input {
        padding: 4px 8px;
        border: 1px solid #ccc;
        border-radius: 4px;
        text-align: center;
        background-color: black;
        color: white;
    }

    button {
        padding: 6px 12px;
        font-weight: bold;
        background-color: #007bff;
        color: white;
        border: none;
        border-radius: 4px;
        cursor: pointer;
    }

    button:hover {
        background-color: #0056b3;
    }

    .export-buttons {
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        gap: 1rem;
        margin: 1rem 0;
    }

    .export-buttons button {
        background-color: #28a745;
    }

    .export-buttons button:hover {
        background-color: #218838;
    }
</style>