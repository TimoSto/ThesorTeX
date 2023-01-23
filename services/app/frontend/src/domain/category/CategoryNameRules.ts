import {i18nKeys} from "../../i18n/keys";

export default function getCategoryNameRules(existing: string[], initial: string, t: (ik: string) => string): (name: string) => boolean | string {
    const fn = (name: string) => {
        if (name === "") {
            return t(i18nKeys.Rules.NoEmpty);
        }

        if (name.indexOf(" ") >= 0) {
            return t(i18nKeys.Rules.NoSpaces);
        }

        const i = existing.indexOf(name);

        if (i >= 0 && existing[i] != initial) {
            return t(i18nKeys.Rules.CategoryAlreadyExists);
        }

        return true;
    };

    return fn;
}
