export type ThesisDoc = {
    Main: string,
    ChapterNumbering: string,
    HeaderFooter: string,
    Abbreviations: string,
    Appendix: string,
    Bibliography: string
}

export async function GetThesisDocumentation(lang: string): Promise<ThesisDoc> {
    const resp = await fetch(`/documentation?doc=thesis&lang=${lang}`);

    if (resp.ok) {
        return await resp.json();
    }

    return {} as ThesisDoc;
}
