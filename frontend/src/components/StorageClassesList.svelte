<script type="ts">
  import StorageClass from "./StorageClass.svelte";
  import Icon from "mdi-svelte";
  import { mdiLoading } from "@mdi/js";
  import { storageClasses, loadStorageClasses } from "../store";
  import { onMount } from "svelte";

  const color = getComputedStyle(document.documentElement).getPropertyValue(
    "--color-primary"
  );

  onMount(async () => {
    await loadStorageClasses();
  });
</script>

<div class="container">
  {#each $storageClasses.classes as item}
    <StorageClass storageClass={item} />
  {:else}
    <Icon path={mdiLoading} size="4rem" spin="2" {color} />
  {/each}
  {#if $storageClasses.error !== ''}
    <p class="text-error">$storageClasses.err}</p>
  {/if}
</div>
