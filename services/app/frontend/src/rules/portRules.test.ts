import GetPortRules from "./portRules";
import {i18nKeys} from "../i18n/keys";

const t = (key: string) => key
describe('PortRules', () => {
  it('should fail on empty', () => {
    expect(GetPortRules(t)[0]('')).toEqual(i18nKeys.Rules.NotEmpty)
  })
  it('should fail on space', () => {
    expect(GetPortRules(t)[0]('8 8')).toEqual(i18nKeys.Rules.NoSpaces)
  })
  it('should fail on letter', () => {
    expect(GetPortRules(t)[0]('8a8b')).toEqual(i18nKeys.Rules.InvalidPortNumber)
  })
  it('should fail on port below 1024', () => {
    expect(GetPortRules(t)[0]('1023')).toEqual(i18nKeys.Rules.InvalidPortNumber)
  })
  it('should fail on port above 49151', () => {
    expect(GetPortRules(t)[0]('49152')).toEqual(i18nKeys.Rules.InvalidPortNumber)
  })
})
