<script>
	import { onMount } from 'svelte';
	import Table from '../Components/Table.svelte';
	import { exportToPDF, exportToCSV, exportToExcel } from '../utils/exportUtils.js';
	import Chart from 'chart.js/auto';

	let maintenances = [];
	let view = 'tabla';
	let chartRef;
	let chartInstance;

	// Filtros
	let start_date = '';
	let end_date = '';
	let car_plate = '';
	let min_cost = '';
	let max_cost = '';
	let status_name = '';
	let keyword = '';

	const headers = ['ID', 'Placa Vehículo', 'Fecha Mantenimiento', 'Descripción', 'Costo', 'Estado'];
	const keys = ['id', 'car_plate', 'maintenance_date', 'description', 'cost', 'status_name'];

	const formatters = {
		maintenance_date: (date) => date ? new Date(date).toLocaleDateString() : '',
		cost: (cost) => `Q${parseFloat(cost).toFixed(2)}`
	};

	const statusOptions = ['Pending', 'Confirmed', 'Completed', 'Cancelled'];

	function buildUrl() {
		const params = new URLSearchParams();
		if (start_date) params.append('start_date', start_date);
		if (end_date) params.append('end_date', end_date);
		if (car_plate) params.append('car_plate', car_plate);
		if (min_cost) params.append('min_cost', min_cost);
		if (max_cost) params.append('max_cost', max_cost);
		if (status_name) params.append('status_name', status_name);
		if (keyword) params.append('description', keyword);
		return `http://localhost:9000/report/maintenance?${params.toString()}`;
	}

	function fetchMaintenances() {
		fetch(buildUrl())
			.then(res => res.json())
			.then(data => {
				maintenances = data;
				if (view === 'grafica') renderChart();
			})
			.catch(err => console.error('Error fetching:', err));
	}

	function renderChart() {
		if (chartInstance) chartInstance.destroy();

		const totals = {};
		for (const m of maintenances) {
			const status = m.status_name || 'Desconocido';
			totals[status] = (totals[status] || 0) + parseFloat(m.cost || 0);
		}

		const labels = Object.keys(totals);
		const data = Object.values(totals);

		chartInstance = new Chart(chartRef, {
			type: 'bar',
			data: {
				labels,
				datasets: [{
					label: 'Costo Total por Estado',
					data,
					backgroundColor: 'rgba(255, 159, 64, 0.6)',
					borderColor: 'rgba(255, 159, 64, 1)',
					borderWidth: 1
				}]
			},
			options: {
				responsive: true,
				plugins: {
					legend: { position: 'top' },
					title: {
						display: true,
						text: 'Mantenimiento total por estado'
					}
				}
			}
		});
	}

	function switchView(mode) {
		view = mode;
		if (view === 'grafica') setTimeout(renderChart, 0);
	}

	onMount(fetchMaintenances);
</script>

<main>
	<h1>Reporte de Mantenimientos</h1>

	<div class="filtros">
		<label>Fecha Inicio: <input type="date" bind:value={start_date} /></label>
		<label>Fecha Fin: <input type="date" bind:value={end_date} /></label>
		<label>Placa Vehículo: <input type="text" bind:value={car_plate} placeholder="Ej: P-123ABC" /></label>
		<label>Costo Mínimo: <input type="number" bind:value={min_cost} min="0" step="0.01" /></label>
		<label>Costo Máximo: <input type="number" bind:value={max_cost} min="0" step="0.01" /></label>
		<label>Estado:
			<select bind:value={status_name}>
				<option value="">Todos</option>
				{#each statusOptions as option}
					<option value={option}>{option}</option>
				{/each}
			</select>
		</label>
		<label>Buscar Descripción: <input type="text" bind:value={keyword} placeholder="aceite, frenos, etc." /></label>
		<button on:click={fetchMaintenances}>Aplicar filtros</button>
	</div>

	<!-- Alternador de vista -->
	<div class="vista-toggle">
		<button on:click={() => switchView('tabla')} disabled={view === 'tabla'}>Ver Tabla</button>
		<button on:click={() => switchView('grafica')} disabled={view === 'grafica'}>Ver Gráfica</button>
	</div>

	<!-- Exportar -->
	{#if view === 'tabla'}
		<div class="export-buttons">
			<button on:click={() => exportToPDF(maintenances, headers, keys)}>Exportar a PDF</button>
			<button on:click={() => exportToCSV(maintenances)}>Exportar a CSV</button>
			<button on:click={() => exportToExcel(maintenances)}>Exportar a Excel</button>
		</div>
	{/if}

	<!-- Contenido -->
	{#if view === 'tabla'}
		<div id="maintenance-table" class="raport-container">
			<Table data={maintenances} headers={headers} keys={keys} formatters={formatters} />
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