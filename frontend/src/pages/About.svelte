<script type="ts">
  import type crayon from "crayon";
  import Header from "../components/Header.svelte";

  // export let req: crayon.Context;
  export let nav: crayon.Router;
</script>

<Header />
<nav>
  <button on:click={() => nav.navigate('/home')}>Home</button>
  <button on:click={() => nav.navigate('/about')}>About</button>
</nav>
