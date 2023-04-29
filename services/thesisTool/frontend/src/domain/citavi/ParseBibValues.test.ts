import {describe, expect, it} from "vitest";
import {trimAndParseValue} from "./ParseBibValues";

describe("trimAndParseBibValues", () => {
    it("should rm surrounding brackets and quotations", function () {
        expect(trimAndParseValue("{\"hallo\"}")).toEqual("hallo");
    });
    it("should rm surrounding brackets", function () {
        expect(trimAndParseValue("{hallo}")).toEqual("hallo");
    });
    it("should rm surrounding quotations", function () {
        expect(trimAndParseValue("\"hallo\"")).toEqual("hallo");
    });
    it("should replace escaped umlauts", function () {
        expect(trimAndParseValue("hallo {\\\"a}{\\\"u}{\\\"o}")).toEqual("hallo äüö");
    });
    it("should replace ß and _ umlauts", function () {
        expect(trimAndParseValue("hallo {{\\textunderscore}}{{\\ss}}")).toEqual("hallo {{\\_}}ß");
    });
});
