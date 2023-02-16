import {
    PathPart,
    SVGPartial,
    TemplateSVG
} from "@thesortex/vue-component-library/src/components/SVGTemplate/helper/SVG";
import {ThesisSVG} from "./ThesisSVG";

const display: PathPart[] = [
    {
        vector: {
            x: -210,
            y: -200
        }
    },
    {
        vector: {
            x: 210,
            y: -200
        }
    },
    {
        arc: {
            radius: 20,
            rotation: 0,
            clockwise: false,
            end: {
                x: 230,
                y: -180
            }
        }
    },
    {
        vector: {
            x: 230,
            y: 50
        }
    },
    {
        vector: {
            x: -230,
            y: 50
        }
    },
    {
        vector: {
            x: -230,
            y: -180
        }
    },
    {
        arc: {
            radius: 20,
            rotation: 0,
            clockwise: false,
            end: {
                x: -210,
                y: -200
            }
        }
    },
];

const keyboard: PathPart[] = [
    {
        vector: {
            x: -240,
            y: 70
        }
    },
    {
        vector: {
            x: 240,
            y: 70
        }
    },
    {
        arc: {
            radius: 10,
            rotation: 0,
            clockwise: false,
            end: {
                x: 250,
                y: 80
            }
        }
    },
    {
        vector: {
            x: -250,
            y: 80
        }
    },
    {
        arc: {
            radius: 10,
            rotation: 0,
            clockwise: false,
            end: {
                x: -240,
                y: 70
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

const thesisPart1 = JSON.parse(JSON.stringify(ThesisSVG.partials[0]));
thesisPart1.scale = 0.45;
thesisPart1.translate = {
    x: 40,
    y: -160
};

const thesisPart2 = JSON.parse(JSON.stringify(ThesisSVG.partials[1]));
thesisPart2.scale = 0.45;
thesisPart2.translate = {
    x: 0,
    y: -160
};

export const LaptopSVG: TemplateSVG = {
    viewBox: "-250 -250 500 500",
    width: 500,
    height: 500,
    partials: [
        laptop,
        thesisPart1,
        thesisPart2
    ]
};
