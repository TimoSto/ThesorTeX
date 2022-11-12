import {TranslationKeys} from "./Keys";

export const GermanTranslations: TranslationKeys = {
    Common: {
        EMail: 'E-Mail',
        Password: 'Passwort',
        Register: 'Registrieren'
    },
    Login: {
        Title: 'ThesorTeX - Anmeldung',
        Login: 'Anmelden',
        Forgot: 'Passwort vergessen',
        NotRegistered: 'Noch nicht registriert?',
        ContinueWithoutLogin: 'Ohne Anmeldung fortfahren'
    },
    Register: {
        Title: 'ThesorTeX - Registrierung',
        RepeatPassword: 'Passwort wiederholden',
        PasswordRules: {
            Length: 'Passwort muss mindestens 8 Zeichen beinhalten',
            Contains: 'Passwort muss mindestens 1 Buchstaben, 1 Zahl und 1 Sonderzeichen beinhalten',
            Equal: 'Passwörter stimmen nicht überein'
        }
    }
}