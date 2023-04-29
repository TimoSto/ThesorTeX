import {i18nKeys} from "../../i18n/keys";

export default function getPortRules(t: (ik: string) => string): (port: string) => boolean | string {
    const fn = (v: string) => {
        if (v === "") {
            return t(i18nKeys.Rules.NoEmpty);
        }

        if (!/^\d+$/.test(v)) {
            return t(i18nKeys.Config.PortNumInvalid);
        }

        const num = parseInt(v);
        if (num < 1024 || num > 49151) {
            return t(i18nKeys.Config.PortReserved);
        }


        return true;
    };

    return fn;
}
