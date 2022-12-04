import ProjectOverviewData from "@/pages/app/api/projects/ProjectOverviewData";

export async function GetProjects(): Promise<ProjectOverviewData[]> {
  const resp = await fetch('/app/projects');

  if( resp.ok ) {
    return await resp.json() as ProjectOverviewData[]
  }
  return []
}
