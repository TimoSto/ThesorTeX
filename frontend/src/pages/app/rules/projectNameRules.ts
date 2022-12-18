import {i18nKeys} from "@/pages/app/i18n/keys";

// eslint-disable-next-line no-unused-vars
export default function getProjectNameRules(projectNames: string [], initial: string, t: (ik: string) => string): ((name: string) => boolean|string)[] {
  const fn = (name: string) => {

    if( name === '' ) {
      return t(i18nKeys.Rules.NotEmpty)
    }

    const i = projectNames.indexOf(name);

    if( i >= 0 && projectNames[i] != initial ) {
      return t(i18nKeys.Rules.ProjectAlreadyExists);
    }

    return true
  }

  return [fn]
}
