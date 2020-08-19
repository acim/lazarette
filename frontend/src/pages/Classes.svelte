<script type="ts">
  import type crayon from "crayon";
  import Nav from "../components/Nav.svelte";

  export let req: crayon.Context;
  export let nav: crayon.Router;
</script>

<Nav {nav} />
Storage Classes
