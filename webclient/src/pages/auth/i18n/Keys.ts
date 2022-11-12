
export type TranslationKeys = typeof i18nDictionary

export const i18nDictionary = {
    Common: {
        EMail: 'Common.EMail',
        Password: 'Common.Password',
        Register: 'Common.Register'
    },
    Login: {
        Title: 'Login.Title',
        Login: 'Login.Login',
        Forgot: 'Login.Forgot',
        NotRegistered: 'Login.NotRegistered',
        ContinueWithoutLogin: 'Login.ContinueWithoutLogin'
    },
    Register: {
        Title: 'Register.Title',
        RepeatPassword: 'Register.RepeatPassword',
        PasswordRules: {
            Length: 'Register.PasswordRules.Length',
            Contains: 'Register.PasswordRules.Contains',
        }
    }
}