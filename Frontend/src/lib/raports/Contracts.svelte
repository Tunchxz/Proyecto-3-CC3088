<script>
	import { onMount } from 'svelte';
	import Table from '../Components/Table.svelte';
	import { exportToPDF, exportToCSV, exportToExcel } from '../utils/exportUtils.js';
	import Chart from 'chart.js/auto';

	let contracts = [];
	let view = 'tabla';
	let chartRef;
	let chartInstance;

	// Filtros
	let start_date = '';
	let end_date = '';
	let reservation_id = '';
	let status_name = '';

	const headers = ['ID', 'Reservación ID', 'Fecha Inicio', 'Fecha Fin', 'Estado'];
	const keys = ['id', 'reservation_id', 'start_date', 'end_date', 'status_name'];

	const statusOptions = ['Pending', 'Confirmed', 'Completed', 'Cancelled', 'Available'];

	const formatters = {
		start_date: (date) => date ? new Date(date).toLocaleDateString() : '',
		end_date: (date) => date ? new Date(date).toLocaleDateString() : ''
	};

	function buildUrl() {
		const params = new URLSearchParams();
		if (start_date) params.append('start_date', start_date);
		if (end_date) params.append('end_date', end_date);
		if (reservation_id) params.append('reservation_id', reservation_id);
		if (status_name) params.append('status_name', status_name);
		return `http://localhost:9000/report/contracts?${params.toString()}`;
	}

	function fetchContracts() {
		const url = buildUrl();
		fetch(url)
			.then(res => res.json())
			.then(data => {
				contracts = data;
				if (view === 'grafica') renderChart();
			})
			.catch(err => console.error('Error:', err));
	}

	function renderChart() {
		if (chartInstance) chartInstance.destroy();

		const totals = {};
		for (const c of contracts) {
			const estado = c.status_name || 'Desconocido';
			totals[estado] = (totals[estado] || 0) + 1;
		}

		const labels = Object.keys(totals);
		const values = Object.values(totals);

		chartInstance = new Chart(chartRef, {
			type: 'bar',
			data: {
				labels,
				datasets: [{
					label: 'Número de Contratos por Estado',
					data: values,
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
						text: 'Contratos por Estado'
					}
				}
			}
		});
	}

	function switchView(mode) {
		view = mode;
		if (view === 'grafica') setTimeout(renderChart, 0);
	}

	onMount(fetchContracts);
</script>

<main>
	<h1>Reporte de Contratos</h1>

	<!-- Filtros -->
	<div class="filtros">
		<label>Fecha Inicio: <input type="date" bind:value={start_date} /></label>
		<label>Fecha Fin: <input type="date" bind:value={end_date} /></label>
		<label>Reservación ID: <input type="number" bind:value={reservation_id} min="1" /></label>
		<label>Estado:
			<select bind:value={status_name}>
				<option value=''>Todos</option>
				{#each statusOptions as option}
					<option value={option}>{option}</option>
				{/each}
			</select>
		</label>
		<button on:click={fetchContracts}>Aplicar filtros</button>
	</div>

	<!-- Alternar vista -->
	<div class="vista-toggle">
		<button on:click={() => switchView('tabla')} disabled={view === 'tabla'}>Ver Tabla</button>
		<button on:click={() => switchView('grafica')} disabled={view === 'grafica'}>Ver Gráfica</button>
	</div>

	<!-- Exportar -->
	{#if view === 'tabla'}
		<div class="export-buttons">
			<button on:click={() => exportToPDF(contracts, headers, keys)}>Exportar a PDF</button>
			<button on:click={() => exportToCSV(contracts)}>Exportar a CSV</button>
			<button on:click={() => exportToExcel(contracts)}>Exportar a Excel</button>
		</div>
	{/if}

	<!-- Contenido -->
	{#if view === 'tabla'}
		<div id="contracts-table" class="raport-container">
			<Table data={contracts} headers={headers} keys={keys} formatters={formatters} />
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

	.chart-container {
		max-width: 700px;
		margin: 2rem auto;
	}
</style>