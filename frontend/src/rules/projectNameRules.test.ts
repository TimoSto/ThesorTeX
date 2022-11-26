import getProjectNameRules from "@/rules/projectNameRules";
import {i18nKeys} from "@/pages/app/i18n/keys";

describe("projectNameRules", () => {
  it('should fail on space in name', () => {
    expect(getProjectNameRules([])[0]('hallo test')).toEqual(i18nKeys.Rules.NoSpaces)
  })
})
