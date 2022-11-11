import {GetStringDifference} from "@/pages/auth/components/MaskValue";

describe("Handling masked input", () => {
    describe("GetDiff", () => {
        it('should work on adding to empty', () => {
            expect(GetStringDifference('', 'a')).toEqual({
                Removed: 0,
                Changed: [],
                Added: ['a']
            });
        })
        it('should work on adding to existing', () => {
            expect(GetStringDifference('a', 'ab')).toEqual({
                Removed: 0,
                Changed: [],
                Added: ['b']
            });
        })
    })
})
