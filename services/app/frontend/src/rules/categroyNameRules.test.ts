import {i18nKeys} from "../i18n/keys";
import GetCategoryNameRules from "./categoryNameRules";

describe('CategoryKeyRules', () => {
  it('should fail on existing with empty initial key', () => {
    const existing = ['t1','t2'];
    expect(GetCategoryNameRules(existing, '', (k: string) => k)[0]('t1')).toEqual(i18nKeys.Rules.CategoryNameAlreadyExists)
  })
  it('should fail on existing with different initial key', () => {
    const existing = ['t1','t2'];
    expect(GetCategoryNameRules(existing, 't2', (k: string) => k)[0]('t1')).toEqual(i18nKeys.Rules.CategoryNameAlreadyExists)
  })
  it('should fail on empty', () => {
    const existing = ['t1','t2'];
    expect(GetCategoryNameRules(existing, 't1', (k: string) => k)[0]('')).toEqual(i18nKeys.Rules.NotEmpty)
  })
  it('should fail on special char _', () => {
    expect(GetCategoryNameRules([], 't1', (k: string) => k)[0]('t_e')).toEqual(i18nKeys.Rules.NoSpecialChars)
  })

  it('should not fail on existing with same initial key', () => {
    const existing = ['t1','t2'];
    expect(GetCategoryNameRules(existing, 't1', (k: string) => k)[0]('t1')).toEqual(true)
  })
  it('should not fail on new', () => {
    const existing = ['t1','t2'];
    expect(GetCategoryNameRules(existing, 't1', (k: string) => k)[0]('t3')).toEqual(true)
  })
})
