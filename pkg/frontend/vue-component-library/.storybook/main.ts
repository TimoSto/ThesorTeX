const path = require("path");
module.exports = {
    stories: ["../stories/**/*.stories.ts"],
    addons: [
        "@storybook/addon-links",
        "@storybook/addon-essentials",
        "@storybook/addon-interactions",
    ],
    framework: {
        name: "@storybook/vue3-vite",
        options: {},
    },
    docs: {
        docsPage: true,
    },
    typescript: {
        check: false,
        checkOptions: {},
    },
};