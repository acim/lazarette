<script type="ts">
  import type crayon from "crayon";
  import Nav from "../components/Nav.svelte";
  import StorageClass from "../components/StorageClass.svelte";
  import Icon from "mdi-svelte";
  import { mdiLoading } from "@mdi/js";
  import store from "../storageClassesStore";
  import { onMount } from "svelte";

  export let nav: crayon.Router;

  const color = getComputedStyle(document.documentElement).getPropertyValue(
    "--color-primary"
  );

  let error: string;

  onMount(async () => {
    try {
      store.load();
    } catch (e) {
      error = e;
    }
  });

  const setDefault = async (name: string) => {
    try {
      store.setDefault(name);
    } catch (e) {
      error = e;
    }
  };
</script>

<Nav {nav} />

<div class="container">
  {#each $store as item (item.metadata.uid)}
    <StorageClass storageClass={item} {setDefault} />
  {:else}
    <Icon path={mdiLoading} size="4rem" spin="2" {color} />
  {/each}
  {#if error}
    <p class="text-error">{error}</p>
  {/if}
</div>
