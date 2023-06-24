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

export const WithTitleAndSidebar: Story = {
    args: {
        documentProp: document,
        mainTitle: "Foobar",
        hasSidebar: true,
        i18n: (k: string) => k
    }
};

export const WithTitleAndDocumentationLink: Story = {
    args: {
        documentProp: document,
        mainTitle: "Foobar",
        documentationTarget: "test",
        i18n: (k: string) => k
    }
};

export const WithTitleDocumentationAndA11y: Story = {
    args: {
        documentProp: document,
        mainTitle: "Foobar",
        documentationTarget: "test",
        showA11y: true,
        i18n: (k: string) => k
    }
};
