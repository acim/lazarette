import crayon from "crayon";
import transition from "crayon-transition";
import animate from "crayon-animate";
import svelte from "crayon-svelte";
import PersistentVolumes from "./pages/PersistentVolumes.svelte";
import StorageClasses from "./pages/StorageClasses.svelte";

const target = document.getElementById("app");
const nav = crayon.create();

nav.use(svelte.router(target));
nav.use(transition.loader());
nav.use(
  animate.defaults({
    name: transition.fade,
    duration: 350,
  })
);

nav.path("/", (req, res) => req.redirect("/volumes"));

nav.path("/volumes", (req, res) => req.mount(PersistentVolumes, { nav }));

nav.path("/classes", (req, res) => req.mount(StorageClasses, { nav, req }));

nav.load();
