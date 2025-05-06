<script>
    import { onMount } from 'svelte';
    import Table from '../Components/Table.svelte';
    import { exportToPDF, exportToCSV, exportToExcel } from '../utils/exportUtils.js';
    import Chart from 'chart.js/auto';

    let reservations = [];
    let view = 'tabla'; // 'tabla' o 'grafica'
    let chartRef;
    let chartInstance;

    // Filtros
    let start_date = '';
    let end_date = '';
    let customer_name = '';
    let car_plate = '';
    let status_name = '';

    const headers = ['ID', 'Cliente', 'Placa Vehículo', 'Fecha Inicio', 'Fecha Fin', 'Estado'];
    const keys = ['id', 'cliente', 'car_plate', 'start_date', 'end_date', 'status_name'];

    const statusOptions = ['Pending', 'Confirmed', 'Completed', 'Cancelled', 'Available'];

    const formatters = {
        start_date: (date) => date ? new Date(date).toLocaleDateString() : '',
        end_date: (date) => date ? new Date(date).toLocaleDateString() : ''
    };

    function buildUrl() {
        const params = new URLSearchParams();
        if (start_date) params.append('start_date', start_date);
        if (end_date) params.append('end_date', end_date);
        if (customer_name) params.append('customer_name', customer_name);
        if (car_plate) params.append('car_plate', car_plate);
        if (status_name) params.append('status_name', status_name);
        return `http://localhost:9000/report/reservations?${params.toString()}`;
    }

    function fetchReservations() {
        const url = buildUrl();
        fetch(url)
            .then(response => response.json())
            .then(data => {
                reservations = data.map(r => ({
                    ...r,
                    cliente: `${r.first_name ?? ''} ${r.last_name ?? ''}`.trim()
                }));
                if (view === 'grafica') {
                    renderChart();
                }
            })
            .catch(error => {
                console.error('Error fetching data:', error);
            });
    }

    function renderChart() {
        if (chartInstance) {
            chartInstance.destroy();
        }

        const counts = {};
        for (const r of reservations) {
            counts[r.status_name] = (counts[r.status_name] || 0) + 1;
        }

        const labels = Object.keys(counts);
        const data = Object.values(counts);

        chartInstance = new Chart(chartRef, {
            type: 'bar',
            data: {
                labels,
                datasets: [{
                    label: 'Reservaciones por estado',
                    data,
                    backgroundColor: 'rgba(54, 162, 235, 0.6)',
                    borderColor: 'rgba(54, 162, 235, 1)',
                    borderWidth: 1
                }]
            },
            options: {
                responsive: true,
                plugins: {
                    legend: { position: 'top' },
                    title: {
                        display: true,
                        text: 'Cantidad de reservaciones por estado'
                    }
                }
            }
        });
    }

    function switchView(mode) {
        view = mode;
        if (view === 'grafica') {
            setTimeout(renderChart, 0); // esperar render
        }
    }

    onMount(fetchReservations);
</script>

<main>
    <h1>Reporte de Reservaciones</h1>

    <!-- Filtros -->
    <div class="filtros">
        <label>Fecha Inicio:<input type="date" bind:value={start_date} /></label>
        <label>Fecha Fin:<input type="date" bind:value={end_date} /></label>
        <label>Nombre Cliente:<input type="text" bind:value={customer_name} placeholder="Ej. Isa Smallpiece" /></label>
        <label>Placa Vehículo:<input type="text" bind:value={car_plate} placeholder="Ej. TWP-849" /></label>
        <label>Estado:
            <select bind:value={status_name}>
                <option value=''>Todos</option>
                {#each statusOptions as option}
                    <option value={option}>{option}</option>
                {/each}
            </select>
        </label>
        <button on:click={fetchReservations}>Aplicar filtros</button>
    </div>

    <!-- Vista toggle -->
    <div class="vista-toggle">
        <button on:click={() => switchView('tabla')} disabled={view === 'tabla'}>Ver Tabla</button>
        <button on:click={() => switchView('grafica')} disabled={view === 'grafica'}>Ver Gráfica</button>
    </div>

    <!-- Export -->
    {#if view === 'tabla'}
        <div class="export-buttons">
            <button on:click={() => exportToPDF(reservations, headers, keys)}>Exportar a PDF</button>
            <button on:click={() => exportToCSV(reservations)}>Exportar a CSV</button>
            <button on:click={() => exportToExcel(reservations)}>Exportar a Excel</button>
        </div>
    {/if}

    <!-- Tabla o Gráfica -->
    {#if view === 'tabla'}
        <div id="reservations-table" class="raport-container">
            <Table 
                data={reservations}
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
    main {
        padding: 1rem;
    }

    .filtros {
        display: flex;
        flex-wrap: wrap;
        gap: 1rem;
        margin-bottom: 1rem;
        align-items: center;
        justify-content: center;
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

    .raport-container {
        margin-top: 2rem;
    }

    .chart-container {
        max-width: 700px;
        margin: 2rem auto;
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
        margin: 1rem 0 2rem 0;
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