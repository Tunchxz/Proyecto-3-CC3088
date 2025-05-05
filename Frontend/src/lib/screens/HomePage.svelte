<script>
    import NavBar from "../Components/NavBar.svelte";
    import Contracts from "../raports/Contracts.svelte";
    import Payments from "../raports/Payments.svelte";
    import Reservations from "../raports/reservations.svelte";
    import Maintenance from "../raports/Maintenance.svelte";
    import Fines from "../raports/Fines.svelte";

    const SCREENS = {
        RESERVATIONS: 'reservations',
        CONTRACTS: 'contracts',
        PAYMENTS: 'payments',
        MAINTENANCE: 'maintenance',
        FINES: 'fines'
    };

    let currentScreen = SCREENS.RESERVATIONS;

    function switchScreen(screen) {
        currentScreen = screen;
    }
</script>

<div class="HomeScreen">
    <div class="Content">
        <NavBar />
    </div>
    <div class="button-container">
        <button
            class:active={currentScreen === SCREENS.RESERVATIONS}
            on:click={() => switchScreen(SCREENS.RESERVATIONS)}>
            Reservaciones
        </button>
        <button
            class:active={currentScreen === SCREENS.CONTRACTS}
            on:click={() => switchScreen(SCREENS.CONTRACTS)}>
            Contratos
        </button>
        <button
            class:active={currentScreen === SCREENS.PAYMENTS}
            on:click={() => switchScreen(SCREENS.PAYMENTS)}>
            Pagos
        </button>
        <button
            class:active={currentScreen === SCREENS.MAINTENANCE}
            on:click={() => switchScreen(SCREENS.MAINTENANCE)}>
            Mantenimientos
        </button>
        <button
            class:active={currentScreen === SCREENS.FINES}
            on:click={() => switchScreen(SCREENS.FINES)}>
            Multas
        </button>
    </div>
    <div class="Content">
        {#if currentScreen === SCREENS.RESERVATIONS}
            <Reservations />
        {:else if currentScreen === SCREENS.CONTRACTS}
            <Contracts />
        {:else if currentScreen === SCREENS.PAYMENTS}
            <Payments />
        {:else if currentScreen === SCREENS.MAINTENANCE}
            <Maintenance />
        {:else if currentScreen === SCREENS.FINES}
            <Fines />
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

    .button-container {
        display: flex;
        justify-content: center;
        gap: 1rem;
        margin: 1rem 0;
    }

    button {
        padding: 0.5rem 1rem;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        background-color: #919191;
        transition: background-color 0.2s;
    }

    button:hover {
        background-color: #464646;
    }

    button.active {
        background-color: #2b2b2b;
        color: white;
    }
</style>
