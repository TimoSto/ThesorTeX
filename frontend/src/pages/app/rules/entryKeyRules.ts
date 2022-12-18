import {i18nKeys} from "@/pages/app/i18n/keys";
import StringContainsSpecialChars from "@/pages/app/rules/specialCharRegex";

export default function GetEntryKeyRules(existing: string[], initialKey: string, t: (ik: string) => string): ((key: string) => boolean|string)[] {
  const fn = (key: string) => {
    if( key === '' ) {
      return t(i18nKeys.Rules.NotEmpty);
    }
    if( StringContainsSpecialChars(key) ) {
      return t(i18nKeys.Rules.NoSpecialChars)
    }
    const i = existing.indexOf(key)
    if( i >= 0 && existing[i] != initialKey ) {
      return t(i18nKeys.Rules.KeyAlreadyExists);
    }
    return true
  }

  return [fn]
}
