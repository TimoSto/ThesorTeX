import {ThesorTeXDocumentation} from "./GetDocumentation";


export default async function GetDSGVO(): Promise<ThesorTeXDocumentation> {
    const resp = await fetch(`/dsgvo`, {
        headers: {
            "Accept": "application/json"
        }
    });

    const data = await resp.json();

    return data;
}
