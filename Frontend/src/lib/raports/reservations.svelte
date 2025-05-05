<script>
    import { onMount } from 'svelte';
    import Table from '../Components/Table.svelte';
    
    let reservations = [];
    const headers = ['ID', 'Cliente ID', 'Vehiculo ID', 'Fecha Inicio', 'Fecha Fin', 'Estado ID'];
    const keys = ['id', 'customer_id', 'vehicle_id', 'start_date', 'end_date', 'status_id'];
    const formatters = {
        start_date: (date) => new Date(date).toLocaleDateString(),
        end_date: (date) => new Date(date).toLocaleDateString()
    };

    onMount(() => {
        fetch('http://localhost:9000/report/reservations')
            .then(response => response.json())
            .then(data => {
                reservations = data;
            })
            .catch(error => {
                console.error('Error fetching data:', error);
            });
    });
</script>

<div class="">
    <h1>Reporte de Reservaciones</h1>
    <div class="raport-container">
        <Table 
            data={reservations}
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