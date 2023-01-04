import {Field} from "./BibCategory";

export default function GenerateModelForFields(fields: Field[]): string {
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
