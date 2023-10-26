import {Meta, StoryObj} from "@storybook/vue3";
import Mascot from "./Mascot.vue";

const meta: Meta<typeof Mascot> = {
    title: "Mascot",
    component: Mascot,
};

export default meta;

type Story = StoryObj<typeof Mascot>

export const Simple: Story = {
    args: {
        // props
    }
};

export const Right: Story = {
    args: {
        // props
        right: true,
    }
};
