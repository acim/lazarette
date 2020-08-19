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
    volumes: Volume[];
    error: string;
  }

  const promise = get<Volumes>("/volumes");
  function setVolume(vol: k8s.V1PersistentVolume) {
    $volume = vol;
  }
</script>

<div class="container">
  {#await promise}
    <p>loading...</p>
  {:then response}
    {#each response.parsedBody.volumes as item}
      <section>
        <p>
          <a href="." on:click|preventDefault={() => setVolume(item.volume)}>
            {item.volume.metadata.name}
          </a>
        </p>
        <p>PVC: {item.claim.metadata.name}</p>
        {#each item.pods as pod}
          <p>Pod: {pod.metadata.name}</p>
        {/each}
      </section>
    {/each}
  {:catch error}
    <p class="text-error">{error.message}</p>
  {/await}
</div>
