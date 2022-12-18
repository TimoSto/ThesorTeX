import GetEntryKeyRules, {existingKeyError} from "@/rules/entryKeyRules";

describe('EntryKeyRules', () => {
  it('should fail on existing with empty initial key', () => {
    const existing = ['t1','t2'];
    expect(GetEntryKeyRules(existing, '')[0]('t1')).toEqual(existingKeyError)
  })
  it('should fail on existing with different initial key', () => {
    const existing = ['t1','t2'];
    expect(GetEntryKeyRules(existing, 't2')[0]('t1')).toEqual(existingKeyError)
  })
  it('should not fail on existing with same initial key', () => {
    const existing = ['t1','t2'];
    expect(GetEntryKeyRules(existing, 't1')[0]('t1')).toEqual(true)
  })
  it('should not fail on new', () => {
    const existing = ['t1','t2'];
    expect(GetEntryKeyRules(existing, 't1')[0]('t3')).toEqual(true)
  })
})
