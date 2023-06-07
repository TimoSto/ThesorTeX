import type {Meta, StoryObj} from "@storybook/vue3";

import Button from "./Button.vue";

// More on how to set up stories at: https://storybook.js.org/docs/vue/writing-stories/introduction
const meta = {
    title: "Example/Button",
    component: Button,
    // This component will have an automatically generated docsPage entry: https://storybook.js.org/docs/vue/writing-docs/autodocs
    tags: ["autodocs"],
    argTypes: {
        size: {control: "select", options: ["small", "medium", "large"]},
        backgroundColor: {control: "color"},
        onClick: {action: "clicked"},
    },
    args: {primary: false}, // default value
} satisfies Meta<typeof Button>;

const Template = (args, {argTypes}) => ({
    components: {Button},
    props: Object.keys(argTypes),
    template: `
      <Button v-bind="$props">
      <template v-if="${"test" in args}" #test>${args.test}</template>
      </Button>
    `,
});

export default meta;
type Story = StoryObj<typeof meta>;
/*
 *👇 Render functions are a framework specific feature to allow you control on how the component renders.
 * See https://storybook.js.org/docs/vue/api/csf
 * to learn how to use render functions.
 */
export const Simple = Template.bind({});
Simple.storyName = "Default";
Simple.args = {
    label: "hallo",
    test: "test slot",
};

export const Secondary: Story = {
    args: {
        primary: false,
        label: "Button",
    },
};

export const Large: Story = {
    args: {
        label: "Button",
        size: "large",
    },
};

export const Small: Story = {
    args: {
        label: "Button",
        size: "small",
    },
};
