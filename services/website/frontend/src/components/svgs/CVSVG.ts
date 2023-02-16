import {
    PathPart,
    SVGPartial,
    TemplateSVG
} from "@thesortex/vue-component-library/src/components/SVGTemplate/helper/SVG";

const paperline: PathPart[] = [
    {
        vector: {
            x: -120,
            y: -200
        }
    },
    {
        vector: {
            x: 120,
            y: -200
        }
    },
    {
        arc: {
            radius: 20,
            clockwise: false,
            rotation: 0,
            end: {
                x: 140,
                y: -180
            }
        }
    },
    {
        vector: {
            x: 140,
            y: 180
        }
    },
    {
        arc: {
            radius: 20,
            clockwise: false,
            rotation: 0,
            end: {
                x: 120,
                y: 200
            }
        }
    },
    {
        vector: {
            x: -120,
            y: 200
        }
    },
    {
        arc: {
            radius: 20,
            clockwise: false,
            rotation: 0,
            end: {
                x: -140,
                y: 180
            }
        }
    },
    {
        vector: {
            x: -140,
            y: -180
        }
    },
    {
        arc: {
            radius: 20,
            clockwise: false,
            rotation: 0,
            end: {
                x: -120,
                y: -200
            }
        }
    },
];

const paper: SVGPartial = {
    strokeColor: "rgba(var(--v-theme-on-background), 1)",
    strokeWidth: "10",
    fillColor: "white",
    scale: 1,
    parts: [
        paperline
    ]
};

export const CVSVG: TemplateSVG = {
    width: 500,
    height: 500,
    viewBox: "-250 -250 500 500",
    partials: [
        paper
    ]
};
