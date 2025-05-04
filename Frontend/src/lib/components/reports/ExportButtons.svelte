<script>
  import { reports } from '../../stores/reports';
  import { reportService } from '../../services/reportService';

  const exportToPDF = async () => {
    try {
      const filters = $reports.filters;
      await reportService.exportToPDF(filters.type, filters);
    } catch (error) {
      console.error('Error al exportar a PDF:', error);
    }
  };

  const exportToExcel = async () => {
    try {
      const filters = $reports.filters;
      await reportService.exportToExcel(filters.type, filters);
    } catch (error) {
      console.error('Error al exportar a Excel:', error);
    }
  };

  const exportToCSV = async () => {
    try {
      const filters = $reports.filters;
      await reportService.exportToCSV(filters.type, filters);
    } catch (error) {
      console.error('Error al exportar a CSV:', error);
    }
  };
</script>

<div class="export-buttons">
  <button class="export-btn pdf" on:click={exportToPDF}>
    <i class="fas fa-file-pdf"></i> Exportar a PDF
  </button>
  <button class="export-btn excel" on:click={exportToExcel}>
    <i class="fas fa-file-excel"></i> Exportar a Excel
  </button>
  <button class="export-btn csv" on:click={exportToCSV}>
    <i class="fas fa-file-csv"></i> Exportar a CSV
  </button>
</div>

<style>
  .export-buttons {
    display: flex;
    gap: 1rem;
    margin-top: 2rem;
    justify-content: flex-end;
  }

  .export-btn {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.75rem 1.5rem;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-weight: bold;
    color: white;
  }

  .pdf {
    background-color: #ff4444;
  }

  .excel {
    background-color: #00C851;
  }

  .csv {
    background-color: #33b5e5;
  }

  .export-btn:hover {
    opacity: 0.9;
  }
</style> 