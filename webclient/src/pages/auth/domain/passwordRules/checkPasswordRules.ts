import {i18nDictionary} from "@/pages/auth/i18n/Keys";

export default function checkPasswordRules(password: string): boolean|string {

    if( password.length < 8 ) {
        return i18nDictionary.Register.PasswordRules.Length;
    }

    //check if has number
    if( !/\d/g.test(password) ) {
        return i18nDictionary.Register.PasswordRules.Contains;
    }

    //chck if has symbol
    if( !/[`!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?~]/.test(password) ) {
        return i18nDictionary.Register.PasswordRules.Contains;
    }

    //chck if has letter
    if( !/[a-zA-Z]/.test(password) ) {
        return i18nDictionary.Register.PasswordRules.Contains;
    }

    return true
}
