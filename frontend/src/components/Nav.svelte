<script lang="ts">
  import type crayon from "crayon"
  import { mdiHexagonMultiple, mdiNas } from "@mdi/js"

  export let nav: crayon.Router
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
    line-height: 1.1rem;
    outline: none;
  }

  a svg {
    width: 4rem;
  }

  a svg path {
    fill: var(--color-primary);
    transition: 350ms fill;
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

  a:focus svg path,
  a:hover svg path,
  a:active svg path {
    fill: var(--color-accent);
  }
</style>

<nav>
  <a href="." on:click|preventDefault={() => nav.navigate('/volumes')}>
    <svg viewBox="0 0 24 24">
      <path d={mdiNas} />
    </svg> Volumes
  </a>
  <a href="." on:click|preventDefault={() => nav.navigate('/classes')}>
    <svg viewBox="0 0 24 24">
      <path d={mdiHexagonMultiple} />
    </svg> Storage Classes
  </a>
</nav>
