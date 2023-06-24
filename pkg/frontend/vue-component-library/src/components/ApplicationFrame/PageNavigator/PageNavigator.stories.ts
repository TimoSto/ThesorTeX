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
          <PageNavigator v-bind="args">
          <template #0="{openPage}">
            <div style="width: 100%; height: 100%; background-color: #0d47a1">
              <button @click="openPage('test')">hallo</button>
            </div>
          </template>
          <template #1="{openPage}">
            <div style="width: 100%; height: 100%; background-color: #a10d4b">
              <button @click="openPage('test')">hallo</button>
            </div>
          </template>
          </PageNavigator>
        `,
    }),
};

export default meta;
type Story = StoryObj<typeof PageNavigator>;

export const OnePage: Story = {
    args: {}
};
