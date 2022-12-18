import GetEntryKeyRules from "@/pages/app/rules/entryKeyRules";
import {i18nKeys} from "@/pages/app/i18n/keys";

describe('EntryKeyRules', () => {
  it('should fail on existing with empty initial key', () => {
    const existing = ['t1','t2'];
    expect(GetEntryKeyRules(existing, '', (k: string) => k)[0]('t1')).toEqual(i18nKeys.Rules.KeyAlreadyExists)
  })
  it('should fail on existing with different initial key', () => {
    const existing = ['t1','t2'];
    expect(GetEntryKeyRules(existing, 't2', (k: string) => k)[0]('t1')).toEqual(i18nKeys.Rules.KeyAlreadyExists)
  })
  it('should fail on empty', () => {
    const existing = ['t1','t2'];
    expect(GetEntryKeyRules(existing, 't1', (k: string) => k)[0]('')).toEqual(i18nKeys.Rules.NotEmpty)
  })

  it('should not fail on existing with same initial key', () => {
    const existing = ['t1','t2'];
    expect(GetEntryKeyRules(existing, 't1', (k: string) => k)[0]('t1')).toEqual(true)
  })
  it('should not fail on new', () => {
    const existing = ['t1','t2'];
    expect(GetEntryKeyRules(existing, 't1', (k: string) => k)[0]('t3')).toEqual(true)
  })
})
