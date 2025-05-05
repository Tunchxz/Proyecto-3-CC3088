<script>
    import { onMount } from 'svelte';
    import Table from '../Components/Table.svelte';
    
    let payments = [];
    const headers = ['ID', 'Contrato ID', 'Multa ID', 'Fecha de Pago', 'Monto', 'Método de Pago', 'Estado ID'];
    const keys = ['id', 'rental_contract_id', 'fine_id', 'payment_date', 'amount', 'payment_method', 'status_id'];
    const formatters = {
        payment_date: (date) => new Date(date).toLocaleDateString(),
        amount: (amount) => `Q${amount.toFixed(2)}`
    };

    onMount(() => {
        fetch('http://localhost:9000/report/payments')
            .then(response => response.json())
            .then(data => {
                payments = data;
            })
            .catch(error => {
                console.error('Error fetching data:', error);
            });
    });
</script>

<div class="content">
    <h1>Reporte de Pagos</h1>
    <div class="raport-container">
        <Table 
            data={payments}
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