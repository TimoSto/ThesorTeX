import {host} from "../config";

export default async function DeleteCategory(name: string): Promise<boolean> {
    const resp = await fetch(`${host}/deleteCategory?name=${name}`, {
        method: "DELETE"
    });

    return resp.ok;
}
