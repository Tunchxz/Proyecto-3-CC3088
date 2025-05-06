<script>
	import { onMount } from 'svelte';
	import Table from '../Components/Table.svelte';
	import { exportToPDF, exportToCSV, exportToExcel } from '../utils/exportUtils.js';
	import Chart from 'chart.js/auto';

	let fines = [];
	let view = 'tabla';
	let chartRef;
	let chartInstance;

	// Filtros
	let start_date = '';
	let end_date = '';
	let min_amount = '';
	let max_amount = '';
	let status_name = '';
	let keyword = '';

	const headers = ['ID', 'Contrato ID', 'Fecha de Multa', 'Monto', 'Razón', 'Estado'];
	const keys = ['id', 'rental_contract_id', 'fine_date', 'amount', 'reason', 'status_name'];

	const statusOptions = ['Pending', 'Confirmed', 'Completed', 'Cancelled'];

	const formatters = {
		fine_date: (date) => date ? new Date(date).toLocaleDateString() : '',
		amount: (amount) => `Q${amount?.toFixed(2) ?? '0.00'}`
	};

	function buildUrl() {
		const params = new URLSearchParams();
		if (start_date) params.append('start_date', start_date);
		if (end_date) params.append('end_date', end_date);
		if (min_amount) params.append('min_amount', min_amount);
		if (max_amount) params.append('max_amount', max_amount);
		if (status_name) params.append('status_name', status_name);
		if (keyword) params.append('keyword', keyword);
		return `http://localhost:9000/report/fines?${params.toString()}`;
	}

	function fetchFines() {
		fetch(buildUrl())
			.then(res => res.json())
			.then(data => {
				fines = data;
				if (view === 'grafica') renderChart();
			})
			.catch(err => console.error('Error:', err));
	}

	function renderChart() {
		if (chartInstance) chartInstance.destroy();

		const totals = {};
		for (const fine of fines) {
			const status = fine.status_name || 'Desconocido';
			totals[status] = (totals[status] || 0) + parseFloat(fine.amount || 0);
		}

		const labels = Object.keys(totals);
		const values = Object.values(totals);

		chartInstance = new Chart(chartRef, {
			type: 'bar',
			data: {
				labels,
				datasets: [{
					label: 'Total de Multas por Estado',
					data: values,
					backgroundColor: 'rgba(255, 99, 132, 0.6)',
					borderColor: 'rgba(255, 99, 132, 1)',
					borderWidth: 1
				}]
			},
			options: {
				responsive: true,
				plugins: {
					legend: { position: 'top' },
					title: {
						display: true,
						text: 'Total Q de Multas por Estado'
					}
				}
			}
		});
	}

	function switchView(mode) {
		view = mode;
		if (view === 'grafica') setTimeout(renderChart, 0);
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
		<label>Estado:
			<select bind:value={status_name}>
				<option value="">Todos</option>
				{#each statusOptions as option}
					<option value={option}>{option}</option>
				{/each}
			</select>
		</label>
		<label>Buscar Razón: <input type="text" bind:value={keyword} placeholder="Ej: speeding, accident..." /></label>
		<button on:click={fetchFines}>Aplicar filtros</button>
	</div>

	<!-- Alternar vista -->
	<div class="vista-toggle">
		<button on:click={() => switchView('tabla')} disabled={view === 'tabla'}>Ver Tabla</button>
		<button on:click={() => switchView('grafica')} disabled={view === 'grafica'}>Ver Gráfica</button>
	</div>

	<!-- Exportar -->
	{#if view === 'tabla'}
		<div class="export-buttons">
			<button on:click={() => exportToPDF(fines, headers, keys)}>Exportar a PDF</button>
			<button on:click={() => exportToCSV(fines)}>Exportar a CSV</button>
			<button on:click={() => exportToExcel(fines)}>Exportar a Excel</button>
		</div>
	{/if}

	<!-- Contenido -->
	{#if view === 'tabla'}
		<div id="fines-table" class="raport-container">
			<Table data={fines} headers={headers} keys={keys} formatters={formatters} />
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