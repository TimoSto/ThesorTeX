export type Documentation = {
    Docs: {
        Title: string
        Content: string
    }[]
}

export async function GetThesisDocumentation(lang: string): Promise<Documentation> {
    const resp = await fetch(`/documentation?doc=thesis&lang=${lang}`);

    if (resp.ok) {
        return await resp.json();
    }

    return {} as Documentation;
}

export async function GetThesisToolDocumentation(lang: string): Promise<Documentation> {
    const resp = await fetch(`/documentation?doc=thesis_tool&lang=${lang}`);

    if (resp.ok) {
        return await resp.json();
    }

    return {} as Documentation;
}
