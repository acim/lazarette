<script lang="ts">
  import type * as k8s from "@kubernetes/client-node";

  export let volume: k8s.V1PersistentVolume;
</script>

<ul>
  <li>Storage class: {volume.spec.storageClassName}</li>
  <li>Reclaim policy: {volume.spec.persistentVolumeReclaimPolicy}</li>
  <li>Capacity: {volume.spec.capacity.storage}</li>
  <li>Mode: {volume.spec.accessModes}</li>
  <li>Status: {volume.status.phase}</li>
  <li>Claim kind: {volume.spec.claimRef.kind}</li>
  <li>Claim name: {volume.spec.claimRef.name}</li>
  <li>Claim namespace: {volume.spec.claimRef.namespace}</li>
</ul>
