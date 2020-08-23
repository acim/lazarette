<script type="ts">
  import type crayon from "crayon";
  import Nav from "../components/Nav.svelte";
  import StorageClass from "../components/StorageClass.svelte";
  import Icon from "mdi-svelte";
  import { mdiLoading, mdiCarHatchback } from "@mdi/js";
  import { storageClasses, loadStorageClasses } from "../store";
  import { onMount } from "svelte";

  export let req: crayon.Context;
  export let nav: crayon.Router;

  const color = getComputedStyle(document.documentElement).getPropertyValue(
    "--color-primary"
  );

  let error: string;

  onMount(async () => {
    try {
      await loadStorageClasses();
    } catch (e) {
      error = e;
    }
  });
</script>

<Nav {req} {nav} />

<div class="container">
  {#each $storageClasses as item}
    <StorageClass storageClass={item} />
  {:else}
    <Icon path={mdiLoading} size="4rem" spin="2" {color} />
  {/each}
  {#if error}
    <p class="text-error">{error}</p>
  {/if}
</div>
