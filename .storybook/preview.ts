import {setup} from "@storybook/vue3";
import {withVuetifyTheme} from "./withVuetifyTheme.decorator";
import "vuetify/styles";
import {createVuetify} from "vuetify";
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";

const vuetify = createVuetify({
    components,
    directives,
});

setup((app) => {
    // Registers your app's plugins into Storybook
    app.use(vuetify);
});

export const decorators = [withVuetifyTheme];
