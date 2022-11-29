import ProjectData from "@/pages/app/api/projects/ProjectData";

export default async function CreateProject(name: string): Promise<ProjectData> {
  const resp = await fetch('/app/createProject', {
    method: 'PUT',
    body: name
  });

  if( resp.ok ) {
    return await resp.json() as ProjectData
  } else {
    return {Name: '', Created: '', LastModified: '', NumberOfEntries: -1}
  }
}
