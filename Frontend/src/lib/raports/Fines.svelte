<script>
    import { onMount } from 'svelte';
    import Table from '../Components/Table.svelte';
    import { exportToPDF, exportToCSV, exportToExcel } from '../utils/exportUtils.js';

    let fines = [];

    // Filtros
    let start_date = '';
    let end_date = '';
    let min_amount = '';
    let max_amount = '';
    let status_id = '';
    let keyword = '';

    const headers = ['ID', 'Contrato ID', 'Fecha de Multa', 'Monto', 'Razón', 'Estado ID'];
    const keys = ['id', 'rental_contract_id', 'fine_date', 'amount', 'reason', 'status_id'];
    const formatters = {
        fine_date: (date) => new Date(date).toLocaleDateString(),
        amount: (amount) => `Q${amount.toFixed(2)}`
    };

    function buildUrl() {
        const params = new URLSearchParams();

        if (start_date) params.append('start_date', start_date);
        if (end_date) params.append('end_date', end_date);
        if (min_amount) params.append('min_amount', min_amount);
        if (max_amount) params.append('max_amount', max_amount);
        if (status_id) params.append('status_id', status_id);
        if (keyword) params.append('keyword', keyword);

        return `http://localhost:9000/report/fines?${params.toString()}`;
    }

    function fetchFines() {
        fetch(buildUrl())
            .then(response => response.json())
            .then(data => {
                fines = data;
            })
            .catch(error => {
                console.error('Error fetching data:', error);
            });
    }

    onMount(fetchFines);
</script>

<main>
    <h1>Reporte de Multas</h1>

    <!-- Filtros -->
    <div class="filtros">
        <label>Fecha Inicio: <input type="date" bind:value={start_date} /></label>
        <label>Fecha Fin: <input type="date" bind:value={end_date} /></label>
        <label>Monto Mínimo: <input type="number" bind:value={min_amount} min="0" step="0.01" /></label>
        <label>Monto Máximo: <input type="number" bind:value={max_amount} min="0" step="0.01" /></label>
        <label>Estado ID: <input type="number" bind:value={status_id} min="1" /></label>
        <label>Buscar Razón: <input type="text" bind:value={keyword} placeholder="e.g. racing, accident..." /></label>
        <button on:click={fetchFines}>Aplicar filtros</button>
    </div>

    <!-- Botones de exportación -->
    <div class="export-buttons">
        <button on:click={() => exportToPDF(fines, headers, keys)}>Exportar a PDF</button>
        <button on:click={() => exportToCSV(fines)}>Exportar a CSV</button>
        <button on:click={() => exportToExcel(fines)}>Exportar a Excel</button>
    </div>

    <!-- Tabla -->
    <div id="fines-table" class="raport-container">
        <Table 
            data={fines}
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