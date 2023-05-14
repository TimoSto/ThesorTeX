import {
    PathPart,
    SVGPartial,
    TemplateSVG
} from "@thesortex/vue-component-library/src/components/SVGTemplate/helper/SVG";
import {GetThesisSVG, ThesisSVG} from "./ThesisSVG";

const display: PathPart[] = [
    {
        vector: {
            x: -210,
            y: -150
        }
    },
    {
        vector: {
            x: 210,
            y: -150
        }
    },
    {
        arc: {
            radius: 20,
            rotation: 0,
            clockwise: false,
            end: {
                x: 230,
                y: -130
            }
        }
    },
    {
        vector: {
            x: 230,
            y: 100
        }
    },
    {
        vector: {
            x: -230,
            y: 100
        }
    },
    {
        vector: {
            x: -230,
            y: -130
        }
    },
    {
        arc: {
            radius: 20,
            rotation: 0,
            clockwise: false,
            end: {
                x: -210,
                y: -150
            }
        }
    },
];

const keyboard: PathPart[] = [
    {
        vector: {
            x: -240,
            y: 120
        }
    },
    {
        vector: {
            x: 240,
            y: 120
        }
    },
    {
        arc: {
            radius: 10,
            rotation: 0,
            clockwise: false,
            end: {
                x: 250,
                y: 130
            }
        }
    },
    {
        vector: {
            x: -250,
            y: 130
        }
    },
    {
        arc: {
            radius: 10,
            rotation: 0,
            clockwise: false,
            end: {
                x: -240,
                y: 120
            }
        }
    },
];

const laptop: SVGPartial = {
    strokeColor: "black",
    strokeWidth: "15",
    fillColor: "white",
    parts: [
        display,
        keyboard
    ]
};

const laptopMark: SVGPartial = {
    strokeColor: "white",
    strokeWidth: "5",
    fillColor: "white",
    parts: [
        [
            {
                vector: {
                    x: -50,
                    y: 122
                }
            },
            {
                vector: {
                    x: 50,
                    y: 122
                }
            },
            {
                arc: {
                    radius: 3,
                    rotation: 0,
                    clockwise: false,
                    end: {
                        x: 50,
                        y: 128
                    }
                }
            },
            {
                vector: {
                    x: -50,
                    y: 128
                }
            },
            {
                arc: {
                    radius: 3,
                    rotation: 0,
                    clockwise: false,
                    end: {
                        x: -50,
                        y: 122
                    }
                }
            },
        ]
    ]
};

const thesisPart1 = JSON.parse(JSON.stringify(ThesisSVG.partials[0]));
thesisPart1.scale = 0.45;
thesisPart1.translate = {
    x: 20,
    y: -50
};

const thesisPart2 = JSON.parse(JSON.stringify(ThesisSVG.partials[1]));
thesisPart2.scale = 0.45;
thesisPart2.translate = {
    x: 0,
    y: -50
};

export const LaptopSVG: TemplateSVG = {
    viewBox: "-250 -250 500 500",
    width: 500,
    height: 500,
    partials: [
        laptop,
        laptopMark,
        thesisPart1,
        thesisPart2
    ]
};

export function GetLaptopSVG(fill: string, stroke: string): TemplateSVG {
    const laptopPart: SVGPartial = {
        strokeColor: stroke,
        strokeWidth: "15",
        fillColor: fill,
        parts: [
            display,
            keyboard
        ]
    };

    const laptopMarkPart: SVGPartial = {
        strokeColor: fill,
        strokeWidth: "5",
        fillColor: fill,
        parts: [
            [
                {
                    vector: {
                        x: -50,
                        y: 122
                    }
                },
                {
                    vector: {
                        x: 50,
                        y: 122
                    }
                },
                {
                    arc: {
                        radius: 3,
                        rotation: 0,
                        clockwise: false,
                        end: {
                            x: 50,
                            y: 128
                        }
                    }
                },
                {
                    vector: {
                        x: -50,
                        y: 128
                    }
                },
                {
                    arc: {
                        radius: 3,
                        rotation: 0,
                        clockwise: false,
                        end: {
                            x: -50,
                            y: 122
                        }
                    }
                },
            ]
        ]
    };

    const thesisSVG = GetThesisSVG("white", "black");

    const thesisPart_1 = JSON.parse(JSON.stringify(thesisSVG.partials[0]));
    thesisPart_1.scale = 0.45;
    thesisPart_1.translate = {
        x: 20,
        y: -50
    };

    const thesisPart_2 = JSON.parse(JSON.stringify(thesisSVG.partials[1]));
    thesisPart_2.scale = 0.45;
    thesisPart_2.translate = {
        x: 0,
        y: -50
    };

    return {
        viewBox: "-250 -250 500 500",
        width: 500,
        height: 500,
        partials: [
            laptopPart,
            laptopMarkPart,
            thesisPart_1,
            thesisPart_2
        ]
    };
}
