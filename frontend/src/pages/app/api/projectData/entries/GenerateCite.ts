import {BibEntry} from "@/pages/app/api/projectData/entries/BibEntry";
import {BibCategory, Field} from "@/pages/app/api/projectData/categories/BibCategory";

export default function GenerateCite(entryObj: BibEntry, categories: BibCategory[]): string {
  let cite = '';

  const i = categories.map(c => c.Name).indexOf(entryObj.Category);

  if( i >= 0 ) {
    let found = 0;
    categories[i].CiteFields.forEach((f: Field, n: number) => {
      console.log(f.Field, f.Prefix, f.Suffix)
      cite += f.Prefix ? f.Prefix : '';
      if(f.Italic) {
        cite += '<i>';
      }
      const indexInBibFields = categories[i].Fields.map(bf => bf.Field).indexOf(f.Field)
      if( indexInBibFields >= 0 ) {
        cite += (indexInBibFields < entryObj.Fields.length ? entryObj.Fields[indexInBibFields] : '')
        found++;
      } else {
        n += categories[i].Fields.length - found
        cite += (n < entryObj.Fields.length ? entryObj.Fields[n] : '')
      }
      if(f.Italic ) {
        cite += '</i>';
      }
      cite += f.Suffix ? f.Suffix : '';
    });
  }

  return cite
}
