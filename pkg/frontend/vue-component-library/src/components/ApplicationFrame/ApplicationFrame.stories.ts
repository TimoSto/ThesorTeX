import {Meta, StoryObj} from "@storybook/vue3";
import ApplicationFrame from "./ApplicationFrame.vue";
import {screen} from "@storybook/testing-library";

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
          <ApplicationFrame>
          <template v-if="${"sidebarContent" in args}" #sidebar>${args.sidebarContent}</template>
          </ApplicationFrame>
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
        i18n: (k: string) => k,
        sidebarContent: "<h2>hello foobar sidebar</h2>"
    }
};

export const WithTitleAndSidebarOpened: Story = {
    args: {
        documentProp: document,
        mainTitle: "Foobar",
        hasSidebar: true,
        i18n: (k: string) => k,
        sidebarContent: "<h2>hello foobar sidebar</h2>"
    },
    play: async () => {
        const t = screen.getByTitle("ApplicationFrame.OpenSidebar");
        t.click();
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

export const WithTitleDocumentationAndConfig: Story = {
    args: {
        documentProp: document,
        mainTitle: "Foobar",
        documentationTarget: "test",
        hasConfig: true,
        i18n: (k: string) => k
    }
};

export const WithFullAppBar: Story = {
    args: {
        hasSidebar: true,
        documentProp: document,
        mainTitle: "Foobar",
        documentationTarget: "test",
        showA11y: true,
        hasConfig: true,
        i18n: (k: string) => k
    }
};
