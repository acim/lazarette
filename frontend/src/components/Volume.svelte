<script lang="ts">
  import { fade } from "svelte/transition";
  import { volume } from "../store";
</script>

<style>
  ul {
    margin-top: 2rem;
  }
</style>

<ul transition:fade>
  {#if $volume}
    <li>Storage class: {$volume.spec.storageClassName}</li>
    <li>Reclaim policy: {$volume.spec.persistentVolumeReclaimPolicy}</li>
    <li>Capacity: {$volume.spec.capacity.storage}</li>
    <li>Mode: {$volume.spec.accessModes}</li>
    <li>Status: {$volume.status.phase}</li>
    <li>Claim kind: {$volume.spec.claimRef.kind}</li>
    <li>Claim name: {$volume.spec.claimRef.name}</li>
    <li>Claim namespace: {$volume.spec.claimRef.namespace}</li>
  {:else}
    <li>Click on the PersistentVolume</li>
  {/if}
</ul>
