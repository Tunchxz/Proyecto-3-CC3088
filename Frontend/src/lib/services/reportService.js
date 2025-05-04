const API_BASE_URL = 'http://localhost:9000';

const defaultHeaders = {
  'Content-Type': 'application/json',
  'Accept': 'application/json'
};

export const reportService = {
  async getReports(filters = {}) {
    try {
      const queryParams = new URLSearchParams(filters);
      const response = await fetch(`${API_BASE_URL}/report/${filters.type}?${queryParams}`, {
        method: 'GET',
        headers: defaultHeaders,
        credentials: 'include'
      });
      
      if (!response.ok) {
        const errorData = await response.json().catch(() => ({}));
        throw new Error(errorData.message || 'Error al obtener los reportes');
      }
      
      return await response.json();
    } catch (error) {
      console.error('Error en reportService.getReports:', error);
      throw error;
    }
  },

  async exportToPDF(reportType, filters = {}) {
    try {
      const queryParams = new URLSearchParams(filters);
      const response = await fetch(`${API_BASE_URL}/report/${reportType}/pdf?${queryParams}`, {
        method: 'GET',
        headers: {
          ...defaultHeaders,
          'Accept': 'application/pdf'
        },
        credentials: 'include'
      });
      
      if (!response.ok) {
        throw new Error('Error al exportar a PDF');
      }
      
      const blob = await response.blob();
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.download = `reporte_${reportType}.pdf`;
      document.body.appendChild(a);
      a.click();
      window.URL.revokeObjectURL(url);
      document.body.removeChild(a);
    } catch (error) {
      console.error('Error en reportService.exportToPDF:', error);
      throw error;
    }
  },

  async exportToExcel(reportType, filters = {}) {
    try {
      const queryParams = new URLSearchParams(filters);
      const response = await fetch(`${API_BASE_URL}/report/${reportType}/excel?${queryParams}`, {
        method: 'GET',
        headers: {
          ...defaultHeaders,
          'Accept': 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'
        },
        credentials: 'include'
      });
      
      if (!response.ok) {
        throw new Error('Error al exportar a Excel');
      }
      
      const blob = await response.blob();
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.download = `reporte_${reportType}.xlsx`;
      document.body.appendChild(a);
      a.click();
      window.URL.revokeObjectURL(url);
      document.body.removeChild(a);
    } catch (error) {
      console.error('Error en reportService.exportToExcel:', error);
      throw error;
    }
  },

  async exportToCSV(reportType, filters = {}) {
    try {
      const queryParams = new URLSearchParams(filters);
      const response = await fetch(`${API_BASE_URL}/report/${reportType}/csv?${queryParams}`, {
        method: 'GET',
        headers: {
          ...defaultHeaders,
          'Accept': 'text/csv'
        },
        credentials: 'include'
      });
      
      if (!response.ok) {
        throw new Error('Error al exportar a CSV');
      }
      
      const blob = await response.blob();
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.download = `reporte_${reportType}.csv`;
      document.body.appendChild(a);
      a.click();
      window.URL.revokeObjectURL(url);
      document.body.removeChild(a);
    } catch (error) {
      console.error('Error en reportService.exportToCSV:', error);
      throw error;
    }
  }
}; 