import {describe, expect, it} from "vitest";
import {ReadFile} from "../testdata/testdataReader";
import {AnalyseBibFile} from "./BibAnalytics";

describe("BibAnalytics", () => {
    describe("AnalyseBibFile", () => {
        it("should give one entry", () => {
            const file = ReadFile("files/citavi/oneEntry.bib");

            const citaviEntries = AnalyseBibFile(file);

            expect(citaviEntries[0].Key).toEqual("f1");
            expect(citaviEntries[0].Category).toEqual("book");
            expect(citaviEntries[0].Attributes[0].Attr).toEqual("field1");
            expect(citaviEntries[0].Attributes[0].Value).not.toEqual("");
            expect(citaviEntries[0].Attributes[1].Attr).toEqual("field2");
            expect(citaviEntries[0].Attributes[1].Value).not.toEqual("");
        });
        it("should give two entries", () => {
            const file = ReadFile("files/citavi/twoEntriesNewline.bib");

            const citaviEntries = AnalyseBibFile(file);

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

            const citaviEntries = AnalyseBibFile(file);

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
});
