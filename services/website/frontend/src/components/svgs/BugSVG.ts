import {
    MirrorVectorsY,
    PathPart,
    SVGPartial,
    TemplateSVG
} from "@thesortex/vue-component-library/src/components/SVGTemplate/helper/SVG";

const headPart1: PathPart[] = [
    {
        vector: {
            x: -50,
            y: -100
        }
    },
    {
        arc: {
            radius: 10,
            clockwise: false,
            rotation: 0,
            end: {
                x: 50,
                y: -100
            }
        }
    }
];

const eye1: PathPart[] = [
    {
        vector: {
            x: -20,
            y: -120
        }
    },
    {
        arc: {
            radius: 2,
            clockwise: false,
            rotation: 0,
            end: {
                x: -20,
                y: -124
            }
        }
    },
    {
        arc: {
            radius: 2,
            clockwise: false,
            rotation: 0,
            end: {
                x: -20,
                y: -120
            }
        }
    }
];

const eye2: PathPart[] = [
    {
        vector: {
            x: 20,
            y: -120
        }
    },
    {
        arc: {
            radius: 2,
            clockwise: false,
            rotation: 0,
            end: {
                x: 20,
                y: -124
            }
        }
    },
    {
        arc: {
            radius: 2,
            clockwise: false,
            rotation: 0,
            end: {
                x: 20,
                y: -120
            }
        }
    }
];

const headpieceLeft: PathPart[] = [
    {
        vector: {
            x: -20,
            y: -145
        }
    },
    {
        vector: {
            x: -30,
            y: -175
        }
    },
    {
        vector: {
            x: -50,
            y: -180
        }
    },
    {
        vector: {
            x: -30,
            y: -175
        }
    },
];

const headpieceRight = MirrorVectorsY(JSON.parse(JSON.stringify(headpieceLeft)));

const head: SVGPartial = {
    strokeColor: "rgba(var(--v-theme-on-background), 1)",
    strokeWidth: "10",
    fillColor: "white",
    scale: 1,
    angle: 0,
    translate: {
        x: 0,
        y: 0,
    },
    parts: [
        headPart1,
        eye1,
        eye2,
        headpieceLeft,
        headpieceRight
    ]
};

const bodyLeft: PathPart[] = [
    {
        vector: {
            x: 0,
            y: -120
        }
    },
    {
        arc: {
            radius: 90,
            clockwise: false,
            rotation: 0,
            radiusY: 120,
            end: {
                x: 0,
                y: 120
            }
        }
    },
];

const bodyRight: PathPart[] = [
    {
        vector: {
            x: 0,
            y: -120
        }
    },
    {
        arc: {
            radius: 90,
            clockwise: true,
            rotation: 0,
            radiusY: 120,
            end: {
                x: 0,
                y: 120
            }
        }
    },
];

const arm1Left: PathPart[] = [
    {
        vector: {
            x: -80,
            y: -60
        }
    },
    {
        vector: {
            x: -110,
            y: -75
        }
    },
    {
        vector: {
            x: -140,
            y: -55
        }
    },
    {
        vector: {
            x: -110,
            y: -75
        }
    },
];

const arm1Right = MirrorVectorsY(JSON.parse(JSON.stringify(arm1Left)));

const arm2Left: PathPart[] = [
    {
        vector: {
            x: -90,
            y: 0
        }
    },
    {
        vector: {
            x: -140,
            y: 0
        }
    },
    {
        vector: {
            x: -160,
            y: 30
        }
    },
    {
        vector: {
            x: -140,
            y: 0
        }
    },
];

const arm2Right = MirrorVectorsY(JSON.parse(JSON.stringify(arm2Left)));

const arm3Left: PathPart[] = [
    {
        vector: {
            x: -70,
            y: 70
        }
    },
    {
        vector: {
            x: -100,
            y: 70
        }
    },
    {
        vector: {
            x: -120,
            y: 120
        }
    },
    {
        vector: {
            x: -130,
            y: 120
        }
    },
    {
        vector: {
            x: -120,
            y: 120
        }
    },
    {
        vector: {
            x: -100,
            y: 70
        }
    },
];

const arm3Right = MirrorVectorsY(JSON.parse(JSON.stringify(arm3Left)));

const body: SVGPartial = {
    strokeColor: "rgba(var(--v-theme-on-background), 1)",
    strokeWidth: "10",
    fillColor: "white",
    scale: 1,
    angle: 0,
    translate: {
        x: 0,
        y: 0,
    },
    parts: [
        bodyLeft,
        bodyRight,
        arm1Left,
        arm1Right,
        arm2Left,
        arm2Right,
        arm3Left,
        arm3Right
    ]
};

export const BugSVG: TemplateSVG = {
    viewBox: "-250 -250 500 500",
    width: 500,
    height: 500,
    partials: [body, head]
};

export function GetBugSVG(fill: string, stroke: string): TemplateSVG {
    const bodyPart: SVGPartial = {
        strokeColor: stroke,
        strokeWidth: "10",
        fillColor: fill,
        scale: 1,
        angle: 0,
        translate: {
            x: 0,
            y: 0,
        },
        parts: [
            bodyLeft,
            bodyRight,
            arm1Left,
            arm1Right,
            arm2Left,
            arm2Right,
            arm3Left,
            arm3Right
        ]
    };

    const headPart: SVGPartial = {
        strokeColor: stroke,
        strokeWidth: "10",
        fillColor: fill,
        scale: 1,
        angle: 0,
        translate: {
            x: 0,
            y: 0,
        },
        parts: [
            headPart1,
            eye1,
            eye2,
            headpieceLeft,
            headpieceRight
        ]
    };

    return {
        viewBox: "-250 -250 500 500",
        width: 500,
        height: 500,
        partials: [bodyPart, headPart]
    };
}
