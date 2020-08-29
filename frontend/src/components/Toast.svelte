<script type="ts">
  import store from "../toastStore";
</script>

<style>
  div {
    position: absolute;
    top: 5rem;
    margin: 0 auto;
    display: none;
    transition: display 350ms;
    background-color: var(--color-accent);
    border-radius: 0.5rem;
    width: 40ch;
    padding: 1rem;
    left: 50%;
    transform: translate(-50%, -50%);
  }
  .show {
    display: block;
  }
</style>

<div class:show={$store.message}>{$store.message}</div>
