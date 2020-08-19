<script type="ts">
  export let req;
  export let nav;
  import VolumesList from "../components/VolumesList.svelte";
  import Header from "../components/Header.svelte";
  import Volume from "../components/Volume.svelte";
</script>

<style>
  main {
    display: flex;
  }
</style>

<nav>
  <button on:click={() => nav.navigate('/home')}>Home</button>
  <button on:click={() => nav.navigate('/about')}>About</button>
  <button class="anchored" on:click={() => nav.navigate('/more')}>
    More &gt;
  </button>
</nav>
<h1>{req.routePattern}</h1>
<Header />
<main>
  <VolumesList />
  <Volume />
</main>
