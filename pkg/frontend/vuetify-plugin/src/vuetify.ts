/**
 * plugins/vuetify.ts
 *
 * Framework documentation: https://vuetifyjs.com`
 */

// Styles
import "@mdi/font/css/materialdesignicons.css";
import "vuetify/styles";

// Composables
import {createVuetify} from "vuetify";

// https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
export function GetVuetifyInstance(dark: boolean) {
    console.log(dark);
    return createVuetify({
        theme: {
            defaultTheme: dark ? "dark" : "light",
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
                dark: {
                    dark: true,
                    colors: {
                        primary: "#15a835",
                        secondary: "#b9ccb4",
                        accent: "#FFFFFF",
                        error: "#CF6679",
                        background: "#001220",
                        primaryDark: "#15a835",
                    }
                },
                contrastLight: {
                    dark: false,
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
                contrastDark: {
                    dark: true,
                    colors: {
                        primary: "#15a835",
                        secondary: "#b9ccb4",
                        accent: "#FFFFFF",
                        error: "#CF6679",
                        background: "#001220",
                        primaryDark: "#15a835",
                    }
                }
            },
        },
    });
}
