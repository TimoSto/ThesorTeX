import {Entry} from "../../domain/entry/Entry";
import {host} from "../config";

export default async function SaveEntry(project: string, key: string, entry: Entry): Promise<boolean> {
    const resp = await fetch(`${host}/saveEntry?project=${project}&key=${key}`, {
        method: "POST",
        body: JSON.stringify(entry)
    });

    return resp.ok;
}
