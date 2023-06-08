import SVGTemplate from "./SVGTemplate.vue";
import {Meta, StoryObj} from "@storybook/vue3";
import {PathPart, SVGPartial, TemplateSVG} from "./helper/SVG";

const meta: Meta<typeof SVGTemplate> = {
    title: "VueComponentLibrary/SVGTemplate",
    component: SVGTemplate,
};

export default meta;
type Story = StoryObj<typeof SVGTemplate>;

const testrect: PathPart[] = [
    {
        vector: {
            x: 20,
            y: 20,
        }
    },
    {
        vector: {
            x: 180,
            y: 20,
        }
    },
    {
        vector: {
            x: 180,
            y: 180,
        }
    },
    {
        vector: {
            x: 20,
            y: 180,
        }
    },
];

const testpart: SVGPartial = {
    parts: [testrect],
    strokeColor: "red",
    strokeWidth: "4px",
    fillColor: "green",
};

const testsvg: TemplateSVG = {
    partials: [testpart],
    viewBox: "0 0 200 200",
    width: 200,
    height: 200,
};

export const Simple: Story = {
    args: {
        maxHeight: 400,
        maxWidth: 600,
        svg: testsvg
    }
};
