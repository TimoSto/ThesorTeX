import {Entry} from "../../domain/entry/Entry";
import {host} from "../config";
import {Category} from "../../domain/category/Category";

export interface ProjectData {
    Entries: Entry[]
    Categories: Category[]
}

export interface GetProjectDataResponse {
    Ok: boolean
    Data: ProjectData
}

export default async function GetProjectData(project: string): Promise<GetProjectDataResponse> {
    const resp = await fetch(`${host}/getProjectData?project=${project}`);
    //TODO: when here is more logic, add a unittest

    if (!resp.ok) {
        return {
            Ok: false,
            Data: {} as ProjectData
        }
    }

    return {
        Ok: resp.ok,
        Data: await resp.json()
    }
}
