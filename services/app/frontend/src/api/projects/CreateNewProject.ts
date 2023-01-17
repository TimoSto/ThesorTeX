import ProjectMetaData from "../../domain/projects/ProjectMetaData";
import {host} from "../config";


interface ProjectCreationData {
    Name: string
}

interface ProjectCreationResponse {
    Success: boolean
    Data: ProjectMetaData
}

export default async function CreateNewProject(name: string): Promise<ProjectCreationResponse> {
    const data: ProjectCreationData = {
        Name: name
    }
    const resp = await fetch(`${host}/createNewProject`, {
        method: "POST",
        body: JSON.stringify(data)
    });

    if( !resp.ok ) {
        return {
            Success: false,
        } as ProjectCreationResponse
    }

    return {
        Success: true,
        Data: await resp.json()
    }
}
