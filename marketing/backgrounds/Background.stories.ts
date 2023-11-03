import {Meta, StoryObj} from "@storybook/vue3";
import Background from "./Background.vue";

const meta: Meta<typeof Background> = {
    title: "Background",
    component: Background,
};

export default meta;

type Story = StoryObj<typeof Background>

export const Plain: Story = {
    args: {
        // props
    }
};
