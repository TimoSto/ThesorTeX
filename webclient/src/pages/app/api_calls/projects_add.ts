import {ApiResponse} from "@/pages/app/api_calls/response";
import {ProjectData} from "@/pages/app/store/state";

export interface AddProjectsResponse extends ApiResponse{
    Project: ProjectData
}

export default async function AddProject(name: string): Promise<AddProjectsResponse> {
    const resp = await fetch("/app/createProject", {
        method: "PUT",
        body: name
    });

    if( resp.ok ) {
        return {
            Success: resp.ok,
            Project: await resp.json()
        }
    } else {
        return {
            Success: false,
            Project: {} as ProjectData
        }
    }
}
