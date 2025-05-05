<script>
    import { onMount } from 'svelte';
    import Table from '../Components/Table.svelte';
    
    let contracts = [];
    const headers = ['ID', 'Reservación ID', 'Fecha Inicio', 'Fecha Fin', 'Estado ID'];
    const keys = ['id', 'reservation_id', 'start_date', 'end_date', 'status_id'];
    const formatters = {
        start_date: (date) => new Date(date).toLocaleDateString(),
        end_date: (date) => new Date(date).toLocaleDateString()
    };

    onMount(() => {
        fetch('http://localhost:9000/report/contracts')
            .then(response => response.json())
            .then(data => {
                contracts = data;
            })
            .catch(error => {
                console.error('Error fetching data:', error);
            });
    });
</script>

<div class="content">
    <h1>Reporte de Contratos</h1>
    <div class="raport-container">
        <Table 
            data={contracts}
            headers={headers}
            keys={keys}
            formatters={formatters}
        />
    </div>
</div>

<style>
    .raport-container {
        margin: auto;
    }
</style>

