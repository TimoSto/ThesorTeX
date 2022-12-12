import {Field} from "@/pages/app/api/projectData/categories/BibCategory";

export default function GenerateModelForCategory(fields: Field[]): string {
  let model = '';

  fields.forEach((f: Field) => {
    let addition = f.Field;
    if( f.Italic ) {
      addition = '<i>' + addition + '</i>';
    }
    if( f.Prefix ) {
      addition = f.Prefix + addition;
    }
    if( f.Suffix ) {
      addition += f.Suffix
    }

    model += addition;
  });

  return model;
}
