import {describe, expect, it} from "vitest";
import {i18nKeys} from "../../i18n/keys";
import getPortRules from "./PortRules";

describe("AttributeNameRules", () => {
    it("should fail on empty", () => {
        const rules = getPortRules(k => k);
        expect(rules("")).toEqual(i18nKeys.Rules.NoEmpty);
    });
    it("should fail on char in num", () => {
        const rules = getPortRules(k => k);
        expect(rules("7134f")).toEqual(i18nKeys.Config.PortNumInvalid);
    });
    it("should fail on reserved port", () => {
        const rules = getPortRules(k => k);
        expect(rules("65")).toEqual(i18nKeys.Config.PortReserved);
    });
    it("should succeed on valid", () => {
        const rules = getPortRules(k => k);
        expect(rules("8334")).toBe(true);
    });
});
