import getProjectNameRules from "@/rules/projectNameRules";
import {i18nKeys} from "@/pages/app/i18n/keys";

describe('ProjectNameRules', () => {
  it('should throw error', () => {
    expect(getProjectNameRules([])[0]('t t')).toEqual(i18nKeys.Rules.NoSpaces)
  })
})
