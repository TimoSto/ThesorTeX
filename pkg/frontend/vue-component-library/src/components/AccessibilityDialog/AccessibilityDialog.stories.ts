import {Meta, StoryObj} from "@storybook/vue3";
import AccessibilityDialog from "./AccessibilityDialog.vue";
import {screen} from "@storybook/testing-library";

const meta: Meta<typeof AccessibilityDialog> = {
    title: "VueComponentLibrary/AccessibilityDialog",
    component: AccessibilityDialog,
    render: (args: any) => ({
        components: {
            AccessibilityDialog,
        },
        setup() {
            return {args};
        },
        template: `
          <AccessibilityDialog v-bind="args" v-slot="scope">
          <v-btn icon @click="scope.openDialog" data-testid="activator">
            <v-icon>mdi-human</v-icon>
          </v-btn>
          </AccessibilityDialog>
        `,
    }),
};

export default meta;
type Story = StoryObj<typeof AccessibilityDialog>;

export const Closed: Story = {
    args: {
        keydownTarget: document,
        i18n: (k: string) => k
    }
};

export const Opened: Story = {
    args: {
        keydownTarget: document,
        i18n: (k: string) => k
    },
    play: async () => {
        const t = screen.getByTestId("activator");
        t.click();
    }
};
