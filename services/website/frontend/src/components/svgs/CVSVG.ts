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

const line1: PathPart[] = [
    {
        vector: {
            x: -100,
            y: -50
        },
    },
    {
        vector: {
            x: 100,
            y: -50
        },
    }
];

const line2: PathPart[] = [
    {
        vector: {
            x: -100,
            y: 0
        },
    },
    {
        vector: {
            x: 100,
            y: 0
        },
    }
];

const line3: PathPart[] = [
    {
        vector: {
            x: -100,
            y: 50
        },
    },
    {
        vector: {
            x: 100,
            y: 50
        },
    }
];

const paper: SVGPartial = {
    strokeColor: "rgba(var(--v-theme-on-background), 1)",
    strokeWidth: "10",
    fillColor: "white",
    scale: 1,
    parts: [
        paperline,
        line1,
        line2,
        line3
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
