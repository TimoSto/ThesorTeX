
export default async function DeleteEntry(project: string, key: string) {
  const resp = await fetch('/app/deleteEntry', {
    method: 'DELETE',
    body: JSON.stringify({
      Project: project,
      Key: key
    })
  });

  return resp.ok;
}
