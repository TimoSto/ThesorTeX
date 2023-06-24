import {Meta, StoryObj} from "@storybook/vue3";
import ApplicationFrame from "./ApplicationFrame.vue";

const meta: Meta<typeof ApplicationFrame> = {
    title: "VueComponentLibrary/ApplicationFrame",
    component: ApplicationFrame,
    render: (args: any) => ({
        components: {
            ApplicationFrame,
        },
        setup() {
            return {args};
        },
        template: `
          <ApplicationFrame></ApplicationFrame>
        `,
    }),
};

export default meta;
type Story = StoryObj<typeof ApplicationFrame>;

export const OnlyTitle: Story = {
    args: {
        documentProp: document,
        mainTitle: "Foobar",
        i18n: (k: string) => k
    }
};
