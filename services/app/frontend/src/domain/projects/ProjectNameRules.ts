import {i18nKeys} from "../../i18n/keys";

export default function getProjectNameRules(projectNames: string [], initial: string, t: (ik: string) => string): (name: string) => boolean|string {
    const fn = (name: string) => {

        if( name === '' ) {
            return t(i18nKeys.Rules.NoEmpty)
        }

        if( name.indexOf(' ') >= 0 ) {
            return t(i18nKeys.Rules.NoSpaces)
        }

        console.log(projectNames, name)

        const i = projectNames.indexOf(name);

        if( i >= 0 && projectNames[i] != initial ) {
            return t(i18nKeys.Rules.ProjectAlreadyExists);
        }

        return true
    }

    return fn
}

