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

export async function GetThesisTemplateDocumentation(lang: string): Promise<Documentation[]> {
    const resp = await fetch(`/documentation?doc=thesisTemplate&lang=${lang}`);

    if (resp.ok) {
        return await resp.json();
    }

    return [] as Documentation[];
}

export async function GetThesisToolDocumentation(lang: string): Promise<Documentation> {
    const resp = await fetch(`/documentation?doc=thesis_tool&lang=${lang}`);

    if (resp.ok) {
        return await resp.json();
    }

    return {} as Documentation;
}
