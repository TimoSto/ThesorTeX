import Button from "./Button.vue";
import {Meta, StoryObj} from "@storybook/vue3";

const meta: Meta<typeof Button> = {
    title: "Example/Button",
    component: Button,
    render: (args: any) => ({
        components: {
            Button,
        },
        setup() {
            return {args};
        },
        template: `
          <Button v-bind="args">
          <template v-if="${"test" in args}" #test>${args.test}</template>
          </Button>
        `,
    }),
};

export default meta;
type Story = StoryObj<typeof Button>;

export const Default: Story = {
    args: {
        label: "hallo",
        test: "test slot",
    }
};