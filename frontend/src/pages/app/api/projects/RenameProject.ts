import ProjectOverviewData from "@/pages/app/api/projects/ProjectOverviewData";

export default async function RenameProject(initial: string, data: ProjectOverviewData): Promise<boolean> {
  const resp = await fetch('/app/renameProject', {
    method: 'POST',
    body: JSON.stringify({
      Project: initial,
      Data: data
    })
  });

  return resp.ok;
}
