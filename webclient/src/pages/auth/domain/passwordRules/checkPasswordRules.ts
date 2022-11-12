import {i18nDictionary} from "@/pages/auth/i18n/Keys";

export default function checkPasswordRules(password: string): boolean|string {

    if( password.length < 8 ) {
        return i18nDictionary.Register.PasswordRules.Length;
    }

    return true
}
