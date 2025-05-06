<script>
    import { onMount } from 'svelte';
    import Table from '../Components/Table.svelte';
    import { exportToPDF, exportToCSV, exportToExcel } from '../utils/exportUtils.js';
    import Chart from 'chart.js/auto';

    let payments = [];
    let view = 'tabla'; // tabla | grafica
    let chartRef;
    let chartInstance;

    // Filtros
    let start_date = '';
    let end_date = '';
    let payment_method = '';
    let min_amount = '';
    let max_amount = '';
    let status_name = '';

    const headers = ['ID', 'Contrato ID', 'Multa ID', 'Fecha de Pago', 'Monto', 'Método de Pago', 'Estado'];
    const keys = ['id', 'rental_contract_id', 'fine_id', 'payment_date', 'amount', 'payment_method', 'status_name'];

    const statusOptions = ['Pending', 'Confirmed', 'Completed', 'Cancelled'];
    const methodOptions = ['Cash', 'Card', 'Transfer'];

    const formatters = {
        payment_date: (date) => date ? new Date(date).toLocaleDateString() : '',
        amount: (amount) => `Q${amount?.toFixed(2) ?? '0.00'}`
    };

    function buildUrl() {
        const params = new URLSearchParams();
        if (start_date) params.append('start_date', start_date);
        if (end_date) params.append('end_date', end_date);
        if (payment_method) params.append('payment_method', payment_method);
        if (min_amount) params.append('min_amount', min_amount);
        if (max_amount) params.append('max_amount', max_amount);
        if (status_name) params.append('status_name', status_name);
        return `http://localhost:9000/report/payments?${params.toString()}`;
    }

    function fetchPayments() {
        fetch(buildUrl())
            .then(response => response.json())
            .then(data => {
                payments = data;
                if (view === 'grafica') renderChart();
            })
            .catch(error => {
                console.error('Error fetching data:', error);
            });
    }

    function renderChart() {
        if (chartInstance) chartInstance.destroy();

        const totals = {};
        for (const p of payments) {
            const method = p.payment_method || 'Desconocido';
            totals[method] = (totals[method] || 0) + Number(p.amount || 0);
        }

        const labels = Object.keys(totals);
        const data = Object.values(totals);

        chartInstance = new Chart(chartRef, {
            type: 'bar',
            data: {
                labels,
                datasets: [{
                    label: 'Total Pagado por Método',
                    data,
                    backgroundColor: 'rgba(75, 192, 192, 0.6)',
                    borderColor: 'rgba(75, 192, 192, 1)',
                    borderWidth: 1
                }]
            },
            options: {
                responsive: true,
                plugins: {
                    legend: { position: 'top' },
                    title: {
                        display: true,
                        text: 'Pagos agrupados por método'
                    }
                }
            }
        });
    }

    function switchView(mode) {
        view = mode;
        if (view === 'grafica') setTimeout(renderChart, 0);
    }

    onMount(fetchPayments);
</script>

<main>
    <h1>Reporte de Pagos</h1>

    <!-- Filtros -->
    <div class="filtros">
        <label>Fecha Inicio: <input type="date" bind:value={start_date} /></label>
        <label>Fecha Fin: <input type="date" bind:value={end_date} /></label>
        <label>Método de Pago:
            <select bind:value={payment_method}>
                <option value="">Todos</option>
                {#each methodOptions as method}
                    <option value={method}>{method}</option>
                {/each}
            </select>
        </label>
        <label>Monto Mínimo: <input type="number" bind:value={min_amount} min="0" step="0.01" /></label>
        <label>Monto Máximo: <input type="number" bind:value={max_amount} min="0" step="0.01" /></label>
        <label>Estado:
            <select bind:value={status_name}>
                <option value="">Todos</option>
                {#each statusOptions as status}
                    <option value={status}>{status}</option>
                {/each}
            </select>
        </label>
        <button on:click={fetchPayments}>Aplicar filtros</button>
    </div>

    <!-- Toggle tabla/grafica -->
    <div class="vista-toggle">
        <button on:click={() => switchView('tabla')} disabled={view === 'tabla'}>Ver Tabla</button>
        <button on:click={() => switchView('grafica')} disabled={view === 'grafica'}>Ver Gráfica</button>
    </div>

    <!-- Export -->
    {#if view === 'tabla'}
        <div class="export-buttons">
            <button on:click={() => exportToPDF(payments, headers, keys)}>Exportar a PDF</button>
            <button on:click={() => exportToCSV(payments)}>Exportar a CSV</button>
            <button on:click={() => exportToExcel(payments)}>Exportar a Excel</button>
        </div>
    {/if}

    <!-- Tabla o gráfica -->
    {#if view === 'tabla'}
        <div id="payments-table" class="raport-container">
            <Table 
                data={payments}
                headers={headers}
                keys={keys}
                formatters={formatters}
            />
        </div>
    {:else}
        <div class="chart-container">
            <canvas bind:this={chartRef}></canvas>
        </div>
    {/if}
</main>

<style>
    .raport-container {
        margin-top: 1rem;
    }

    .chart-container {
        max-width: 700px;
        margin: 2rem auto;
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

    input, select {
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

    .vista-toggle {
        display: flex;
        justify-content: center;
        gap: 1rem;
        margin: 1rem 0;
    }

    .vista-toggle button {
        background-color: #6c757d;
    }

    .vista-toggle button[disabled] {
        background-color: #343a40;
        cursor: not-allowed;
    }
</style>