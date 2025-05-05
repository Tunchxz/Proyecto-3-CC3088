<script>
    import { onMount } from 'svelte';
    import Table from '../Components/Table.svelte';
    import { exportToPDF, exportToCSV, exportToExcel } from '../utils/exportUtils.js';

    let payments = [];

    // Filtros
    let start_date = '';
    let end_date = '';
    let payment_method = '';
    let min_amount = '';
    let max_amount = '';
    let status_id = '';

    const headers = ['ID', 'Contrato ID', 'Multa ID', 'Fecha de Pago', 'Monto', 'Método de Pago', 'Estado ID'];
    const keys = ['id', 'rental_contract_id', 'fine_id', 'payment_date', 'amount', 'payment_method', 'status_id'];
    const formatters = {
        payment_date: (date) => new Date(date).toLocaleDateString(),
        amount: (amount) => `Q${amount.toFixed(2)}`
    };

    function buildUrl() {
        const params = new URLSearchParams();

        if (start_date) params.append('start_date', start_date);
        if (end_date) params.append('end_date', end_date);
        if (payment_method) params.append('payment_method', payment_method);
        if (min_amount) params.append('min_amount', min_amount);
        if (max_amount) params.append('max_amount', max_amount);
        if (status_id) params.append('status_id', status_id);

        return `http://localhost:9000/report/payments?${params.toString()}`;
    }

    function fetchPayments() {
        fetch(buildUrl())
            .then(response => response.json())
            .then(data => {
                payments = data;
            })
            .catch(error => {
                console.error('Error fetching data:', error);
            });
    }

    onMount(fetchPayments);
</script>

<main>
    <h1>Reporte de Pagos</h1>

    <!-- Filtros -->
    <div class="filtros">
        <label>Fecha Inicio: <input type="date" bind:value={start_date} /></label>
        <label>Fecha Fin: <input type="date" bind:value={end_date} /></label>
        <label>Método de Pago: <input type="text" bind:value={payment_method} placeholder="Cash, Card, Transfer..." /></label>
        <label>Monto Mínimo: <input type="number" bind:value={min_amount} min="0" step="0.01" /></label>
        <label>Monto Máximo: <input type="number" bind:value={max_amount} min="0" step="0.01" /></label>
        <label>Estado ID: <input type="number" bind:value={status_id} min="1" /></label>
        <button on:click={fetchPayments}>Aplicar filtros</button>
    </div>

    <!-- Botones de exportación -->
    <div class="export-buttons">
        <button on:click={() => exportToPDF(payments, headers, keys)}>Exportar a PDF</button>
        <button on:click={() => exportToCSV(payments)}>Exportar a CSV</button>
        <button on:click={() => exportToExcel(payments)}>Exportar a Excel</button>
    </div>

    <!-- Tabla -->
    <div id="payments-table" class="raport-container">
        <Table 
            data={payments}
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