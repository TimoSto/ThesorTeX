import {Meta, StoryObj} from "@storybook/vue3";
import ImageViewer from "./ImageViewer.vue";
import testimg from "./testimg.png";
import {screen} from "@storybook/testing-library";

const meta: Meta<typeof ImageViewer> = {
    title: "VueComponentLibrary/ImageViewer",
    component: ImageViewer,
};

export default meta;
type Story = StoryObj<typeof ImageViewer>

export const Closed: Story = {
    args: {
        image: testimg
    }
};

export const Opened: Story = {
    args: {
        image: testimg
    },
    play: async () => {
        const t = screen.getByTestId("plainimg");
        t.click();
    }
};
