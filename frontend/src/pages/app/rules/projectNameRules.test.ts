import getProjectNameRules from "@/pages/app/rules/projectNameRules";
import {i18nKeys} from "@/pages/app/i18n/keys";

describe('ProjectNameRules', () => {
  it('should throw error on space', () => {
    expect(getProjectNameRules([], '', (k: string) => k)[0]('t t')).toEqual(i18nKeys.Rules.NoSpaces)
  });
  it('should throw error on existing', () => {
    expect(getProjectNameRules(['t'], '', (k: string) => k)[0]('t')).toEqual(i18nKeys.Rules.ProjectAlreadyExists)
  });
  it('should throw error on space', () => {
    expect(getProjectNameRules(['t'], '', (k: string) => k)[0]('tt')).toBe(true);
  })
})
