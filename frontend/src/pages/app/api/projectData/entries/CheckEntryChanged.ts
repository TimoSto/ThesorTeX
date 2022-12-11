import {BibEntry} from "@/pages/app/api/projectData/entries/BibEntry";

export default function CheckEntryChanged(initial: BibEntry, key: string, category: string, fields: string[]): boolean {
  return initial.Key !== key ||
    initial.Category !== category ||
    initial.Fields.length !== fields.length ||
    !initial.Fields.every((v, i) => v === fields[i]);
}
