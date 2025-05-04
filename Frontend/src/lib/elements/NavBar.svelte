<script>
  import { createEventDispatcher } from 'svelte';
  
  const dispatch = createEventDispatcher();
  
  const handlePageChange = (page) => {
    dispatch('pageChange', page);
  };
</script>

<div>
    <nav class="navbar">
        <div class="navbar__brand">
        <a href="/">AutoVentas</a>
        </div>
        <ul class="navbar__menu">
        <li><a href="/">Home</a></li>
        <li><a href="/about">About</a></li>
        <li><a href="/services">Services</a></li>
        </ul>
    </nav>
    <style>
        .navbar {
        display: flex;
        justify-content: space-around;
        align-items: center;
        color: white;
        padding: 1rem;        
        }
        .navbar__brand a {
        color: white;
        text-decoration: none;
        font-size: 1.5rem;
        }
        .navbar__menu {
        list-style-type: none;
        display: flex;
        gap: 1rem;
        }
        .navbar__menu li a {
        color: white;
        text-decoration: none;
        }
    </style>
</div>

<nav class="navbar">
  <div class="nav-brand">
    <h1>Sistema de Reportes</h1>
  </div>
  
  <div class="nav-links">
    <button class="nav-link" on:click={() => handlePageChange('home')}>
      Inicio
    </button>
    <button class="nav-link" on:click={() => handlePageChange('reports')}>
      Reportes
    </button>
  </div>
</nav>

<style>
  .navbar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1rem 2rem;
    background-color: #2c3e50;
    color: white;
  }

  .nav-brand h1 {
    margin: 0;
    font-size: 1.5rem;
  }

  .nav-links {
    display: flex;
    gap: 1rem;
  }

  .nav-link {
    background: none;
    border: none;
    color: white;
    padding: 0.5rem 1rem;
    cursor: pointer;
    font-size: 1rem;
    border-radius: 4px;
    transition: background-color 0.3s;
  }

  .nav-link:hover {
    background-color: #34495e;
  }
</style>