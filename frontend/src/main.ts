import crayon from "crayon";
import transition from "crayon-transition";
import animate from "crayon-animate";
import svelte from "crayon-svelte";
import Volumes from "./pages/Volumes.svelte";
import Classes from "./pages/Classes.svelte";

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

nav.path("/volumes", (req, res) => req.mount(Volumes, { req, nav }));

nav.path("/classes", (req, res) => req.mount(Classes, { req, nav }));

nav.load();
