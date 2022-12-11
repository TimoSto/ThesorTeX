import {BibCategory, Field} from "@/pages/app/api/projectData/categories/BibCategory";

export default function CheckCategoryChanged(initial: BibCategory, name: string, citaviCategory: string, citaviFilter: string[], fields: Field[], citeFields: Field[]) {
  if( !citaviFilter ) {
    citaviFilter = [];
  }
  if( !initial.CitaviNecessaryFields ) {
    initial.CitaviNecessaryFields = [];
  }
  return name !== initial.Name ||
    citaviCategory !== initial.CitaviType ||
    citaviFilter.length !== initial.CitaviNecessaryFields.length ||
    !citaviFilter.every((v, i) => v === initial.CitaviNecessaryFields[i]) ||
    fields.length !== initial.Fields.length ||
    !fields.every((f, i) => JSON.stringify(f) === JSON.stringify(initial.Fields[i])) ||
    citeFields.length !== initial.CiteFields.length ||
    !citeFields.every((f, i) => JSON.stringify(f) === JSON.stringify(initial.CiteFields[i]));
}
