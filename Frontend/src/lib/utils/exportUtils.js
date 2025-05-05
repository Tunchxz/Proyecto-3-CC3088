import jsPDF from 'jspdf';
import autoTable from 'jspdf-autotable';
import Papa from 'papaparse';
import * as XLSX from 'xlsx';
import { saveAs } from 'file-saver';

// Exportar a PDF (usando jsPDF + autoTable)
export function exportToPDF(data, headers, keys, filename = 'reporte.pdf') {
    const doc = new jsPDF({
        orientation: 'portrait',
        unit: 'pt',
        format: 'letter' // tamaño carta
    });

    const tableHeaders = headers;
    const tableData = data.map(row =>
        keys.map(key => row[key])
    );

    autoTable(doc, {
        head: [tableHeaders],
        body: tableData,
        margin: { top: 40 },
        styles: { fontSize: 9 },
        headStyles: { fillColor: [0, 0, 0], textColor: [255, 255, 255] }
    });

    doc.save(filename);
}

// Exportar a CSV
export function exportToCSV(data, filename = 'reporte.csv') {
    const csv = Papa.unparse(data);
    const blob = new Blob([csv], { type: 'text/csv;charset=utf-8;' });
    saveAs(blob, filename);
}

// Exportar a Excel
export function exportToExcel(data, filename = 'reporte.xlsx') {
    const worksheet = XLSX.utils.json_to_sheet(data);
    const workbook = XLSX.utils.book_new();
    XLSX.utils.book_append_sheet(workbook, worksheet, 'Reporte');
    const excelBuffer = XLSX.write(workbook, { bookType: 'xlsx', type: 'array' });
    const blob = new Blob([excelBuffer], { type: 'application/octet-stream' });
    saveAs(blob, filename);
}