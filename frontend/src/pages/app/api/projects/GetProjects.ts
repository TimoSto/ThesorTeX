import ProjectData from "@/pages/app/api/projects/ProjectData";

export async function GetProjects(): Promise<ProjectData[]> {
  const resp = await fetch('/app/projects');

  if( resp.ok ) {
    return await resp.json() as ProjectData[]
  }
  return []
}
