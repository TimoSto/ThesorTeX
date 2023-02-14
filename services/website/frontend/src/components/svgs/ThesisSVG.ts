import {PathSegment} from "@thesortex/svg-generator/src/helper/generatePath";

// book
const BookAngle = -10;

const roundedRect: PathSegment = {
    strokeColor: "rgba(var(--v-theme-on-background), 1)",
    strokeWidth: "10",
    fillColor: "white",
    parts: [
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
    ],
    angle: BookAngle
};

const bookMark: PathSegment = {
    strokeColor: "rgba(var(--v-theme-on-background), 1)",
    strokeWidth: "10",
    fillColor: "white",
    parts: [
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
    ],
    angle: BookAngle
};

const leftLine: PathSegment = {
    strokeColor: "rgba(var(--v-theme-on-background), 1)",
    strokeWidth: "10",
    fillColor: "white",
    parts: [
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
    ],
    angle: BookAngle
};

const circle: PathSegment = {
    strokeColor: "rgba(var(--v-theme-on-background), 1)",
    strokeWidth: "10",
    fillColor: "white",
    parts: [
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
    ],
    angle: BookAngle
};

const bottomLine1: PathSegment = {
    strokeColor: "rgba(var(--v-theme-on-background), 1)",
    strokeWidth: "10",
    fillColor: "white",
    parts: [
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
    ],
    angle: BookAngle
};

const bottomLine2: PathSegment = {
    strokeColor: "rgba(var(--v-theme-on-background), 1)",
    strokeWidth: "10",
    fillColor: "white",
    parts: [
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
    ],
    angle: BookAngle
};

const bottomLine3: PathSegment = {
    strokeColor: "rgba(var(--v-theme-on-background), 1)",
    strokeWidth: "10",
    fillColor: "white",
    parts: [
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
    ],
    angle: BookAngle
};

// hat
const hat: PathSegment = {
    strokeColor: "rgba(var(--v-theme-on-background), 1)",
    strokeWidth: "10",
    fillColor: "white",
    parts: [
        {
            vector: {
                x: 50,
                y: -30
            }
        },
        {
            vector: {
                x: 150,
                y: -80
            }
        },
        {
            vector: {
                x: 250,
                y: -30
            }
        },
        {
            vector: {
                x: 150,
                y: 20
            }
        },
        {
            vector: {
                x: 50,
                y: -30
            }
        }
    ],
    angle: 0
};

const hat2: PathSegment = {
    strokeColor: "rgba(var(--v-theme-on-background), 1)",
    strokeWidth: "10",
    fillColor: "white",
    parts: [
        {
            vector: {
                x: 100,
                y: -5
            }
        },
        {
            vector: {
                x: 100,
                y: 45
            }
        },
        {
            arc: {
                radius: 50,
                clockwise: true,
                rotation: 0,
                xEnd: 200,
                yEnd: 45
            }
        },
        {
            vector: {
                x: 200,
                y: -5
            }
        },
    ],
    angle: 0
};

export const Paths: PathSegment[] = [
    roundedRect,
    bookMark,
    leftLine,
    circle,
    bottomLine1,
    bottomLine2,
    bottomLine3,
    hat2,
    hat,
];
