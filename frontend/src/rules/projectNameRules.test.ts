import getProjectNameRules from "@/rules/projectNameRules";
import {i18nKeys} from "@/pages/app/i18n/keys";

describe('ProjectNameRules', () => {
  it('should throw error on space', () => {
    expect(getProjectNameRules([])[0]('t t')).toEqual(i18nKeys.Rules.NoSpaces)
  });
  it('should throw error on existing', () => {
    expect(getProjectNameRules(['t'])[0]('t')).toEqual(i18nKeys.Rules.ProjectAlreadyExists)
  });
  it('should throw error on space', () => {
    expect(getProjectNameRules(['t'])[0]('tt')).toBe(true);
  })
})
