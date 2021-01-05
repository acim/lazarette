<script lang="ts">
  import { mdiHexagonMultiple, mdiNas } from '@mdi/js'
  import { Link } from '@bjornlu/svelte-router'
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
