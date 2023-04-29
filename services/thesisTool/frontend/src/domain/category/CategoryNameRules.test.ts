import {describe, expect, it} from "vitest";
import {i18nKeys} from "../../i18n/keys";
import getCategoryNameRules from "./CategoryNameRules";

describe("CategoryNameRules", () => {
    it("should fail on empty", () => {
        const rules = getCategoryNameRules([], "", k => k);
        expect(rules("")).toEqual(i18nKeys.Rules.NoEmpty);
    });
    it("should fail on space", () => {
        const rules = getCategoryNameRules([], "", k => k);
        expect(rules("t s")).toEqual(i18nKeys.Rules.NoSpaces);
    });
    it("should fail on existing", () => {
        const rules = getCategoryNameRules(["test"], "", k => k);
        expect(rules("test")).toEqual(i18nKeys.Rules.NameAlreadyExists);
    });
    it("should succeed on valid", () => {
        const rules = getCategoryNameRules(["test"], "", k => k);
        expect(rules("te")).toBe(true);
    });
});
