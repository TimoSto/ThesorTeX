import {describe, expect, it} from "vitest";
import {i18nKeys} from "../../i18n/keys";
import getAttributeNameRules from "./AttributeNameRules";

describe("AttributeNameRules", () => {
    it("should fail on empty", () => {
        const rules = getAttributeNameRules(k => k);
        expect(rules("")).toEqual(i18nKeys.Rules.NoEmpty);
    });
    it("should succeed on valid", () => {
        const rules = getAttributeNameRules(k => k);
        expect(rules("te")).toBe(true);
    });
});
