import {BibCategory, Field} from "@/pages/app/api/projectData/categories/BibCategory";

export default function GenerateModelForCategory(category: BibCategory): string {
  let model = '';

  category.Fields.forEach((f: Field) => {
    let addition = f.Field;
    if( f.Style === 'italic' ) {
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
