import {DocumentationPack} from "./GetDocumentation";


export default async function GetDSGVO(): Promise<DocumentationPack> {
    const resp = await fetch(`/dsgvo`, {
        headers: {
            "Accept": "application/json"
        }
    });

    const data = await resp.json();

    return data;
}
