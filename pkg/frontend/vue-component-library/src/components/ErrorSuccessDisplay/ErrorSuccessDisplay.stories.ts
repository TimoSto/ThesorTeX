import ErrorSuccessDisplay from "./ErrorSuccessDisplay.vue";
import {Meta, StoryObj} from "@storybook/vue3";
import ToolbarAndContent from "../ToolbarAndContent/ToolbarAndContent.vue";

const meta: Meta<typeof ErrorSuccessDisplay> = {
    title: "VueComponentLibrary/ErrorSuccessDisplay",
    component: ErrorSuccessDisplay,
    render: (args: any) => ({
        components: {
            ErrorSuccessDisplay,
        },
        setup() {
            return {args};
        },
        template: `
          <ErrorSuccessDisplay v-bind="args">
          <template v-if="${"suffix" in args}" #suffix>${args.suffix}</template>
          </ErrorSuccessDisplay>
        `,
    }),
};

export default meta;
type Story = StoryObj<typeof ToolbarAndContent>;

export const DisplayError: Story = {
    args: {
        valid: false,
        errorTitle: "Some error",
        message: "this is a message",
        suffix: "this suffix",
        close: "close"
    }
};

export const DisplaySuccess: Story = {
    args: {
        valid: true,
        errorTitle: "Some error",
        message: "this is a message",
        suffix: "this suffix",
        close: "close"
    }
};