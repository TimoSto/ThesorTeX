import {i18nKeys} from "@/pages/app/i18n/keys";

export default function GetEntryKeyRules(existing: string[], initialKey: string, t: (ik: string) => string): ((key: string) => boolean|string)[] {
  const fn = (key: string) => {
    const i = existing.indexOf(key)
    if( i >= 0 && existing[i] != initialKey ) {
      return t(i18nKeys.Rules.KeyAlreadyExists)
    }
    return true
  }

  return [fn]
}
