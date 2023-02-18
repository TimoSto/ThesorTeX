import {describe, expect, it} from "vitest";
import {analyseLine, AnalyseLineResult} from "./analyseLine";

describe("analyseLine", () => {
    it("h1", () => {
        const expected: AnalyseLineResult = {
            type: "H1",
            content: "TestHeading"
        };
        expect(analyseLine("# TestHeading")).toEqual(expected);
    });
    it("h2", () => {
        const expected: AnalyseLineResult = {
            type: "H2",
            content: "TestHeading"
        };
        expect(analyseLine("## TestHeading")).toEqual(expected);
    });
    it("h3", () => {
        const expected: AnalyseLineResult = {
            type: "H3",
            content: "TestHeading"
        };
        expect(analyseLine("### TestHeading")).toEqual(expected);
    });
    it("h4", () => {
        const expected: AnalyseLineResult = {
            type: "H4",
            content: "TestHeading"
        };
        expect(analyseLine("#### TestHeading")).toEqual(expected);
    });
    it("h5", () => {
        const expected: AnalyseLineResult = {
            type: "H5",
            content: "TestHeading"
        };
        expect(analyseLine("##### TestHeading")).toEqual(expected);
    });
});
