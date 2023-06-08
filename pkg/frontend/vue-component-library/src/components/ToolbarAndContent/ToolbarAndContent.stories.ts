import ToolbarAndContent from "./ToolbarAndContent.vue";
import {Meta, StoryObj} from "@storybook/vue3";
import {screen} from "@storybook/testing-library";

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
          <template v-if="${"content" in args}" #content>${args.content}</template>
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
        `,
        content: `<p>halllooo</p>`
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
        `,
        content: `<p>halllooo</p>`
    }
};

export const WithScrollbar: Story = {
    args: {
        barColor: "primary",
        bar: `
        <v-toolbar-title>Test Title</v-toolbar-title>
        <v-spacer />
        <v-btn color="white">
            btn
        </v-btn>
        `,
        content: `<div style="height: 1500px">hallo</div>`
    }
};

export const WithScrollbarScrolled: Story = {
    args: {
        barColor: "primary",
        bar: `
        <v-toolbar-title>Test Title</v-toolbar-title>
        <v-spacer />
        <v-btn color="white">
            btn
        </v-btn>
        `,
        content: `<div style="height: 1500px">hallo</div>`
    },
    play: async () => {
        const t = screen.getByText("hallo", {
            selector: "div"
        });
        console.log("hi", t);
        t.parentElement!.scrollTo(0, 300);
    }
};
