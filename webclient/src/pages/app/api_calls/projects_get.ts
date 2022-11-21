import {MutationTypes} from "@/pages/app/store/mutation-types";
import {ApiResponse} from "@/pages/app/api_calls/response";
import {ProjectData} from "@/pages/app/store/state";

export interface GetProjectsResponse extends ApiResponse{
    Projects: ProjectData[]
}

export default async function GetProjects(): Promise<GetProjectsResponse> {
    const resp = await fetch("/app/projects");

    if( resp.ok ) {
        const data = await resp.json() as ProjectData[];
        return {
            Success: true,
            Projects: data
        }
    }
    return {
        Success: false,
        Projects: []
    }
}
