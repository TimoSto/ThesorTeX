import {i18nKeys} from "@/pages/app/i18n/keys";

// eslint-disable-next-line no-unused-vars
export default function getProjectNameRules(projectNames: string [], t: (ik: string) => string): ((name: string) => boolean|string)[] {
  const fn = (name: string) => {
    if( projectNames.indexOf(name) >= 0 ) {
      return t(i18nKeys.Rules.ProjectAlreadyExists);
    }

    if( name.indexOf(' ') >= 0 ) {
      return t(i18nKeys.Rules.NoSpaces);
    }

    return true
  }

  return [fn]
}
