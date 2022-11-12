import checkPasswordRules from "@/pages/auth/domain/passwordRules/checkPasswordRules";
import {i18nDictionary} from "@/pages/auth/i18n/Keys";

describe('checkPasswordRules', () => {
    describe('length', () => {
        it('should fail on 7 chars', () => {
            expect(checkPasswordRules('ab!1234')).toEqual(i18nDictionary.Register.PasswordRules.Length);
        });
        it('should not fail on 8 chars', () => {
            expect(checkPasswordRules('ab!12345')).toBe(true);
        });
    })
})
