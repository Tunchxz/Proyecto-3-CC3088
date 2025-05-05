<script>
    import { onMount } from 'svelte';
    import Table from '../Components/Table.svelte';
    import { exportToPDF, exportToCSV, exportToExcel } from '../utils/exportUtils.js';

    let contracts = [];

    // Filtros
    let start_date = '';
    let end_date = '';
    let customer_id = '';
    let status_id = '';

    const headers = ['ID', 'Reservación ID', 'Fecha Inicio', 'Fecha Fin', 'Estado ID'];
    const keys = ['id', 'reservation_id', 'start_date', 'end_date', 'status_id'];
    const formatters = {
        start_date: (date) => new Date(date).toLocaleDateString(),
        end_date: (date) => new Date(date).toLocaleDateString()
    };

    function buildUrl() {
        const params = new URLSearchParams();

        if (start_date) params.append('start_date', start_date);
        if (end_date) params.append('end_date', end_date);
        if (customer_id) params.append('customer_id', customer_id);
        if (status_id) params.append('status_id', status_id);

        return `http://localhost:9000/report/contracts?${params.toString()}`;
    }

    function fetchContracts() {
        const url = buildUrl();
        fetch(url)
            .then(response => response.json())
            .then(data => {
                contracts = data;
            })
            .catch(error => {
                console.error('Error fetching data:', error);
            });
    }

    onMount(fetchContracts);
</script>

<main>
    <h1>Reporte de Contratos</h1>

    <!-- Filtros -->
    <div class="filtros">
        <label>Fecha Inicio: <input type="date" bind:value={start_date} /></label>
        <label>Fecha Fin: <input type="date" bind:value={end_date} /></label>
        <label>Cliente ID: <input type="number" bind:value={customer_id} min="1" /></label>
        <label>Estado ID: <input type="number" bind:value={status_id} min="1" /></label>
        <button on:click={fetchContracts}>Aplicar filtros</button>
    </div>

    <!-- Botones de exportación -->
    <div class="export-buttons">
        <button on:click={() => exportToPDF(contracts, headers, keys)}>Exportar a PDF</button>
        <button on:click={() => exportToCSV(contracts)}>Exportar a CSV</button>
        <button on:click={() => exportToExcel(contracts)}>Exportar a Excel</button>
    </div>

    <!-- Tabla -->
    <div id="contracts-table" class="raport-container">
        <Table 
            data={contracts}
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