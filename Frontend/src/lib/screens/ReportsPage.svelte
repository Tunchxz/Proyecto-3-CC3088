<script>
  import ReportList from '../components/reports/ReportList.svelte';
  import ReportFilters from '../components/reports/ReportFilters.svelte';
  import ExportButtons from '../components/reports/ExportButtons.svelte';
  import { onMount } from 'svelte';
  import { reports } from '../stores/reports';
  import { reportService } from '../services/reportService';

  onMount(async () => {
    try {
      reports.setLoading(true);
      const data = await reportService.getReports({ type: 'reservations' });
      reports.setReports(data);
      reports.setFilters({ type: 'reservations' });
    } catch (e) {
      reports.setError(e.message);
    } finally {
      reports.setLoading(false);
    }
  });
</script>

<div class="reports-container">
  <h1>Reportes</h1>
  
  {#if $reports.loading}
    <div class="loading">Cargando reportes...</div>
  {:else if $reports.error}
    <div class="error">{$reports.error}</div>
  {:else}
    <ReportFilters />
    <ReportList />
    <ExportButtons />
  {/if}
</div>

<style>
  .reports-container {
    padding: 2rem;
    max-width: 1200px;
    margin: 0 auto;
  }

  .loading, .error {
    text-align: center;
    padding: 2rem;
    font-size: 1.2rem;
  }

  .error {
    color: red;
  }
</style> 