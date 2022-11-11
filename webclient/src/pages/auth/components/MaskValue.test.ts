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
                    InputState: 'a'
                });

                expect(res).toEqual({//TODO: unicode &bull; für mask verwenden
                    CurrentValue: 'a',
                    CurrentMask: '•',
                    InputState: ''
                })
            })
            it('second input', () => {
                const res = MaskValue({
                    CurrentValue: 'a',
                    CurrentMask: '•',
                    InputState: '•b'
                });

                expect(res).toEqual({//TODO: unicode &bull; für mask verwenden
                    CurrentValue: 'ab',
                    CurrentMask: '••',
                    InputState: ''
                })
            })
            it('third input', () => {
                const res = MaskValue({
                    CurrentValue: 'ab',
                    CurrentMask: '••',
                    InputState: '••c'
                });

                expect(res).toEqual({//TODO: unicode &bull; für mask verwenden
                    CurrentValue: 'abc',
                    CurrentMask: '•••',
                    InputState: ''
                })
            })
            it('multiple char input', () => {
                const res = MaskValue({
                    CurrentValue: 'ab',
                    CurrentMask: '••',
                    InputState: '••cdefg'
                });

                expect(res).toEqual({//TODO: unicode &bull; für mask verwenden
                    CurrentValue: 'abcdefg',
                    CurrentMask: '•••••••',
                    InputState: ''
                })
            })
        })
        describe('remove from end', () => {
            it('rm second char', () => {
                const res = MaskValue({
                    CurrentValue: 'ab',
                    CurrentMask: '••',
                    InputState: '•'
                });

                expect(res).toEqual({
                    CurrentValue: 'a',
                    CurrentMask: '•',
                    InputState: ''
                })
            })
            it('rm first char', () => {
                const res = MaskValue({
                    CurrentValue: 'a',
                    CurrentMask: '•',
                    InputState: ''
                });

                expect(res).toEqual({
                    CurrentValue: '',
                    CurrentMask: '',
                    InputState: ''
                })
            })
            it('rm multiple chars', () => {
                const res = MaskValue({
                    CurrentValue: 'abcd',
                    CurrentMask: '••••',
                    InputState: '•'
                });

                expect(res).toEqual({
                    CurrentValue: 'a',
                    CurrentMask: '•',
                    InputState: ''
                })
            })
        })
    })
})
