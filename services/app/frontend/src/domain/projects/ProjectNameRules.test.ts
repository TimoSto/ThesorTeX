import {describe, expect, it} from "vitest";
import getProjectNameRules from "./ProjectNameRules";
import {i18nKeys} from "../../i18n/keys";

describe("ProjectNameRules", () => {
    it("should fail on empty", () => {
        const rules = getProjectNameRules([], "", k => k);
        expect(rules("")).toEqual(i18nKeys.Rules.NoEmpty);
    });
    it("should fail on space", () => {
        const rules = getProjectNameRules([], "", k => k);
        expect(rules("t s")).toEqual(i18nKeys.Rules.NoSpaces);
    });
    it("should fail on existing", () => {
        const rules = getProjectNameRules(["test"], "", k => k);
        expect(rules("test")).toEqual(i18nKeys.Rules.ProjectAlreadyExists);
    });
    it("should succeed on valid", () => {
        const rules = getProjectNameRules(["test"], "", k => k);
        expect(rules("te")).toBe(true);
    });
})
