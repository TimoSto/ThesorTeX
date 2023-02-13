import {PathPart} from "@thesortex/svg-generator/src/helper/generatePath";

const roundedRect: PathPart[] = [
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
            xEnd: -180,
            yEnd: 160
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
            xEnd: -140,
            yEnd: -200
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

const leftLine: PathPart[] = [
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
            xEnd: 40,
            yEnd: -60,
            rotation: 0
        }
    },
    {
        arc: {
            radius: 40,
            clockwise: false,
            xEnd: 0,
            yEnd: -20,
            rotation: 0
        }
    },
    {
        arc: {
            radius: 40,
            clockwise: false,
            xEnd: -40,
            yEnd: -60,
            rotation: 0
        }
    },
    {
        arc: {
            radius: 40,
            clockwise: false,
            xEnd: 0,
            yEnd: -100,
            rotation: 0
        }
    }
];

export const Paths: PathPart[][] = [
    roundedRect,
    bookMark,
    leftLine,
    circle
];

export const Angle = 0;
