<script>
    import { onMount } from 'svelte';
    import Table from '../Components/Table.svelte';
    
    let fines = [];
    const headers = ['ID', 'Contrato ID', 'Fecha de Multa', 'Monto', 'Razón', 'Estado ID'];
    const keys = ['id', 'rental_contract_id', 'fine_date', 'amount', 'reason', 'status_id'];
    const formatters = {
        fine_date: (date) => new Date(date).toLocaleDateString(),
        amount: (amount) => `Q${amount.toFixed(2)}`
    };

    onMount(() => {
        fetch('http://localhost:9000/report/fines')
            .then(response => response.json())
            .then(data => {
                fines = data;
            })
            .catch(error => {
                console.error('Error fetching data:', error);
            });
    });
</script>

<div class="content">
    <h1>Reporte de Multas</h1>
    <div class="raport-container">
        <Table 
            data={fines}
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