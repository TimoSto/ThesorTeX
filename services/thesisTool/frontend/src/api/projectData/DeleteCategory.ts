import {host} from "../config";

export default async function DeleteCategory(project: string, name: string): Promise<boolean> {
    const resp = await fetch(`${host}/deleteCategory?project=${project}&name=${name}`, {
        method: "DELETE"
    });

    return resp.ok;
}
