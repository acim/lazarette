<script type="ts">
  import type crayon from "crayon";
  import Icon from "mdi-svelte";
  import { mdiHarddisk } from "@mdi/js";
  import { mdiNas } from "@mdi/js";

  //   export let req: crayon.Context;
  export let nav: crayon.Router;

  const color = getComputedStyle(document.documentElement).getPropertyValue(
    "--color-primary"
  );
</script>

<style>
  span {
    display: block;
    color: var(--color-primary);
    font-size: 0.8rem;
    line-height: 1rem;
  }
  hr {
    color: var(--color-primary);
    width: 80%;
    margin: 1rem auto;
    margin-top: 1.5rem;
  }
</style>

<nav>
  <a href="." on:click|preventDefault={() => nav.navigate('/volumes')}>
    <Icon path={mdiHarddisk} size="4rem" {color} />
    <span>Volumes</span>
  </a>
  <hr />
  <a href="." on:click|preventDefault={() => nav.navigate('/classes')}>
    <Icon path={mdiNas} size="4rem" {color} />
    <span>Storage Classes</span>
  </a>
</nav>
