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
        it('should work on removing last', () => {
            expect(GetStringDifference('ab', 'a')).toEqual({
                Removed: 1,
                Changed: [],
                Added: []
            });
        })
        it('should work on removing first', () => {
            expect(GetStringDifference('acb', 'cb')).toEqual({
                Removed: 1,
                Changed: [
                    {
                        Index: 0,
                        Old: 'a',
                        New: 'c'
                    },
                    {
                        Index: 1,
                        Old: 'c',
                        New: 'b'
                    }
                ],
                Added: []
            });
        })
    })
})
