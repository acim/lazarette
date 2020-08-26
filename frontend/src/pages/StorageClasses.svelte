<script type="ts">
  import type crayon from "crayon";
  import Nav from "../components/Nav.svelte";
  import StorageClass from "../components/StorageClass.svelte";
  import Icon from "mdi-svelte";
  import { mdiLoading } from "@mdi/js";
  import store from "../storageClassesStore";
  import { onMount } from "svelte";
  import Toast from "../components/Toast.svelte";

  export let nav: crayon.Router;

  const color = getComputedStyle(document.documentElement).getPropertyValue(
    "--color-primary"
  );

  let error: string;

  onMount(() => {
    try {
      store.load();
    } catch (err) {
      error = err;
    }
  });
</script>

<Nav {nav} />

<div class="container">
  {#each $store as item, i (item.metadata.uid)}
    <StorageClass {i} />
  {:else}
    <Icon path={mdiLoading} size="4rem" spin="2" {color} />
  {/each}
  {#if error}
    <p class="text-error">{error}</p>
  {/if}
</div>

<Toast />
