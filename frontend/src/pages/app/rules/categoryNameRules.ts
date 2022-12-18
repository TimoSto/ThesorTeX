import {i18nKeys} from "@/pages/app/i18n/keys";
import StringContainsSpecialChars from "@/pages/app/rules/specialCharRegex";

export default function GetCategoryNameRules(existing: string [], initialName: string, t: (key: string) => string): ((name: string) => boolean|string)[] {
  const fn = function(n: string) {
    if( n === '' ) {
      return t(i18nKeys.Rules.NotEmpty);
    }
    if( StringContainsSpecialChars(n) ) {
      return t(i18nKeys.Rules.NoSpecialChars)
    }
    const i = existing.indexOf(n)
    if( i >= 0 && existing[i] != initialName ) {
      return t(i18nKeys.Rules.CategoryNameAlreadyExists);
    }

    return true
  }

  return [fn]
}
