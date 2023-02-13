import {describe, expect, it} from "vitest";
import generatePath from "./generatePath";

describe("generatePath", () => {
    it("simple rect", () => {

        const points = [
            {
                x: 20,
                y: 20
            },
            {
                x: 100,
                y: 20
            },
            {
                x: 100,
                y: 50
            },
            {
                x: 20,
                y: 50
            }
        ];

        expect(generatePath(points, 0)).toEqual("M20,20 L100,20 L100,50 L20,50 z");
    });
    it("simple rect at 90 deg", () => {

        const points = [
            {
                x: 20,
                y: 20
            },
            {
                x: 100,
                y: 20
            },
            {
                x: 100,
                y: 50
            },
            {
                x: 20,
                y: 50
            }
        ];

        expect(generatePath(points, -90)).toEqual("M20,20 L20,100 L50,100 L50,20 z");
    });
});
