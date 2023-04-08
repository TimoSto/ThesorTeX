export type ThesorTeXDocumentation = {
    ThesisTemplate: DocumentationPack
}

export type DocumentationPack = {
    Title: string
    Docs: Documentation[]
}

export type Documentation = {
    Title: string
    Groups: Group[]
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
        ThesisTemplate: {} as DocumentationPack
    };

    if (resp.ok) {
        let data = await resp.json();
        docs.ThesisTemplate = JSON.parse(data.ThesisTemplate);
    }

    return docs;
}

export async function GetThesisToolDocumentation(lang: string): Promise<Documentation> {
    const resp = await fetch(`/documentation?doc=thesis_tool&lang=${lang}`);

    if (resp.ok) {
        return await resp.json();
    }

    return {} as Documentation;
}
