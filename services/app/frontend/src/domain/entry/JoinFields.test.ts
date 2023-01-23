import {describe, expect, it} from "vitest";
import JoinFields from "./JoinFields";
import {Field} from "../category/Category";

describe("joinFields", () => {
    it("no overlap", () => {
        const b = [
            {
                Name: "t1"
            },
            {
                Name: "t2"
            },
            {
                Name: "t3"
            },
        ];

        const c = [
            {
                Name: "t4"
            },
            {
                Name: "t5"
            },
        ];

        const result = JoinFields(b as Field[], c as Field[]);
        expect(result.map(f => f.Name)).toEqual(["t1", "t2", "t3", "t4", "t5"]);
    });
    it("partial overlap", () => {
        const b = [
            {
                Name: "t1"
            },
            {
                Name: "t2"
            },
            {
                Name: "t3"
            },
        ];

        const c = [
            {
                Name: "t4"
            },
            {
                Name: "t3"
            },
        ];

        const result = JoinFields(b as Field[], c as Field[]);
        expect(result.map(f => f.Name)).toEqual(["t1", "t2", "t3", "t4"]);
    });

    it("full overlap", () => {
        const b = [
            {
                Name: "t1"
            },
            {
                Name: "t2"
            },
            {
                Name: "t3"
            },
        ];

        const c = [
            {
                Name: "t2"
            },
            {
                Name: "t3"
            },
        ];

        const result = JoinFields(b as Field[], c as Field[]);
        expect(result.map(f => f.Name)).toEqual(["t1", "t2", "t3"]);
    });
});
