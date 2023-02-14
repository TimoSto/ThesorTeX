import {PathSegment} from "@thesortex/svg-generator/src/helper/generatePath";

const display: PathSegment = {
    strokeWidth: "10",
    strokeColor: "black",
    fillColor: "white",
    angle: 0,
    parts: [
        {
            vector: {
                x: -190,
                y: -200,
            },
        },
        {
            vector: {
                x: 190,
                y: -200,
            }
        },
        {
            arc: {
                radius: 10,
                rotation: 0,
                clockwise: false,
                xEnd: 200,
                yEnd: -190
            }
        },
        {
            vector: {
                x: 200,
                y: 50,
            }
        },
        {
            vector: {
                x: -200,
                y: 50,
            }
        },
        {
            vector: {
                x: -200,
                y: -190,
            }
        },
        {
            arc: {
                radius: 10,
                rotation: 0,
                clockwise: false,
                xEnd: -190,
                yEnd: -200
            }
        },
    ]
};

export const Paths: PathSegment[] = [
    display
];
