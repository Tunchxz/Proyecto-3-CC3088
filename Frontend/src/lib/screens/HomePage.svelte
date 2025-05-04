<script>
    import NavBar from "../elements/NavBar.svelte";
    import Raport from "../elements/raport.svelte";
    import Sales from "../elements/sales.svelte";

    const SCREENS = {
        RAPORT: 'raport',
        SALES: 'sales'
    };

    let screen = SCREENS.RAPORT;

    function toggleScreen() {
        screen = screen === SCREENS.RAPORT ? SCREENS.SALES : SCREENS.RAPORT;
    }
</script>

<div class="HomeScreen">
    <div class="Content">
        <NavBar />
    </div>
    <div>
        <button
            aria-label="Toggle Screen"
            on:click={toggleScreen}>
            Show {screen === SCREENS.RAPORT ? 'Sales' : 'Report'}
        </button>
    </div>
    <div class="Content">
        {#if screen === SCREENS.RAPORT}
            <Raport />
        {:else}
            <Sales />
        {/if}
    </div>
</div>

<style>
    .HomeScreen {
        position: fixed; 
        background-color: #b8b8ba;
        height: 100%;
        width: 100%;
        top: 0px;
        left: 0px;
        z-index: 1000;
    }
</style>
