import {i18nKeys} from "../../i18n/keys";

export default function getAttributeNameRules(t: (ik: string) => string): (name: string) => boolean | string {
    const fn = (name: string) => {
        if (name === "") {
            return t(i18nKeys.Rules.NoEmpty);
        }

        return true;
    };

    return fn;
}
