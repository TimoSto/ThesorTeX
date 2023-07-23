import {App} from "vue";
import * as components from "./components";
import {de as accessibilityDialog_de} from "./components/AccessibilityDialog/i18n/de";
import {en as accessibilityDialog_en} from "./components/AccessibilityDialog/i18n/en";
import {i18nKeys as accessibilityDialog_keys} from "./components/AccessibilityDialog/i18n/keys";
import {de as applicationFrame_de} from "./components/ApplicationFrame/i18n/de";
import {en as applicationFrame_en} from "./components/ApplicationFrame/i18n/en";
import {i18nKeys as applicationFrame_keys} from "./components/ApplicationFrame/i18n/keys";

export function installLib(app: App) {
    for (const key in components) {
        // @ts-expect-error
        app.component(key, components[key]);
    }
}

export const accessibilityDialogDe = accessibilityDialog_de;
export const accessibilityDialogEn = accessibilityDialog_en;
export const accessibilityDialogKeys = accessibilityDialog_keys;

export const applicationFrameDe = applicationFrame_de;
export const applicationFrameEn = applicationFrame_en;
export const applicationFrameKeys = applicationFrame_keys;
