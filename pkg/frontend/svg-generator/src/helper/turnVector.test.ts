import {describe, expect, it} from "vitest";
import turnVector from "./turnVector";

describe("turnVector", () => {
    it("should work on 90°", () => {
        const result = turnVector({x: 5, y: 0}, 90);
        const expected = {
            x: 0,
            y: 5
        };
        expect(Math.abs(result.x - expected.x)).toBeLessThan(0.05);
        expect(Math.abs(result.y - expected.y)).toBeLessThan(0.05);
    });
    it("should work on -90°", () => {
        const result = turnVector({x: 5, y: 0}, -90);
        const expected = {
            x: 0,
            y: -5
        };
        expect(Math.abs(result.x - expected.x)).toBeLessThan(0.05);
        expect(Math.abs(result.y - expected.y)).toBeLessThan(0.05);
    });
    it("should work on already tilted", () => {
        const result = turnVector({x: 5, y: 5}, -45);
        const expected = {
            x: 7.07,
            y: 0
        };
        console.log(result);
        expect(Math.abs(result.x - expected.x)).toBeLessThan(0.05);
        expect(Math.abs(result.y - expected.y)).toBeLessThan(0.05);
    });
});
