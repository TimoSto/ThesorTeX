import {
    PathPart,
    SVGPartial,
    TemplateSVG
} from "@thesortex/vue-component-library/src/components/SVGTemplate/helper/SVG";

const outerBorder: PathPart[] = [
    {
        vector: {
            x: -140,
            y: -200
        }
    },
    {
        vector: {
            x: 140,
            y: -200
        }
    },
    {
        vector: {
            x: 140,
            y: 200
        },
    },
    {
        vector: {
            x: -140,
            y: 200
        }
    },
    {
        arc: {
            radius: 40,
            rotation: 0,
            clockwise: false,
            end: {
                x: -180,
                y: 160
            },
        }
    },
    {
        vector: {
            x: -180,
            y: -160
        }
    },
    {
        arc: {
            radius: 40,
            rotation: 0,
            clockwise: false,
            end: {
                x: -140,
                y: -200
            }
        }
    },
];

const bookMArk: PathPart[] = [
    {
        vector: {
            x: 40,
            y: -200
        },
    },
    {
        vector: {
            x: 40,
            y: -120
        },
    },
    {
        vector: {
            x: 70,
            y: -150
        },
    },
    {
        vector: {
            x: 100,
            y: -120
        },
    },
    {
        vector: {
            x: 100,
            y: -200
        },
    }
];

const p1: SVGPartial = {
    strokeColor: "rgba(var(--v-theme-on-background), 1)",
    strokeWidth: "10",
    fillColor: "white",
    scale: 1,
    angle: -10,
    translate: {
        x: 0,
        y: 0,
    },
    parts: [
        outerBorder,
        bookMArk
    ]
};

export const TestSVG: TemplateSVG = {
    viewBox: "-250 -250 500 500",
    width: 500,
    height: 500,
    partials: [
        p1,
    ]
};
