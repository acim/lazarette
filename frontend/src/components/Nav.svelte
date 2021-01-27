<script lang="ts">
  import { mdiHexagonMultiple, mdiNas } from '@mdi/js'
  import { Link } from '@bjornlu/svelte-router'
</script>

<nav>
  <Link to="/">
    <svg viewBox="0 0 24 24">
      <path d={mdiNas} />
    </svg>
    Volumes
  </Link>
  <Link to="/classes">
    <svg viewBox="0 0 24 24">
      <path d={mdiHexagonMultiple} />
    </svg>
    Storage Classes
  </Link>
</nav>

<style>
  nav {
    text-align: center;
    background-color: var(--color-canvas);
    width: 5rem;
    min-height: 100vh;
  }

  a svg path {
    fill: var(--color-primary);
    transition: 350ms fill;
  }

  a:focus svg path,
  a:hover svg path,
  a:active svg path {
    fill: var(--color-accent);
  }
</style>
