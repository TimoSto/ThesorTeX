import {SVGPartial, TemplateSVG} from "@thesortex/vue-component-library/src/components/SVGTemplate/helper/SVG";

const p1: SVGPartial = {
    scale: 1,
    angle: 0,
    parts: [
        {
            vector: {
                x: 0,
                y: 0
            }
        },
        {
            vector: {
                x: 200,
                y: 0
            }
        }
    ],
    strokeColor: "green",
    strokeWidth: "10",
    fillColor: "white"
};

export const TestSVG: TemplateSVG = {
    viewBox: "-250 -250 500 500",
    width: 500,
    height: 500,
    partials: [
        p1,
    ]
};
