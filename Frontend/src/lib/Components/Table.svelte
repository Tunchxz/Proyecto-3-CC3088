<script>
    export let data = [];
    export let headers = [];
    export let keys = [];
    export let formatters = {};
</script>

<div class="table-container">
    <table>
        <thead>
            <tr>
                {#each headers as header}
                    <th>{header}</th>
                {/each}
            </tr>
        </thead>
        <tbody>
            {#each data as row}
                <tr>
                    {#each keys as key}
                        <td>
                            {#if formatters[key]}
                                {formatters[key](row[key])}
                            {:else}
                                {row[key]}
                            {/if}
                        </td>
                    {/each}
                </tr>
            {/each}
        </tbody>
    </table>
</div>

<style>
    .table-container {
        max-height: 55vh;         /* Limita la altura vertical */
        max-width: 120vh;         /* Se mantiene el ancho horizontal que usabas */
        overflow-y: auto;         /* Scroll vertical si hay muchos registros */
        margin: 0 auto 2rem auto; /* Centrado + espacio inferior para los botones */
        border: 1px solid #444;
        border-radius: 6px;
    }

    table {
        width: 100%;
        border-collapse: collapse;
    }
    
    th, td {
        padding: 0.5rem;
        text-align: left;
        border: 1px solid #ddd;
    }
    
    th {
        background-color: #0f0f0f;
        font-weight: bold;
        color: white;
        position: sticky;
        top: 0;
        z-index: 1;
    }

    tr:nth-child(even) {
        background-color: rgb(102, 102, 102);
    }
</style>