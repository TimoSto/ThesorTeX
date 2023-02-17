import {describe, expect, it} from "vitest";
import {PathPart, SVGPartial} from "./SVG";
import {generatePath, generateTransform} from "./generatePath";

describe("generatePath", () => {
    it("give rect", () => {
        const shape: PathPart[] = [
            {
                vector: {
                    x: 20,
                    y: 20,
                }
            },
            {
                vector: {
                    x: 50,
                    y: 20,
                }
            },
            {
                vector: {
                    x: 50,
                    y: 50,
                }
            },
            {
                vector: {
                    x: 20,
                    y: 50,
                }
            },
        ];

        const path = generatePath(shape);

        expect(path).toEqual("M 20 20 L 50 20 L 50 50 L 20 50 z");
    });
});

describe("generateTransform", () => {
    it("all set", () => {
        const partial: SVGPartial = {
            angle: 0,
            translate: {
                x: 20,
                y: 30
            },
            scale: 0.7
        } as SVGPartial;

        const transform = generateTransform(partial);
        expect(transform).toEqual("scale(0.7) rotate(0, 0, 0) translate(20 30)");
    });
});
