import {BibEntry} from "@/pages/app/api/projectData/entries/BibEntry";
import {BibCategory, Field} from "@/pages/app/api/projectData/categories/BibCategory";

export default function GenerateEntryExample(entryObj: BibEntry, categories: BibCategory[]): string {
  let entry = '';

  const i = categories.map(c => c.Name).indexOf(entryObj.Category);

  if( i >= 0 ) {
    categories[i].Fields.forEach((f: Field, n: number) => {
      entry += f.Prefix ? f.Prefix : '';
      if(f.Style === 'italic') {
        entry += '<i>';
      }
      entry += (n < entryObj.Fields.length ? entryObj.Fields[n] : '')
      if(f.Style === 'italic') {
        entry += '</i>';
      }
      entry += f.Suffix ? f.Suffix : '';
    });
  }

  return entry
}
