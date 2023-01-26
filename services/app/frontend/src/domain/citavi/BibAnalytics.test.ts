import {describe, expect, it} from "vitest";
import {ReadFile} from "../testdata/testdataReader";
import {AssignCategory, getCategoryScore, GetEntries} from "./BibAnalytics";
import {AttributeValue, CitaviEntry} from "./Citavi";
import {Category} from "../category/Category";
import {Entry} from "../entry/Entry";

describe("BibAnalytics", () => {
    describe("GetEntries", () => {
        it("should give one entry", () => {
            const file = ReadFile("files/citavi/oneEntry.bib");

            const citaviEntries = GetEntries(file);

            expect(citaviEntries[0].Key).toEqual("f1");
            expect(citaviEntries[0].Category).toEqual("book");
            expect(citaviEntries[0].Attributes[0].Attr).toEqual("field1");
            expect(citaviEntries[0].Attributes[0].Value).not.toEqual("");
            expect(citaviEntries[0].Attributes[1].Attr).toEqual("field2");
            expect(citaviEntries[0].Attributes[1].Value).not.toEqual("");
        });
        it("should give two entries", () => {
            const file = ReadFile("files/citavi/twoEntriesNewline.bib");

            const citaviEntries = GetEntries(file);

            expect(citaviEntries[0].Key).toEqual("f1");
            expect(citaviEntries[0].Category).toEqual("book");
            expect(citaviEntries[0].Attributes[0].Attr).toEqual("field1");
            expect(citaviEntries[0].Attributes[0].Value).not.toEqual("");
            expect(citaviEntries[0].Attributes[1].Attr).toEqual("field2");
            expect(citaviEntries[0].Attributes[1].Value).not.toEqual("");

            expect(citaviEntries[1].Key).toEqual("f2");
            expect(citaviEntries[1].Category).toEqual("wer");
            expect(citaviEntries[1].Attributes[0].Attr).toEqual("field1");
            expect(citaviEntries[1].Attributes[0].Value).not.toEqual("");
            expect(citaviEntries[1].Attributes[1].Attr).toEqual("field2");
            expect(citaviEntries[1].Attributes[1].Value).not.toEqual("");
        });
        it("should give two entries", () => {
            const file = ReadFile("files/citavi/twoEntriesBracketsSameLine.bib");

            const citaviEntries = GetEntries(file);

            expect(citaviEntries[0].Key).toEqual("f1");
            expect(citaviEntries[0].Category).toEqual("book");
            expect(citaviEntries[0].Attributes[0].Attr).toEqual("field1");
            expect(citaviEntries[0].Attributes[0].Value).not.toEqual("");
            expect(citaviEntries[0].Attributes[1].Attr).toEqual("field2");
            expect(citaviEntries[0].Attributes[1].Value).not.toEqual("");

            expect(citaviEntries[1].Key).toEqual("f2");
            expect(citaviEntries[1].Category).toEqual("wer");
            expect(citaviEntries[1].Attributes[0].Attr).toEqual("field1");
            expect(citaviEntries[1].Attributes[0].Value).not.toEqual("");
            expect(citaviEntries[1].Attributes[1].Attr).toEqual("field2");
            expect(citaviEntries[1].Attributes[1].Value).not.toEqual("");
        });
        // it("should get all attributes", () => {
        //     const file = ReadFile("files/citavi/moreAttributes.bib");
        //
        //     const citaviEntries = AnalyseBibFile(file);
        //
        //     const expected = [
        //         {
        //             Key: "Thiemann2008",
        //             Category: "book",
        //             Attributes: [
        //                 {Attr: "title", Value: "Barrierefreiheit"},
        //                 {
        //                     Attr: "bookTitle",
        //                     Value: "Benutzerfreundliche Online-Hilfen: Grundlagen und Umsetzung mit MadCap Flare"
        //                 },
        //                 {Attr: "year", Value: "2008"},
        //                 {Attr: "publisher", Value: "Vieweg+Teubner"},
        //                 {Attr: "address", Value: "Wiesbaden"},
        //                 {Attr: "pages", Value: "143--157"},
        //                 {
        //                     Attr: "abstract",
        //                     Value: "Barrierefreies Design bedeutet, dass Menschen mit Behinderungen ein elektronisches Angebot uneingeschränkt und selbstständig nutzen können. Im Sinne der Nutzbarkeit muss die Barrierefreiheit beispielsweise einer Online-Hilfe im Web einen Schritt über die reine Zugänglichkeit hinausgehen: auch behinderte Nutzer sollen mit ihren Fähigkeiten und Hilfsmitteln elektronische Angebote nutzen können. Barrierefreiheit ist somit Gebrauchstauglichkeit vor dem Hintergrund einer Behinderung oder Einschränkung."
        //                 },
        //                 {Attr: "isbn", Value: "978-3-8348-9483-0"},
        //                 {Attr: "doi", Value: "10.1007/978-3-8348-9483-0_7"},
        //                 {Attr: "url", Value: "https://doi.org/10.1007/978-3-8348-9483-0_7"}
        //             ]
        //         }
        //     ] as CitaviEntry[];
        //
        //     console.log(citaviEntries[0].Attributes[6]);
        //
        //     expect(citaviEntries).toEqual(expected);
        // });
    });
    describe("getCategoryScore", () => {
        it("should give -1 if category name does not match", () => {
            expect(getCategoryScore({
                Key: "f",
                Category: "test",
                Attributes: [] as AttributeValue[]
            }, {
                Name: "teste",
                CitaviCategory: "teste",
                CitaviFilter: [] as string[],
            } as Category)).toEqual(-1);
        });
        it("should give 0 if category name does match", () => {
            expect(getCategoryScore({
                Key: "f",
                Category: "test",
                Attributes: [] as AttributeValue[]
            }, {
                Name: "teste",
                CitaviCategory: "test",
                CitaviFilter: [] as string[],
            } as Category)).toEqual(0);
        });
        it("should give 1 if one attribute matches", () => {
            expect(getCategoryScore({
                Key: "f",
                Category: "test",
                Attributes: [
                    {Attr: "t1", Value: "t2"}
                ]
            }, {
                Name: "teste",
                CitaviCategory: "test",
                CitaviFilter: ["t1"],
            } as Category)).toEqual(1);
        });
        it("should give 2 if 2 attributes matches", () => {
            expect(getCategoryScore({
                Key: "f",
                Category: "test",
                Attributes: [
                    {Attr: "t1", Value: "t2"},
                    {Attr: "ta", Value: "t2"}
                ]
            }, {
                Name: "teste",
                CitaviCategory: "test",
                CitaviFilter: ["t1", "ta"],
            } as Category)).toEqual(2);
        });
        it("should give -1 if 2 attributes matche but not category name", () => {
            expect(getCategoryScore({
                Key: "f",
                Category: "tester",
                Attributes: [
                    {Attr: "t1", Value: "t2"},
                    {Attr: "ta", Value: "t2"}
                ]
            }, {
                Name: "teste",
                CitaviCategory: "test",
                CitaviFilter: ["t1", "ta"],
            } as Category)).toEqual(-1);
        });
        it("should give -1 if an attribute is missing", () => {
            expect(getCategoryScore({
                Key: "f",
                Category: "tester",
                Attributes: [
                    {Attr: "t1", Value: "t2"}
                ]
            }, {
                Name: "teste",
                CitaviCategory: "test",
                CitaviFilter: ["t1", "ta"],
            } as Category)).toEqual(-1);
        });
    });
    describe("AssignCategory", () => {
        it("should give c1 if only c1 exists but matches", () => {
            const cs: Category[] = [
                {
                    Name: "c1",
                    CitaviCategory: "t1",
                    CitaviFilter: [] as string[],
                    BibFields: [
                        {
                            Name: "test1",
                            CitaviMapping: ["attr1"],
                            Format: {
                                Preformatted: false
                            }
                        }
                    ]
                } as Category
            ];

            const e: CitaviEntry = {
                Key: "test",
                Category: "t1",
                Attributes: [
                    {
                        Attr: "attr1",
                        Value: "test"
                    }
                ]
            };

            const expectedEntry: Entry = {
                Key: "test",
                Category: "c1",
                Fields: [
                    "test"
                ]
            };

            expect(AssignCategory(e, cs)).toEqual(expectedEntry);
        });
        it("should give undefined if only c1 exists but matches", () => {
            const cs: Category[] = [
                {
                    Name: "c1",
                    CitaviCategory: "t1",
                    CitaviFilter: [] as string[],
                } as Category
            ];

            const e: CitaviEntry = {
                Key: "test",
                Category: "t2",
                Attributes: []
            };

            expect(AssignCategory(e, cs)).toEqual(undefined);
        });
        it("should give c1 if it has a higher score that c2", () => {
            const cs: Category[] = [
                {
                    Name: "c1",
                    CitaviCategory: "t1",
                    CitaviFilter: ["attr1"],
                    BibFields: [
                        {
                            Name: "test1",
                            CitaviMapping: ["attr1"],
                            Format: {
                                Preformatted: false
                            }
                        }
                    ]
                } as Category,
                {
                    Name: "c2",
                    CitaviCategory: "t1",
                    CitaviFilter: [] as string[],
                    BibFields: [
                        {
                            Name: "test1",
                            CitaviMapping: ["attr1"],
                            Format: {
                                Preformatted: false
                            }
                        }
                    ]
                } as Category
            ];

            const e: CitaviEntry = {
                Key: "test",
                Category: "t1",
                Attributes: [
                    {Attr: "attr1", Value: "test"}
                ]
            };

            const expectedEntry: Entry = {
                Key: "test",
                Category: "c1",
                Fields: [
                    "test"
                ]
            };

            expect(AssignCategory(e, cs)).toEqual(expectedEntry);
        });
        it("should give c2 if an attribute from c1 is missing", () => {
            const cs: Category[] = [
                {
                    Name: "c1",
                    CitaviCategory: "t1",
                    CitaviFilter: ["attr1", "attr2"],
                    BibFields: [
                        {
                            Name: "test1",
                            CitaviMapping: ["attr1"],
                            Format: {
                                Preformatted: false
                            }
                        }
                    ]
                } as Category,
                {
                    Name: "c2",
                    CitaviCategory: "t1",
                    CitaviFilter: [] as string[],
                    BibFields: [
                        {
                            Name: "test1",
                            CitaviMapping: ["attr1"],
                            Format: {
                                Preformatted: false
                            }
                        }
                    ]
                } as Category
            ];

            const e: CitaviEntry = {
                Key: "test",
                Category: "t1",
                Attributes: [
                    {Attr: "attr1", Value: "test"}
                ]
            };

            const expectedEntry: Entry = {
                Key: "test",
                Category: "c2",
                Fields: [
                    "test"
                ]
            };

            expect(AssignCategory(e, cs)).toEqual(expectedEntry);
        });
    });
});
