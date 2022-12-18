import {i18nKeys} from "@/pages/app/i18n/keys";

export default function GetFieldNameRules(existing: string[], initialName: string, t: (ik: string) => string): ((key: string) => boolean|string)[] {
  const fn = (f: string) => {
    if( f === '' ) {
      return t(i18nKeys.Rules.NotEmpty);
    }
    if( f.indexOf(' ') >= 0 ) {
      return t(i18nKeys.Rules.NoSpaces);
    }

    const i = existing.indexOf(f)
    if( i >= 0 && i < existing.length - 1 ) {
      return t(i18nKeys.Rules.FieldAlreadyExists);
    }

    return true
  }

  return [fn]
}
