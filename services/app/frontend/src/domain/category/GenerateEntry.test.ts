import {describe, expect, it} from "vitest";
import {Category} from "./Category";
import GenerateEntryForCategory from "./GenerateEntry";

describe("GenerateEntry", () => {
    it("should work on a simple category", () => {
        const category: Category = {
            Name: "test",
            CitaviCategory: "",
            BibFields: [
                {
                    Name: "f1",
                    Format: {
                        Prefix: "",
                        Suffix: " ",
                        Italic: false,
                        Preformatted: false
                    },
                    CitaviMapping: []
                },
                {
                    Name: "f2",
                    Format: {
                        Prefix: "",
                        Suffix: ", ",
                        Italic: false,
                        Preformatted: false
                    },
                    CitaviMapping: []
                },
                {
                    Name: "f3",
                    Format: {
                        Prefix: "",
                        Suffix: "",
                        Italic: false,
                        Preformatted: false
                    },
                    CitaviMapping: []
                }
            ],
            CiteFields: [],
            CitaviFilter: []
        };

        const got = GenerateEntryForCategory(category.BibFields, ["field1", "field2", "field3"]);
        const expected = "field1 field2, field3";

        expect(got).toEqual(expected);
    });

    it("should work on a category with italic fields", () => {
        const category: Category = {
            Name: "test",
            CitaviCategory: "",
            BibFields: [
                {
                    Name: "f1",
                    Format: {
                        Prefix: "",
                        Suffix: " ",
                        Italic: true,
                        Preformatted: false
                    },
                    CitaviMapping: []
                },
                {
                    Name: "f2",
                    Format: {
                        Prefix: "",
                        Suffix: ", ",
                        Italic: false,
                        Preformatted: false
                    },
                    CitaviMapping: []
                },
                {
                    Name: "f3",
                    Format: {
                        Prefix: "",
                        Suffix: "",
                        Italic: true,
                        Preformatted: false
                    },
                    CitaviMapping: []
                }
            ],
            CiteFields: [],
            CitaviFilter: []
        };

        const got = GenerateEntryForCategory(category.BibFields, ["field1", "field2", "field3"]);
        const expected = "<i>field1</i> field2, <i>field3</i>";

        expect(got).toEqual(expected);
    });

    it("should give empty on unknown category", () => {
        expect(GenerateEntryForCategory([], ["t", "ta", "to"])).toEqual("");
    });

    it("should print field empty, if it is not set", () => {
        const category: Category = {
            Name: "test",
            CitaviCategory: "",
            BibFields: [
                {
                    Name: "f1",
                    Format: {
                        Prefix: "",
                        Suffix: " ",
                        Italic: true,
                        Preformatted: false
                    },
                    CitaviMapping: []
                },
                {
                    Name: "f2",
                    Format: {
                        Prefix: "",
                        Suffix: ", ",
                        Italic: false,
                        Preformatted: false
                    },
                    CitaviMapping: []
                },
                {
                    Name: "f3",
                    Format: {
                        Prefix: "",
                        Suffix: "",
                        Italic: true,
                        Preformatted: false
                    },
                    CitaviMapping: []
                }
            ],
            CiteFields: [],
            CitaviFilter: []
        };

        const got = GenerateEntryForCategory(category.BibFields, ["field1", "field2"]);
        const expected = "<i>field1</i> field2, <i></i>";

        expect(got).toEqual(expected);
    });
});
