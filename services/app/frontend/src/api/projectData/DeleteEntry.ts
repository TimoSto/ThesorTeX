import {host} from "../config";

export default async function DeleteEntry(project: string, key: string): Promise<boolean> {
    const resp = await fetch(`${host}/deleteEntry?project=${project}&key=${key}`, {
        method: "DELETE"
    });

    return resp.ok;
}
