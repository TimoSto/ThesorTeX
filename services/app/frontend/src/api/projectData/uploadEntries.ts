import {Entry} from "../../domain/entry/Entry";
import {host} from "../config";

export default async function UploadEntries(entries: Entry[]): Promise<boolean> {
    const resp = await fetch(`${host}/uploadEntries`, {
        method: "POST",
        body: JSON.stringify(entries)
    });

    return resp.ok;
}
