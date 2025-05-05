<script>
    import { onMount } from 'svelte';
    import Table from '../Components/Table.svelte';
    
    let maintenances = [];
    const headers = ['ID', 'Vehiculo ID', 'Fecha Mantenimiento', 'Descripción', 'Costo', 'Estado ID'];
    const keys = ['id', 'vehicle_id', 'maintenance_date', 'description', 'cost', 'status_id'];
    const formatters = {
        maintenance_date: (date) => new Date(date).toLocaleDateString(),
        cost: (cost) => `$${parseFloat(cost).toFixed(2)}`
    };

    onMount(() => {
        fetch('http://localhost:9000/report/maintenance')
            .then(response => response.json())
            .then(data => {
                maintenances = data;
            })
            .catch(error => {
                console.error('Error fetching data:', error);
            });
    });
</script>

<div class="">
    <h1>Reporte de Mantenimientos</h1>
    <div class="raport-container">
        <Table 
            data={maintenances}
            headers={headers}
            keys={keys}
            formatters={formatters}
        />
    </div>
</div>

<style>
    .raport-container{
        margin: auto;
    }
</style>