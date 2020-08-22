<script type="ts">
  import type * as k8s from "@kubernetes/client-node";
  import { get } from "../fetch";
  import { volume } from "../store";
  import Icon from "mdi-svelte";
  import { mdiLoading } from "@mdi/js";

  interface Volume {
    volume: k8s.V1PersistentVolume;
    claim: k8s.V1PersistentVolumeClaim;
    pods: k8s.V1Pod[];
  }

  interface Volumes {
    volumes: Volume[];
    error: string;
  }

  const promise = get<Volumes>("/v1/volumes.json");
  function setVolume(vol: k8s.V1PersistentVolume) {
    $volume = vol;
  }

  const color = getComputedStyle(document.documentElement).getPropertyValue(
    "--color-primary"
  );
</script>

<div class="container">
  {#await promise}
    <Icon path={mdiLoading} size="4rem" spin="2" {color} />
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
