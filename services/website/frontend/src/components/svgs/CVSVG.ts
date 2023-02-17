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
            y: -30
        },
    },
    {
        vector: {
            x: 100,
            y: -30
        },
    }
];

const line2: PathPart[] = [
    {
        vector: {
            x: -100,
            y: 20
        },
    },
    {
        vector: {
            x: 100,
            y: 20
        },
    }
];

const line3: PathPart[] = [
    {
        vector: {
            x: -100,
            y: 70
        },
    },
    {
        vector: {
            x: 100,
            y: 70
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

const body: PathPart[] = [
    {
        vector: {
            x: -50,
            y: -70
        },
    },
    {
        vector: {
            x: 50,
            y: -70
        },
    },
    {
        arc: {
            radius: 30,
            radiusY: 50,
            rotation: 0,
            clockwise: true,
            end: {
                x: 20,
                y: -120
            }
        }
    },
    {
        arc: {
            radius: 20,
            radiusY: 10,
            rotation: 0,
            clockwise: false,
            end: {
                x: -20,
                y: -120
            }
        }
    },
    {
        arc: {
            radius: 30,
            radiusY: 50,
            rotation: 0,
            clockwise: true,
            end: {
                x: -50,
                y: -70
            }
        }
    },
];

const head: PathPart[] = [
    {
        vector: {
            x: 0,
            y: -130
        }
    },
    {
        arc: {
            radius: 25,
            rotation: 0,
            clockwise: true,
            end: {
                x: 0,
                y: -180
            }
        }
    },
    {
        arc: {
            radius: 25,
            rotation: 0,
            clockwise: true,
            end: {
                x: 0,
                y: -130
            }
        }
    }
];

const person: SVGPartial = {
    strokeColor: "rgba(var(--v-theme-on-background), 1)",
    strokeWidth: "7",
    fillColor: "white",
    scale: 1,
    parts: [
        body,
        head
    ]
};

//TODO: diploma icon or smth

export const CVSVG: TemplateSVG = {
    width: 500,
    height: 500,
    viewBox: "-250 -250 500 500",
    partials: [
        paper,
        person
    ]
};
