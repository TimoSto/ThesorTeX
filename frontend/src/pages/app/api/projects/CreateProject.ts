import ProjectOverviewData from "@/pages/app/api/projects/ProjectOverviewData";

export interface CreateProjectResponse {
  Project: ProjectOverviewData,
  Status: number
}

export default async function CreateProject(name: string): Promise<CreateProjectResponse> {
  const resp = await fetch('/app/createProject', {
    method: 'PUT',
    body: name
  });

  if( resp.ok ) {
    const obj = await resp.json() as ProjectOverviewData;
    return {
      Project: obj,
      Status: 200
    }
  } else {
    return {Project: {} as ProjectOverviewData, Status: resp.status}
  }
}
