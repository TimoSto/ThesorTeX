import ProjectMetaData from "../../domain/projects/ProjectMetaData";
import {host} from "../config";

interface GetAllProjectsResponse {
    Ok: boolean
    Projects: ProjectMetaData[]
}

export default async function GetProjectsMetaData(): Promise<GetAllProjectsResponse> {
    const resp = await fetch(host + "/getAllProjects");

    if( !resp.ok ) {
        return {
            Ok: false,
            Projects: []
        }
    }

    const projects = await resp.json() as ProjectMetaData[];

    return {
        Ok: true,
        Projects: projects
    }
}
