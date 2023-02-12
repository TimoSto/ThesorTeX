import SVGTemplate from "./SVGTemplate.vue";
import {App} from "vue";

export function installLib(app: App) {
    app.component("SVGTemplate", SVGTemplate);
}