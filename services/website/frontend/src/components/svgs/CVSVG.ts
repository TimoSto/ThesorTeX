import {
    PathPart,
    SVGPartial,
    TemplateSVG
} from "@thesortex/vue-component-library/src/components/SVGTemplate/helper/SVG";

const paperline: PathPart[] = [
    {
        vector: {
            x: -110,
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
        arc: {
            radius: 20,
            clockwise: false,
            rotation: 0,
            end: {
                x: 160,
                y: -180
            }
        }
    },
    {
        vector: {
            x: 160,
            y: 180
        }
    },
    {
        arc: {
            radius: 20,
            clockwise: false,
            rotation: 0,
            end: {
                x: 140,
                y: 200
            }
        }
    },
    {
        vector: {
            x: -140,
            y: 200
        }
    },
    {
        arc: {
            radius: 20,
            clockwise: false,
            rotation: 0,
            end: {
                x: -160,
                y: 180
            }
        }
    },
    {
        vector: {
            x: -160,
            y: -150
        }
    },
    {
        vector: {
            x: -110,
            y: -200
        }
    },
];

const lash: PathPart[] = [
    {
        vector: {
            x: -110,
            y: -200
        }
    },
    {
        vector: {
            x: -110,
            y: -160
        }
    },
    {
        arc: {
            radius: 10,
            clockwise: false,
            rotation: 0,
            end: {
                x: -120,
                y: -150
            }
        }
    },
    {
        vector: {
            x: -160,
            y: -150
        }
    },
    {
        vector: {
            x: -120,
            y: -150
        }
    },
    {
        arc: {
            radius: 10,
            clockwise: true,
            rotation: 0,
            end: {
                x: -110,
                y: -160
            }
        }
    },
];

const line1: PathPart[] = [
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

const line2: PathPart[] = [
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

const line3: PathPart[] = [
    {
        vector: {
            x: -100,
            y: 100
        },
    },
    {
        vector: {
            x: 100,
            y: 100
        },
    }
];

const line4: PathPart[] = [
    {
        vector: {
            x: -100,
            y: -40
        },
    },
    {
        vector: {
            x: -20,
            y: -40
        },
    }
];

const line5: PathPart[] = [
    {
        vector: {
            x: -100,
            y: -60
        },
    },
    {
        vector: {
            x: -20,
            y: -60
        },
    }
];

const line6: PathPart[] = [
    {
        vector: {
            x: -100,
            y: -80
        },
    },
    {
        vector: {
            x: -20,
            y: -80
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
        line3,
        line4,
        line5,
        line6,
        lash
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

const border: PathPart[] = [
    {
        vector: {
            x: -50,
            y: -60
        },
    },
    {
        arc: {
            radius: 10,
            rotation: 0,
            clockwise: false,
            end: {
                x: -60,
                y: -70
            }
        }
    },
    {
        vector: {
            x: -60,
            y: -190
        },
    },
    {
        arc: {
            radius: 10,
            rotation: 0,
            clockwise: false,
            end: {
                x: -50,
                y: -200
            }
        }
    },
    {
        vector: {
            x: 50,
            y: -200
        },
    },
    {
        arc: {
            radius: 10,
            rotation: 0,
            clockwise: false,
            end: {
                x: 60,
                y: -190
            }
        }
    },
    {
        vector: {
            x: 60,
            y: -70
        },
    },
    {
        arc: {
            radius: 10,
            rotation: 0,
            clockwise: false,
            end: {
                x: 50,
                y: -60
            }
        }
    }
];

const person: SVGPartial = {
    strokeColor: "rgba(var(--v-theme-on-background), 1)",
    strokeWidth: "7",
    fillColor: "white",
    scale: 0.8,
    parts: [
        border,
        body,
        head,
    ],
    translate: {
        x: 60,
        y: 20
    }
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

export function GetCvSVG(fill: string, stroke: string) {
    const paperPart: SVGPartial = {
        strokeColor: stroke,
        strokeWidth: "10",
        fillColor: fill,
        scale: 1,
        parts: [
            paperline,
            line1,
            line2,
            line3,
            line4,
            line5,
            line6,
            lash
        ]
    };

    const personPart: SVGPartial = {
        strokeColor: stroke,
        strokeWidth: "7",
        fillColor: fill,
        scale: 0.8,
        parts: [
            border,
            body,
            head,
        ],
        translate: {
            x: 60,
            y: 20
        }
    };

    return {
        width: 500,
        height: 500,
        viewBox: "-250 -250 500 500",
        partials: [
            paperPart,
            personPart
        ]
    };
}
