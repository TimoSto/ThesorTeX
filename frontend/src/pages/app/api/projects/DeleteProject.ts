
export default async function DeleteProject(name: string): Promise<boolean> {
  const resp = await fetch('/app/deleteProject', {
    method: 'DELETE',
    body: name
  });

  return resp.ok;
}
