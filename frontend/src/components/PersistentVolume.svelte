<script lang="ts">
  import { fade } from "svelte/transition";
  import type { PersistentVolume } from "../persistentVolumesStore";

  export let persistentVolume: PersistentVolume;
</script>

<style>
  ul {
    margin-top: 2rem;
  }
</style>

<ul transition:fade>
  {#if persistentVolume}
    <li>Storage class: {persistentVolume.volume.spec.storageClassName}</li>
    <li>
      Reclaim policy: {persistentVolume.volume.spec.persistentVolumeReclaimPolicy}
    </li>
    <li>Capacity: {persistentVolume.volume.spec.capacity.storage}</li>
    <li>Mode: {persistentVolume.volume.spec.accessModes}</li>
    <li>Status: {persistentVolume.volume.status.phase}</li>
    <li>Claim kind: {persistentVolume.volume.spec.claimRef.kind}</li>
    <li>Claim name: {persistentVolume.volume.spec.claimRef.name}</li>
    <li>Claim namespace: {persistentVolume.volume.spec.claimRef.namespace}</li>
  {/if}
</ul>
