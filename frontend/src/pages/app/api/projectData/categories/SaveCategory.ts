import {Field} from "@/pages/app/api/projectData/categories/BibCategory";

interface CategorySaveData {
  Name: string
  InitialName: string,
  CitaviCategory: string,
  CitaviFilter: string[],
  BibFields: Field[],
  CiteFields: Field[]

}
export default async function SaveCategory(name: string, initialName: string, citaviCategory: string, citaviFilter: string[], bibFields: Field[], citeFields: Field[]): Promise<boolean> {
  const obj: CategorySaveData = {
    InitialName: initialName,
    Name: name,
    CitaviCategory: citaviCategory,
    CitaviFilter: citaviFilter,
    BibFields: bibFields,
    CiteFields: citeFields
  }

  const data = JSON.stringify(obj)

  const resp = await fetch('/app/saveCategory', {
    method: 'PUT',
    body: data
  });

  return resp.ok
}
