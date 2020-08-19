import crayon from "crayon";
import transition from "crayon-transition";
import animate from "crayon-animate";
import svelte from "crayon-svelte";
import Base from "./pages/Base.svelte";
import About from "./pages/About.svelte";

const target = document.getElementById("app");
const app = crayon.create();

app.use(svelte.router(target));
app.use(transition.loader());
app.use(
  animate.defaults({
    name: transition.fade,
    duration: 350,
  })
);

app.path("/", (req, res) => req.redirect("/home"));

app.path("/home", (req, res) => req.mount(Base, { req, nav: app }));

app.path("/about", (req, res) => req.mount(About, { req, nav: app }));

app.load();
