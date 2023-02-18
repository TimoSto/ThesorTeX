import {describe, expect, it} from "vitest";
import * as fs from "fs";
import path from "path";

describe("convertMarkdownToVuetify", () => {
    it("simple example", () => {
        const file = fs.readFileSync(path.resolve(__dirname, "./testfiles/simple_example.md"), "utf-8");

        expect(file).toEqual("g");
    });
});
