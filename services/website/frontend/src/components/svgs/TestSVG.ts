import {SVGPartial, TemplateSVG} from "@thesortex/vue-component-library/src/components/SVGTemplate/helper/SVG";

const p1: SVGPartial = {
    scale: 0.7,
    angle: 90,
    translate: {
        x: 50,
        y: 20,
    },
    parts: [
        {
            vector: {
                x: 0,
                y: 0
            }
        },
        {
            vector: {
                x: -200,
                y: 0
            }
        },
        {
            arc: {
                radius: 50,
                rotation: 0,
                clockwise: true,
                end: {
                    x: -250,
                    y: 50
                }
            }
        }
    ],
    strokeColor: "rgba(var(--v-theme-on-background), 1)",
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
