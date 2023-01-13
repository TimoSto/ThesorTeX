import ProjectMetaData from "./ProjectMetaData";

export interface CreateProjectResponse {
  Project: ProjectMetaData,
  Status: number
}

export default async function CreateProject(name: string): Promise<CreateProjectResponse> {
  const resp = await fetch('/app/createProject', {
    method: 'PUT',
    body: name
  });

  if( resp.ok ) {
    const obj = await resp.json() as ProjectMetaData;
    console.log(obj)
    return {
      Project: obj,
      Status: 200
    }
  } else {
    return {Project: {} as ProjectMetaData, Status: resp.status}
  }
}
