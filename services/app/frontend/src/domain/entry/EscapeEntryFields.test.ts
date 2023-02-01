import {describe, expect, it} from "vitest";
import {Field} from "../category/Category";
import {Entry} from "./Entry";
import {EscapeEntryFields, UnEscapeEntryFields} from "./EscapeEntryFields";

const testFields: Field[] = [
    {
        Name: "f1",
        Format: {
            Preformatted: false
        }
    } as Field,
    {
        Name: "f2",
        Format: {
            Preformatted: false
        }
    } as Field,
    {
        Name: "f3",
        Format: {
            Preformatted: false
        }
    } as Field,
    {
        Name: "f4",
        Format: {
            Preformatted: true
        }
    } as Field,
    {
        Name: "f5",
        Format: {
            Preformatted: false
        }
    } as Field,
    {
        Name: "f6",
        Format: {
            Preformatted: false
        }
    } as Field
];

describe("escape and unescape", () => {
    describe("escape", () => {
        it("entry with no parsing necessary", () => {
            const e: Entry = {
                Key: "",
                Category: "",
                Fields: ["v1", "v2", "v3", "v4", "v5", "v6"]
            };

            expect(EscapeEntryFields(e, testFields)).toEqual(["v1", "v2", "v3", "v4", "v5", "v6"]);
        });
        it("entry with one field more", () => {
            const e: Entry = {
                Key: "",
                Category: "",
                Fields: ["v1", "v2", "v3", "v4", "v5", "v6", "v7"]
            };

            expect(EscapeEntryFields(e, testFields)).toEqual(["v1", "v2", "v3", "v4", "v5", "v6", "v7"]);
        });
        it("entry with one field needing escaping at end", () => {
            const e: Entry = {
                Key: "",
                Category: "",
                Fields: ["v1_", "v2", "v3", "v4", "v5", "v6"]
            };

            expect(EscapeEntryFields(e, testFields)).toEqual(["v1{{\\_}}", "v2", "v3", "v4", "v5", "v6"]);
        });
        it("entry with already escaped", () => {
            const e: Entry = {
                Key: "",
                Category: "",
                Fields: ["v1{{\\_}}", "v2", "v3", "v4", "v5", "v6"]
            };

            expect(EscapeEntryFields(e, testFields)).toEqual(["v1{{\\_}}", "v2", "v3", "v4", "v5", "v6"]);
        });
        it("entry with multiple fields needing escaping at different places", () => {
            const e: Entry = {
                Key: "",
                Category: "",
                Fields: ["v1_", "v2_q", "v3", "v4", ";v_5$á", "v6"]
            };

            expect(EscapeEntryFields(e, testFields)).toEqual(["v1{{\\_}}", "v2{{\\_}}q", "v3", "v4", "{{;}}v{{\\_}}5{{\\$}}{{\\'{a}}}", "v6"]);
        });
        it("dont escape in preformatted", () => {
            it("entry with multiple fields needing escaping at different places", () => {
                const e: Entry = {
                    Key: "",
                    Category: "",
                    Fields: ["v1_", "v2_q", "v3", "v_;4", ";v_5$á", "v6"]
                };

                expect(EscapeEntryFields(e, testFields)).toEqual(["v1{{\\_}}", "v2{{\\_}}q", "v3", "v_;4", "{{;}}v{{\\_}}5{{\\$}}{{\\'{a}}}", "v6"]);
            });
        });
    });
    describe("unescape", () => {
        it("nothing to unescape", () => {
            expect(UnEscapeEntryFields(["v1", "v2", "v3", "v4", "v5", "v6"], testFields)).toEqual(["v1", "v2", "v3", "v4", "v5", "v6"]);
        });
        it("unescape", () => {
            expect(UnEscapeEntryFields(["v1{{\\_}}", "{{\\_}}v2", "{{\\_}}v{{\\_}}3{{\\_}}", "v4", "v{{\\_}}5", "v6"], testFields)).toEqual(["v1_", "_v2", "_v_3_", "v4", "v_5", "v6"]);
        });
        it("dont unescape preformatted", () => {
            expect(UnEscapeEntryFields(["v1{{\\_}}", "{{\\_}}v2", "{{\\_}}v{{\\_}}3{{\\_}}", "v{{\\_}}4", "v{{\\_}}5", "v6"], testFields)).toEqual(["v1_", "_v2", "_v_3_", "v{{\\_}}4", "v_5", "v6"]);
        });
    });
});
