import {Entry} from "../../domain/entry/Entry";
import {host} from "../config";

export default async function UploadEntries(project: string, entries: Entry[]): Promise<boolean> {
    const resp = await fetch(`${host}/uploadEntries?project=${project}`, {
        method: "POST",
        body: JSON.stringify(entries)
    });

    return resp.ok;
}
