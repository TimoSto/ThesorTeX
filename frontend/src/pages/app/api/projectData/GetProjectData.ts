import {ProjectData} from "@/pages/app/api/projectData/ProjectData";

export default async function GetProjectData(project: string): Promise<ProjectData> {
  const resp = await fetch(`/app/projectData?project=${project}`);

  //todo: handle error, possibly via additiona attribute in Project data?
  if( resp.ok ) {
    return await resp.json() as ProjectData
  }
  return {
    Entries: [],
    Categories: []
  }
}
