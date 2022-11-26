import {i18nKeys} from "@/pages/app/i18n/keys";

export default function getProjectNameRules(projectNames: string []): ((name: string) => boolean|string)[] {
  const fn = (name: string) => {
    if( projectNames.indexOf(name) >= 0 ) {
      return i18nKeys.Rules.ProjectAlreadyExists;
    }

    if( name.indexOf(' ') >= 0 ) {
      return i18nKeys.Rules.NoSpaces;
    }

    return true
  }

  return [fn]
}
