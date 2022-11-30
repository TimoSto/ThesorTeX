import ProjectData from "@/pages/app/api/projects/ProjectData";

export interface CreateProjectResponse {
  Project: ProjectData,
  Status: number
}

export default async function CreateProject(name: string): Promise<CreateProjectResponse> {
  const resp = await fetch('/app/createProject', {
    method: 'PUT',
    body: name
  });

  if( resp.ok ) {
    const obj = await resp.json() as ProjectData;
    return {
      Project: obj,
      Status: 200
    }
  } else {
    return {Project: {} as ProjectData, Status: resp.status}
  }
}
