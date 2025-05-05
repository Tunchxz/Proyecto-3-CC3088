<script>
    import { onMount } from 'svelte';
    import Chart from 'chart.js/auto';

    export let labels = [];
    export let data = [];
    export let title = 'Gráfico de Pastel';

    let canvas;
    let chartInstance;

    onMount(() => {
        if (chartInstance) {
            chartInstance.destroy();
        }

        chartInstance = new Chart(canvas, {
            type: 'pie',
            data: {
                labels,
                datasets: [{
                    label: title,
                    data,
                    backgroundColor: [
                        '#FF6384', '#36A2EB', '#FFCE56', '#00C49F', '#8A2BE2', '#FFA07A', '#20B2AA'
                    ]
                }]
            },
            options: {
                responsive: true,
                plugins: {
                    legend: { position: 'bottom' },
                    title: { display: true, text: title }
                }
            }
        });
    });
</script>

<canvas bind:this={canvas}></canvas>

<style>
    canvas {
        max-width: 500px;
        margin: 2rem auto;
        display: block;
        background-color: white;
        padding: 1rem;
        border-radius: 10px;
    }
</style>