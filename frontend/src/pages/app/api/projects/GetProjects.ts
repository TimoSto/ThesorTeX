import ProjectMetaData from "@/pages/app/api/projects/ProjectMetaData";

export async function GetProjects(): Promise<ProjectMetaData[]> {
  const resp = await fetch('/app/projects');

  if( resp.ok ) {
    return await resp.json() as ProjectMetaData[]
  }
  return []
}
