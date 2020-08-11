<script type="ts">
  import VolumesList from "./components/VolumesList.svelte";
  import Header from "./components/Header.svelte";
  import Volume from "./components/Volume.svelte";
</script>

<style>
  main {
    display: flex;
  }
</style>

<Header />
<main>
  <VolumesList />
  <Volume />
</main>
