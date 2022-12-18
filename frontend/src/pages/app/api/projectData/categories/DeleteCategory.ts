
export default async function DeleteCategory(project: string, name: string) {
  const resp = await fetch('/app/deleteCategory', {
    method: 'DELETE',
    body: JSON.stringify({
      Project: project,
      Name: name
    })
  });

  return resp.ok;
}
