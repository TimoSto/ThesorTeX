import {TranslationKeys} from "./Keys";

export const EnglishTranslations: TranslationKeys = {
    Common: {
        EMail: 'Email',
        Password: 'Password',
        Register: 'Register'
    },
    Login: {
        Title: 'ThesorTeX - Login',
        Login: 'Login',
        Forgot: 'Forgot password',
        NotRegistered: 'Not registered yet?',
        ContinueWithoutLogin: 'Continue without login'
    },
    Register: {
        Title: 'Registration',
        RepeatPassword: 'Repeat password',
        PasswordRules: {
            Length: 'Passwort muss mindestens 8 Zeichen beinhalten',
            Contains: 'Passwort muss mindestens 1 Buchstaben, 1 Zahl und 1 Sonderzeichen beinhalten',
            Equal: 'Passwords are unequal'
        }
    }
}