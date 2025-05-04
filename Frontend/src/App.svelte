<script>
  import NavBar from "./lib/elements/NavBar.svelte";
  import HomePage from "./lib/screens/HomePage.svelte";
  import ReportsPage from "./lib/screens/ReportsPage.svelte";
  
  let currentPage = 'home';
</script>

<main>
  <NavBar on:pageChange={(e) => currentPage = e.detail} />
  <ReportsPage />
</main>

<style>
  main {
    max-width: 1200px;
    margin: 0 auto;
    padding: 1rem;
  }
</style>
