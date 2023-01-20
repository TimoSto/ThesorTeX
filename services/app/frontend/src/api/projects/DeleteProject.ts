import {host} from "../config";

export default async function DeleteProject(project: string): Promise<boolean> {
    const resp = await fetch(`${host}/deleteProject?project=${project}`, {
        method: "DELETE"
    });

    return resp.ok;
}
