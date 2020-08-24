<script type="ts">
  import type crayon from "crayon";
  import Icon from "mdi-svelte";
  import { mdiNas } from "@mdi/js";
  import { mdiHexagonMultiple } from "@mdi/js";

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

  a {
    margin-top: 1rem;
    margin-bottom: 1rem;
  }
</style>

<nav>
  <a href="." on:click|preventDefault={() => nav.navigate('/volumes')}>
    <Icon path={mdiNas} size="4rem" {color} />
    <span>Volumes</span>
  </a>
  <a href="." on:click|preventDefault={() => nav.navigate('/classes')}>
    <Icon path={mdiHexagonMultiple} size="4rem" {color} />
    <span>Storage Classes</span>
  </a>
</nav>
