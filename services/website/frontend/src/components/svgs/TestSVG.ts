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

const bookMark: PathPart[] = [
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


const leftMark: PathPart[] = [
    {
        vector: {
            x: -140,
            y: -200,
        }
    },
    {
        vector: {
            x: -140,
            y: 200,
        }
    }
];

const circle: PathPart[] = [
    {
        vector: {
            x: 0,
            y: -100,
        }
    },
    {
        arc: {
            radius: 40,
            clockwise: false,
            end: {
                x: 40,
                y: -60,
            },
            rotation: 0
        }
    },
    {
        arc: {
            radius: 40,
            clockwise: false,
            end: {
                x: 0,
                y: -20,
            },
            rotation: 0
        }
    },
    {
        arc: {
            radius: 40,
            clockwise: false,
            end: {
                x: -40,
                y: -60,
            },
            rotation: 0
        }
    },
    {
        arc: {
            radius: 40,
            clockwise: false,
            end: {
                x: 0,
                y: -100,
            },
            rotation: 0
        }
    }
];

const line1: PathPart[] = [
    {
        vector: {
            x: -100,
            y: 70
        }
    },
    {
        vector: {
            x: 100,
            y: 70
        }
    }
];

const line2: PathPart[] = [
    {
        vector: {
            x: -100,
            y: 100
        }
    },
    {
        vector: {
            x: 100,
            y: 100
        }
    }
];

const line3: PathPart[] = [
    {
        vector: {
            x: -100,
            y: 130
        }
    },
    {
        vector: {
            x: 100,
            y: 130
        }
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
        bookMark,
        leftMark,
        circle,
        line1,
        line2,
        line3
    ]
};

const hatPart1: PathPart[] = [
    {
        vector: {
            x: 50,
            y: -20
        }
    },
    {
        vector: {
            x: 150,
            y: -70
        }
    },
    {
        vector: {
            x: 250,
            y: -20
        }
    },
    {
        vector: {
            x: 150,
            y: 30
        }
    },
    {
        vector: {
            x: 50,
            y: -20
        }
    }
];

const hatPart2: PathPart[] = [
    {
        vector: {
            x: 90,
            y: 5
        }
    },
    {
        vector: {
            x: 90,
            y: 55
        }
    },
    {
        arc: {
            radius: 50,
            radiusY: 20,
            clockwise: true,
            rotation: 0,
            end: {
                x: 210,
                y: 55
            }
        }
    },
    {
        vector: {
            x: 210,
            y: 5
        }
    },
];

const hatPart3: PathPart[] = [
    {
        vector: {
            x: 230,
            y: -10
        }
    },
    {
        vector: {
            x: 230,
            y: 15
        }
    },
    {
        arc: {
            radius: 5,
            clockwise: true,
            rotation: 0,
            end: {
                x: 230,
                y: 25
            }
        }
    },
    {
        arc: {
            radius: 5,
            clockwise: true,
            rotation: 0,
            end: {
                x: 230,
                y: 15
            }
        }
    },
    {
        vector: {
            x: 230,
            y: 40
        }
    },
];

const hat: SVGPartial = {
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
        hatPart2,
        hatPart1,
        hatPart3
    ]
};

export const TestSVG: TemplateSVG = {
    viewBox: "-250 -250 500 500",
    width: 500,
    height: 500,
    partials: [
        p1,
        hat
    ]
};
