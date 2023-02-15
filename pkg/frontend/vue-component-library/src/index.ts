import {App} from "vue";
import * as components from "./components";

export function installLib(app: App) {
    for (const key in components) {
        // @ts-expect-error
        app.component(key, components[key]);
    }
}
