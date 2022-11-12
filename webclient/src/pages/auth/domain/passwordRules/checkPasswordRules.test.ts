import checkPasswordRules from "@/pages/auth/domain/passwordRules/checkPasswordRules";
import {i18nDictionary} from "@/pages/auth/i18n/Keys";

describe('checkPasswordRules', () => {
    describe('length', () => {
        it('should fail on 7 chars', () => {
            expect(checkPasswordRules('ab!1234')).toEqual(i18nDictionary.Register.PasswordRules.Length);
        });
    });
    describe('char, number and symbol', () => {
        it('should fail on missing number', () => {
            expect(checkPasswordRules('abchefg!')).toEqual(i18nDictionary.Register.PasswordRules.Contains);
        });
        it('should fail on missing symbol', () => {
            expect(checkPasswordRules('abchefg1')).toEqual(i18nDictionary.Register.PasswordRules.Contains);
        });
        it('should fail on missing letter', () => {
            expect(checkPasswordRules('_!_!_!_!_!')).toEqual(i18nDictionary.Register.PasswordRules.Contains);
        });
    })
    it('should succeed when all are met', () => {
        expect(checkPasswordRules('ab_12345')).toBe(true);
    })
})
