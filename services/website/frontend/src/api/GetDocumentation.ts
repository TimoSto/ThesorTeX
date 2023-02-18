export default async function GetDocumentation(doc: string, lang: string): Promise<string> {
    const resp = await fetch(`/documentation?doc=${doc}&lang=${lang}`);

    if (resp.ok) {
        return await resp.text();
    }

    return "";
}
