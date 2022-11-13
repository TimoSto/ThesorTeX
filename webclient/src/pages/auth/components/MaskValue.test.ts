import MaskValue, {GetStringDifference} from "@/pages/auth/components/MaskValue";

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

    describe('getNewMaskState', () => {
        describe('input at end', () => {
            it('first input', () => {
                const res = MaskValue({
                    CurrentValue: '',
                    CurrentMask: '',
                    InputState: 'a',
                    CaretPosition: 1
                });

                expect(res).toEqual({
                    CurrentValue: 'a',
                    CurrentMask: '•',
                    InputState: '',
                    CaretPosition: 1
                })
            })
            it('second input', () => {
                const res = MaskValue({
                    CurrentValue: 'a',
                    CurrentMask: '•',
                    InputState: '•b',
                    CaretPosition: 2
                });

                expect(res).toEqual({
                    CurrentValue: 'ab',
                    CurrentMask: '••',
                    InputState: '',
                    CaretPosition: 2
                })
            })
            it('third input', () => {
                const res = MaskValue({
                    CurrentValue: 'ab',
                    CurrentMask: '••',
                    InputState: '••c',
                    CaretPosition: 3
                });

                expect(res).toEqual({
                    CurrentValue: 'abc',
                    CurrentMask: '•••',
                    InputState: '',
                    CaretPosition: 3
                })
            })
            it('multiple char input', () => {
                const res = MaskValue({
                    CurrentValue: 'ab',
                    CurrentMask: '••',
                    InputState: '••cdefg',
                    CaretPosition: 7
                });

                expect(res).toEqual({
                    CurrentValue: 'abcdefg',
                    CurrentMask: '•••••••',
                    InputState: '',
                    CaretPosition: 7
                })
            })
        })
        describe('remove from end', () => {
            it('rm second char', () => {
                const res = MaskValue({
                    CurrentValue: 'ab',
                    CurrentMask: '••',
                    InputState: '•',
                    CaretPosition: 1
                });

                expect(res).toEqual({
                    CurrentValue: 'a',
                    CurrentMask: '•',
                    InputState: '',
                    CaretPosition: 1
                })
            })
            it('rm first char', () => {
                const res = MaskValue({
                    CurrentValue: 'a',
                    CurrentMask: '•',
                    InputState: '',
                    CaretPosition: 0
                });

                expect(res).toEqual({
                    CurrentValue: '',
                    CurrentMask: '',
                    InputState: '',
                    CaretPosition: 0
                })
            })
            it('rm multiple chars', () => {
                const res = MaskValue({
                    CurrentValue: 'abcd',
                    CurrentMask: '••••',
                    InputState: '•',
                    CaretPosition: 1
                });

                expect(res).toEqual({
                    CurrentValue: 'a',
                    CurrentMask: '•',
                    InputState: '',
                    CaretPosition: 1
                })
            })
        })
    })
})
