import {Meta, StoryObj} from "@storybook/vue3";
import PageNavigator from "./PageNavigator.vue";

const meta: Meta<typeof PageNavigator> = {
    title: "VueComponentLibrary/PageNavigator",
    component: PageNavigator,
    render: (args: any) => ({
        components: {
            PageNavigator,
        },
        setup() {
            return {args};
        },
        template: `
          <PageNavigator>
          <template #1>
            <div style="width: 100%; height: 100%; background-color: #0d47a1"></div>
          </template>
          <template v-if="args.numberOfPages>1" #2>
            <div style="width: 100%; height: 100%; background-color: red"></div>
          </template>
          </PageNavigator>
        `,
    }),
};

export default meta;
type Story = StoryObj<typeof PageNavigator>;

export const OnePage: Story = {
    args: {
        numberOfPages: 1
    }
};

export const TwoPages: Story = {
    args: {
        numberOfPages: 2
    }
};