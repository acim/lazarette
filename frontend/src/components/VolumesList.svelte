<script type="ts">
  import type * as k8s from "@kubernetes/client-node";
  import { get } from "../fetch";
  import { volume } from "../store";
  import type { claim_component } from "svelte/internal";

  interface Volume {
    volume: k8s.V1PersistentVolume;
    claim: k8s.V1PersistentVolumeClaim;
    pods: k8s.V1Pod[];
  }

  interface Volumes {
    classes: k8s.V1StorageClass[];
    volumes: Volume[];
    error: string;
  }

  const promise = get<Volumes>("/volumes");
  function setVolume(vol: k8s.V1PersistentVolume) {
    $volume = vol;
  }
</script>

<style>
  li:hover {
    cursor: pointer;
  }
</style>

<div class="container">
  {#await promise}
    <p>loading...</p>
  {:then response}
    <section>
      <ul>
        {#each response.parsedBody.classes as item}
          <li>{item.metadata.name}</li>
        {/each}
      </ul>
    </section>
    <section>
      <ul>
        {#each response.parsedBody.volumes as item}
          <li on:click={() => setVolume(item.volume)}>
            {item.volume.metadata.name}
          </li>
          <p>PVC: {item.claim.metadata.name}</p>
          {#each item.pods as pod}
            <p>Pod: {pod.metadata.name}</p>
          {/each}
        {/each}
      </ul>
    </section>
  {:catch error}
    <p class="text-error">{error.message}</p>
  {/await}
</div>
