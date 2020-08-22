<script type="ts">
  import type crayon from "crayon";
  import Nav from "../components/Nav.svelte";
  import VolumesList from "../components/VolumesList.svelte";
  import Volume from "../components/Volume.svelte";

  export let req: crayon.Context;
  export let nav: crayon.Router;
</script>

<Nav {req} {nav} />
<VolumesList />
<Volume />
