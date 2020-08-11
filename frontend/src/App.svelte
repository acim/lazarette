<script type="ts">
  import type * as k8s from "@kubernetes/client-node";

  async function fetchData() {
    const res = await fetch("/volumes");
    const data: {
      volumes: k8s.V1PersistentVolume[];
    } = await res.json();

    if (res.ok) {
      return data;
    } else {
      throw new Error();
    }
  }
</script>

<h1>kubernetes Volumes Explorer</h1>
{#await fetchData()}
  <p>loading</p>
{:then data}
  {#each data.volumes as volume}
    <li>{volume.metadata.name}</li>
  {/each}
{:catch error}
  <p style="color: red">{error.message}</p>
{/await}
