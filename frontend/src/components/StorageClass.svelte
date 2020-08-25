<script type="ts">
  import type { V1StorageClass } from "@kubernetes/client-node";
  import { fade } from "svelte/transition";
  import store from "../storageClassesStore";
  import Icon from "mdi-svelte";
  import { mdiLoading } from "@mdi/js";

  export let i: number;

  let loading = false;

  const color = getComputedStyle(document.documentElement).getPropertyValue(
    "--color-primary"
  );

  const setDefault = (name: string) => {
    try {
      store.setDefault(name);
    } catch (err) {
      console.log(err);
    }
  };

  let isDefault: () => boolean;
  $: isDefault = () => {
    return (
      $store[i].metadata.annotations.hasOwnProperty(
        "storageclass.kubernetes.io/is-default-class"
      ) &&
      $store[i].metadata.annotations[
        "storageclass.kubernetes.io/is-default-class"
      ] === "true"
    );
  };
</script>

<style>
  table {
    margin-bottom: 1rem;
  }
  button {
    margin-left: auto;
  }
</style>

<section class:position-relative={isDefault()} transition:fade>
  <h3>{$store[i].metadata.name}</h3>
  {#if isDefault()}
    <tag>default</tag>
  {/if}
  <table>
    <tr>
      <td>Provisioner</td>
      <td>{$store[i].provisioner}</td>
    </tr>
    <tr>
      <td>Reclaim policy</td>
      <td>{$store[i].reclaimPolicy}</td>
    </tr>
    <tr>
      <td>Allow expansion</td>
      <td>{$store[i].allowVolumeExpansion}</td>
    </tr>
    <tr>
      <td>Binding mode</td>
      <td>{$store[i].volumeBindingMode}</td>
    </tr>
  </table>
  {#if !isDefault()}
    <button
      on:click|once={() => {
        loading = true;
        setDefault($store[i].metadata.name);
      }}>
      {#if loading}
        <Icon path={mdiLoading} spin="2" {color} />
      {:else}Set to default{/if}
    </button>
  {/if}
</section>
