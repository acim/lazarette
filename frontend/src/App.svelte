{#await fetchData()}
  <p>loading</p>
{:then data}
  {#each data.data.volumes as volume}
    <li>{volume.metadata.name}</li>
  {/each}
{:catch error}
  <p style="color: red">{error.message}</p>
{/await}

<script type="ts">
  async function fetchData() {
    const res = await fetch("/volumes");
    const data = await res.json();

    if (res.ok) {
      return data;
    } else {
      throw new Error(data);
    }
  }
</script>
