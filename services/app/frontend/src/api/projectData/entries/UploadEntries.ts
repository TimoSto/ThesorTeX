import {BibEntry} from "./BibEntry";
import SaveResponse from "../SaveResponse";

interface UploadEntriesData {
    Project: string
    Entries: BibEntry[]
}

export default async function UploadEntries(project: string, entries: BibEntry[]): Promise<SaveResponse> {
    const obj: UploadEntriesData = {
        Project: project,
        Entries: entries
    };

    const data = JSON.stringify(obj)

    const resp = await fetch('/app/uploadEntries', {
        method: 'PUT',
        body: data
    });

    const respData = await resp.json();

    return {
        Ok: resp.ok,
        Data: respData
    };
}
