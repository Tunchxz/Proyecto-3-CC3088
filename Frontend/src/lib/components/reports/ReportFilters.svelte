<script>
  import { reports } from '../../stores/reports';
  import { reportService } from '../../services/reportService';

  let dateFrom = '';
  let dateTo = '';
  let reportType = 'reservations';
  let status = '';

  const handleFilter = async () => {
    try {
      reports.setLoading(true);
      const filters = {
        dateFrom,
        dateTo,
        type: reportType,
        status
      };
      
      const data = await reportService.getReports(filters);
      reports.setReports(data);
      reports.setFilters(filters);
    } catch (error) {
      reports.setError(error.message);
    } finally {
      reports.setLoading(false);
    }
  };
</script>

<div class="filters">
  <div class="filter-group">
    <label for="dateFrom">Fecha Desde:</label>
    <input type="date" id="dateFrom" bind:value={dateFrom}>
  </div>

  <div class="filter-group">
    <label for="dateTo">Fecha Hasta:</label>
    <input type="date" id="dateTo" bind:value={dateTo}>
  </div>

  <div class="filter-group">
    <label for="reportType">Tipo de Reporte:</label>
    <select id="reportType" bind:value={reportType}>
      <option value="reservations">Reservaciones</option>
      <option value="contracts">Contratos</option>
      <option value="payments">Pagos</option>
      <option value="fines">Multas</option>
      <option value="maintenance">Mantenimiento</option>
    </select>
  </div>

  <div class="filter-group">
    <label for="status">Estado:</label>
    <select id="status" bind:value={status}>
      <option value="">Todos</option>
      <option value="completado">Completado</option>
      <option value="pendiente">Pendiente</option>
      <option value="error">Error</option>
    </select>
  </div>

  <button on:click={handleFilter}>Aplicar Filtros</button>
</div>

<style>
  .filters {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: 1rem;
    padding: 1rem;
    background-color: #f5f5f5;
    border-radius: 8px;
    margin-bottom: 2rem;
  }

  .filter-group {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
  }

  label {
    font-weight: bold;
  }

  input, select {
    padding: 0.5rem;
    border: 1px solid #ddd;
    border-radius: 4px;
  }

  button {
    padding: 0.75rem 1.5rem;
    background-color: #4CAF50;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    align-self: end;
  }

  button:hover {
    background-color: #45a049;
  }
</style> 