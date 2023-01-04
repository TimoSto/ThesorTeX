import {BibCategory, Field} from "./BibCategory";

export default function CheckCategoryChanged(initial: BibCategory, name: string, citaviCategory: string, citaviFilter: string[], fields: Field[], citeFields: Field[]) {
  if( !citaviFilter ) {
    citaviFilter = [];
  }
  if( !initial.CitaviFilters ) {
    initial.CitaviFilters = [];
  }
  return name !== initial.Name ||
    citaviCategory !== initial.CitaviCategory ||
    citaviFilter.length !== initial.CitaviFilters.length ||
    !citaviFilter.every((v, i) => v === initial.CitaviFilters[i]) ||
    fields.length !== initial.Fields.length ||
    !fields.every((f, i) => JSON.stringify(f) === JSON.stringify(initial.Fields[i])) ||
    citeFields.length !== initial.CiteFields.length ||
    !citeFields.every((f, i) => JSON.stringify(f) === JSON.stringify(initial.CiteFields[i]));
}
