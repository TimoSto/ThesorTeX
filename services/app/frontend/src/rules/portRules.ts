import {i18nKeys} from "../i18n/keys";

export default function GetPortRules(t: (key: string) => string): ((key: string) => boolean|string)[] {
  const fn = (f: string) => {
    if( f === '' ) {
      return t(i18nKeys.Rules.NotEmpty);
    }
    if( f.indexOf(' ') >= 0 ) {
      return t(i18nKeys.Rules.NoSpaces);
    }

    if( !/^\d+$/.test(f) ) {
      return t(i18nKeys.Rules.InvalidPortNumber);
    }

    const v = parseInt(f);

    if( v < 1024 || v > 49151 ) {
      return t(i18nKeys.Rules.InvalidPortNumber);
    }

    return true
  }

  return [fn]
}

