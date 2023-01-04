import ProjectMetaData from "./ProjectMetaData";

export async function GetProjects(): Promise<ProjectMetaData[]> {
  const resp = await fetch('/app/projects');

  if( resp.ok ) {
    return await resp.json() as ProjectMetaData[]
  }
  return []
}
