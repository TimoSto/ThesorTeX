export type ThesorTeXDocumentation = {
    ThesisTemplate: DocumentationPack
    CVTemplate: DocumentationPack
    ThesisTool: DocumentationPack
}

export type DocumentationPack = {
    Title: string
    Docs: Documentation[]
}

export type Footnote = {
    Key: string
    Elements: Element[]
}

export type Documentation = {
    Title: string
    Groups: Group[]
    Footnotes: Footnote[]
}

export type Group = {
    Type: "TEXT",
    Elements: Element[]
}

export type Element = {
    Content: string,
    Style: "PLAIN" | "BOLD" | "ITALIC" | "ITALIC_BOLD"
}

export async function GetDocumentationsJSON(lang: string): Promise<ThesorTeXDocumentation> {
    const resp = await fetch(`/documentation?lang=${lang}`);

    let docs: ThesorTeXDocumentation = {
        ThesisTemplate: {} as DocumentationPack,
        CVTemplate: {} as DocumentationPack,
        ThesisTool: {} as DocumentationPack
    };

    if (resp.ok) {
        let data = await resp.json();
        docs.ThesisTemplate = mapFootnotes(data.ThesisTemplate);
        docs.CVTemplate = mapFootnotes(data.CVTemplate);
        docs.ThesisTool = mapFootnotes(data.ThesisTool);
    }

    console.log(docs.ThesisTool);

    return docs;
}

export async function GetThesisToolDocumentation(lang: string): Promise<Documentation> {
    const resp = await fetch(`/documentation?doc=thesis_tool&lang=${lang}`);

    if (resp.ok) {
        return await resp.json();
    }

    return {} as Documentation;
}

function mapFootnotes(data: string): DocumentationPack {
    const raw = JSON.parse(data);

    raw.Docs.forEach((d: Documentation, i: number) => {
        const rawFootnotes = d.Footnotes;
        raw.Docs[i].Footnotes = [];

        for (let k in rawFootnotes) {
            raw.Docs[i].Footnotes.push({
                Key: k,
                Elements: rawFootnotes[k]
            });
        }
        // for (let [k, v] of rawFootnotes) {
        //     console.log(k, v);
        // }

        console.log(raw.Docs[i].Footnotes);
    });


    return raw;
}
