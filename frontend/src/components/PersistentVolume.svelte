<script lang="ts">
  import { fade } from "svelte/transition";
  import type { PersistentVolume } from "../persistentVolumesStore";
  import store from "../persistentVolumesStore";

  export let i: number;

  // let isDefault: () => boolean;
  // $: isDefault = () => {
  //   return (
  //     $store[i].metadata.annotations.hasOwnProperty(
  //       "storageclass.kubernetes.io/is-default-class"
  //     ) &&
  //     $store[i].metadata.annotations[
  //       "storageclass.kubernetes.io/is-default-class"
  //     ] === "true"
  //   );
  // };
</script>

<section transition:fade>
  <h3>{$store[i].volume.metadata.name}</h3>
  <table>
    <tr>
      <td>Storage class</td>
      <td>{$store[i].volume.spec.storageClassName}</td>
    </tr>
    <tr>
      <td>Capacity</td>
      <td>{$store[i].volume.spec.capacity.storage}</td>
    </tr>
    <tr>
      <td>Mode</td>
      <td>{$store[i].volume.spec.accessModes}</td>
    </tr>
    <tr>
      <td>Reclaim policy</td>
      <td>{$store[i].volume.spec.persistentVolumeReclaimPolicy}</td>
    </tr>
    <tr>
      <td>Status</td>
      <td>{$store[i].volume.status.phase}</td>
    </tr>
    <tr>
      <td>Reference claim kind</td>
      <td>{$store[i].volume.spec.claimRef.kind}</td>
    </tr>
    <tr>
      <td>Referencing claim name</td>
      <td>
        {$store[i].volume.spec.claimRef.namespace}/{$store[i].volume.spec.claimRef.name}
      </td>
    </tr>
    <tr>
      <td>Associated claim name</td>
      <td>
        {$store[i].claim.metadata.namespace}/{$store[i].claim.metadata.name}
      </td>
    </tr>
    <tr>
      <td>Associated claim capacity</td>
      <td>{$store[i].claim.status.capacity.storage}</td>
    </tr>
    <tr>
      <td>Associated claim modes</td>
      <td>{$store[i].claim.status.accessModes}</td>
    </tr>
    <tr>
      <td>Associated claim status</td>
      <td>{$store[i].claim.status.phase}</td>
    </tr>
    <tr>
      {#each $store[i].pods as pod, i (pod.metadata.uid)}
        <td>Mounted by pod</td>
        <td>{$store[i].pods[i].metadata.name}</td>
      {/each}
    </tr>
  </table>
  <!-- {#if !isDefault()}
    <button
      on:click|once={() => {
        setDefault($store[i].metadata.name);
      }}>
      Set to default
    </button>
  {/if} -->
</section>
