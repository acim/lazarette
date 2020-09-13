<script lang="ts">
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
  nav {
    text-align: center;
    background-color: var(--color-canvas);
    width: 5rem;
    min-height: 100vh;
  }

  a {
    display: inline-block;
    margin: 1rem 0;
    font-size: 0.8rem;
    line-height: 1rem;
  }

  a,
  a:link,
  a:visited {
    color: var(--color-primary);
    text-decoration: none;
    transition: 350ms color;
  }

  a:focus,
  a:hover,
  a:active {
    color: var(--color-accent);
  }
</style>

<nav>
  <a href="." on:click|preventDefault={() => nav.navigate('/volumes')}>
    <Icon path={mdiNas} size="4rem" {color} />
    Volumes
  </a>
  <a href="." on:click|preventDefault={() => nav.navigate('/classes')}>
    <Icon path={mdiHexagonMultiple} size="4rem" {color} />
    Storage Classes
  </a>
</nav>
