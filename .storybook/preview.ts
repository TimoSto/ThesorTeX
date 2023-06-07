import {setup} from "@storybook/vue3";
import {withVuetifyTheme} from "./withVuetifyTheme.decorator";
import "vuetify/styles";
import {createVuetify} from "vuetify";
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";

const vuetify = createVuetify({
    components,
    directives,
    theme: {
        defaultTheme: "light",
        themes: {
            light: {
                colors: {
                    primary: "#008833",
                    primaryText: "#FFFFFF",
                    secondary: "#52634f",
                    accent: "#000000",
                    error: "#ba1a1a",
                    background: "#FFFFFF",
                    primaryDark: "#15a835",
                }
            },
        }
    }
});

setup((app) => {
    // Registers your app's plugins into Storybook
    app.use(vuetify);
});

export const decorators = [withVuetifyTheme];
