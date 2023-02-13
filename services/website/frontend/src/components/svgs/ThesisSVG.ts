import {PathPart} from "@thesortex/svg-generator/src/helper/generatePath";


export const Path: PathPart[] = [
    {
        vector: {
            x: -40,
            y: -20
        }
    },
    {
        vector: {
            x: 40,
            y: -20
        }
    },
    {
        arc: {
            radius: 10,
            rotation: 0,
            clockwise: false,
            xEnd: 50,
            yEnd: -10
        }
    },
    {
        vector: {
            x: 50,
            y: 10
        },
    },
    {
        arc: {
            radius: 10,
            rotation: 0,
            clockwise: false,
            xEnd: 40,
            yEnd: 20
        }
    },
    {
        vector: {
            x: -40,
            y: 20
        }
    },
    {
        arc: {
            radius: 10,
            rotation: 0,
            clockwise: false,
            xEnd: -50,
            yEnd: 10
        }
    },
    {
        vector: {
            x: -50,
            y: -10
        }
    },
    {
        arc: {
            radius: 10,
            rotation: 0,
            clockwise: false,
            xEnd: -40,
            yEnd: -20
        }
    },
];

export const Angle = 0;
