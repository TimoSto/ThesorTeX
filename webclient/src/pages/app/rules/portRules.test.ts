import CheckPortValid from "@/pages/app/rules/portRules";

describe('PortRules', () => {
    it('should fail on chars', () => {
        expect(typeof CheckPortValid("808a")).toEqual('string')
    })
})
