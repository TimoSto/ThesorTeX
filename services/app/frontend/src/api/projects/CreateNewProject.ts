import ProjectMetaData from "../../domain/projects/ProjectMetaData";
import {host} from "../config";


interface ProjectCreationData {
    Name: string
}

interface ProjectCreationResponse {
    Success: boolean
    Data: ProjectMetaData
    Status: number
}

export default async function CreateNewProject(name: string): Promise<ProjectCreationResponse> {
    const data: ProjectCreationData = {
        Name: name
    }
    const resp = await fetch(`${host}/createNewProject`, {
        method: "PUT",
        body: JSON.stringify(data)
    });

    if( !resp.ok ) {
        return {
            Success: false,
            Status: resp.status
        } as ProjectCreationResponse
    }

    return {
        Success: true,
        Data: await resp.json(),
        Status: resp.status
    }
}
