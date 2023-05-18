import {setup} from "@storybook/vue3";
import {vuetifyInstance} from "@thesortex/vuetify-plugin";
import {withVuetifyTheme} from "./withVuetifyTheme.decorator";

setup((app) => {
    // Registers your app's plugins into Storybook
    app.use(vuetifyInstance);
});

export const decorators = [withVuetifyTheme];
