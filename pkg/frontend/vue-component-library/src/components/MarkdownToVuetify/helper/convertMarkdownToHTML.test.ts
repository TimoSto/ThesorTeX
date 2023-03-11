import {describe, expect, it} from "vitest";
import * as fs from "fs";
import path from "path";
import {convertMarkdownToHTML} from "./convertMarkdownToHTML";

describe("convertMarkdownToVuetify", () => {
    it("simple example", () => {
        const file = fs.readFileSync(path.resolve(__dirname, "./testfiles/simple_example.md"), "utf-8");
        const expected = fs.readFileSync(path.resolve(__dirname, "./testfiles/simple_example.html"), "utf-8");

        const converted = convertMarkdownToHTML(file);

        expect(converted).toEqual(expected);
    });
});
