import {i18nKeys} from "@/pages/app/i18n/keys";
import GetFieldNameRules from "@/pages/app/rules/fieldNameRules";

describe('FieldNameRules', () => {

  it('should fail on existing', () => {
    const existing = ['t1','t1'];
    expect(GetFieldNameRules(existing,  (k: string) => k)[0]('t1')).toEqual(i18nKeys.Rules.FieldAlreadyExists)
  })
  it('should fail on empty', () => {
    const existing = ['t1','t2'];
    expect(GetFieldNameRules(existing, (k: string) => k)[0]('')).toEqual(i18nKeys.Rules.NotEmpty)
  })

  it('should not fail on existing with same initial key', () => {
    const existing = ['t1','t2'];
    expect(GetFieldNameRules(existing, (k: string) => k)[0]('t1')).toEqual(true)
  })
  it('should not fail on new', () => {
    const existing = ['t1','t2'];
    expect(GetFieldNameRules(existing, (k: string) => k)[0]('t3')).toEqual(true)
  })
})
