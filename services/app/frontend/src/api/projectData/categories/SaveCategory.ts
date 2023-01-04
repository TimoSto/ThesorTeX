import {Field} from "./BibCategory";
import SaveResponse from "../SaveResponse";

interface CategorySaveData {
  Project: string
  Name: string
  InitialName: string,
  CitaviCategory: string,
  CitaviFilter: string[],
  BibFields: Field[],
  CiteFields: Field[]

}
export default async function SaveCategory(project: string, name: string, initialName: string, citaviCategory: string, citaviFilter: string[], bibFields: Field[], citeFields: Field[]): Promise<SaveResponse> {
  const obj: CategorySaveData = {
    Project: project,
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


  const respData = await resp.json();

  return {
    Ok: resp.ok,
    Data: respData
  };
}
