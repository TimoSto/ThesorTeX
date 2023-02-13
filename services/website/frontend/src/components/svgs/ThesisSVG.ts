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
            x: 30,
            y: -200
        },
    },
    {
        vector: {
            x: 30,
            y: -120
        },
    },
    {
        vector: {
            x: 60,
            y: -150
        },
    },
    {
        vector: {
            x: 90,
            y: -120
        },
    },
    {
        vector: {
            x: 90,
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

export const Paths: PathPart[][] = [
    roundedRect,
    bookMark,
    leftLine
];

export const Angle = -10;
