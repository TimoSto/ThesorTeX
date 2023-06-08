import ToolbarAndContent from "./ToolbarAndContent.vue";
import {Meta, StoryObj} from "@storybook/vue3";

const meta: Meta<typeof ToolbarAndContent> = {
    title: "VueComponentLibrary/ToolbarAndContent",
    component: ToolbarAndContent,
    render: (args: any) => ({
        components: {
            ToolbarAndContent,
        },
        setup() {
            return {args};
        },
        template: `
          <ToolbarAndContent v-bind="args">
          <template v-if="${"bar" in args}" #bar>${args.bar}</template>
          </ToolbarAndContent>
        `,
    }),
};

export default meta;
type Story = StoryObj<typeof ToolbarAndContent>;

export const Simple: Story = {
    args: {
        barColor: "background",
        bar: `
        <v-toolbar-title>Test Title</v-toolbar-title>
        <v-spacer />
        <v-btn color="primary">
            btn
        </v-btn>
        `
    }
};

export const Primary: Story = {
    args: {
        barColor: "primary",
        bar: `
        <v-toolbar-title>Test Title</v-toolbar-title>
        <v-spacer />
        <v-btn color="white">
            btn
        </v-btn>
        `
    }
};
