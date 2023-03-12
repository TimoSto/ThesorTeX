import {describe, expect, it} from "vitest";
import fs from "fs";
import path from "path";
import {ConvertMarkdownToVuetify, VuetifyComponent} from "./convertMarkdownToVuetify";

describe("convertMarkdownToVuetify()", () => {
    it("should give the expected array", function () {
        const file = fs.readFileSync(path.resolve(__dirname, "./testfiles/simple_example.md"), "utf-8");

        const expected: VuetifyComponent[] = [
            {
                Tag: "H1",
                Content: "Test Heading"
            },
            {
                Tag: "TEXT",
                Content: "Some content Some more content"
            },
            {
                Tag: "TEXT",
                Content: "Even more content in <b>bold</b> and <i>italic</i> and <b>second bold</b>"
            },
            {
                Tag: "HR",
                Content: ""
            },
            {
                Tag: "H2",
                Content: "Heading 2"
            },
            {
                Tag: "TEXT",
                Content: "hallo"
            },
            {
                Tag: "LIST",
                Content: ["test", "test2"]
            },
            {
                Tag: "TEXT",
                Content: "ende"
            },
            {
                Tag: "CODE",
                Content: ["\\part{Test}\n"]
            },
        ];

        const converted = ConvertMarkdownToVuetify(file);

        //expect(converted.length).toEqual(expected.length);

        converted.forEach((l, i) => {
            expect(l.Tag).toEqual(expected[i].Tag);
            expect(l.Content).toEqual(expected[i].Content);
        });
    });
});
