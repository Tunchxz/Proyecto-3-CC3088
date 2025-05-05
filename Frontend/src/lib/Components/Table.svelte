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
    table {
        width: 100%;
        border-collapse: collapse;
        margin: auto;
        position: relative;

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
    }
    
    tr:nth-child(even) {
        background-color: rgb(102, 102, 102);
    }

    .table-container {
        max-height: 60vh;
        max-width:  120vh;
        align-items: center;
        overflow-y: auto;
        margin: auto;
    }
</style>