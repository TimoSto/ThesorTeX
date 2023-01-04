import ProjectMetaData from "./ProjectMetaData";

export default async function RenameProject(initial: string, data: ProjectMetaData): Promise<boolean> {
  const resp = await fetch('/app/renameProject', {
    method: 'POST',
    body: JSON.stringify({
      Project: initial,
      Data: data
    })
  });

  return resp.ok;
}
