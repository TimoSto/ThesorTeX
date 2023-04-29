import {Category} from "../../domain/category/Category";
import {host} from "../config";

export default async function SaveCategory(name: string, project: string, category: Category): Promise<boolean> {
    const resp = await fetch(`${host}/saveCategory?project=${project}&name=${name}`, {
        method: "POST",
        body: JSON.stringify(category)
    });

    return resp.ok;
}
