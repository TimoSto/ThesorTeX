import {App} from "vue";
import * as components from "./components";
import {de as accessibilityDialog_de} from "./components/AccessibilityDialog/i18n/de";
import {en as accessibilityDialog_en} from "./components/AccessibilityDialog/i18n/en";

export function installLib(app: App) {
    for (const key in components) {
        // @ts-expect-error
        app.component(key, components[key]);
    }
}

export const accessibilityDialogDe = accessibilityDialog_de;
export const accessibilityDialogEn = accessibilityDialog_en;
