import MaskValue, {GetStringDifference} from "@/pages/auth/components/MaskValue";

describe("Handling masked input", () => {
    describe("GetDiff", () => {
        it('should work on adding to empty', () => {
            expect(GetStringDifference('', '','a', 1)).toEqual({
                Removed: 0,
                Added: 'a'
            });
        })
        it('should work on adding to existing', () => {
            expect(GetStringDifference('a', '•', '•b', 2)).toEqual({
                Removed: 0,
                Added: 'b'
            });
        })
        it('should work on removing last', () => {
            expect(GetStringDifference('ab', '••', '•', 1)).toEqual({
                Removed: 1,
                Added: ''
            });
        })
        it('should work on removing first', () => {
            expect(GetStringDifference('abc', '•••', '••', 0)).toEqual({
                Removed: 1,
                Added: ''
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
        describe('adding in between', () => {
            it('adding at start', () => {
                const res = MaskValue({
                    CurrentValue: 'ab',
                    CurrentMask: '••',
                    InputState: 'c••',
                    CaretPosition: 1
                });

                expect(res).toEqual({
                    CurrentValue: 'cab',
                    CurrentMask: '•••',
                    InputState: '',
                    CaretPosition: 1,
                });
            })
            it('adding in between', () => {
                const res = MaskValue({
                    CurrentValue: 'ab',
                    CurrentMask: '••',
                    InputState: '•c•',
                    CaretPosition: 2
                });

                expect(res).toEqual({
                    CurrentValue: 'acb',
                    CurrentMask: '•••',
                    InputState: '',
                    CaretPosition: 2,
                });
            })
            it('adding multiple in between', () => {
                const res = MaskValue({
                    CurrentValue: 'ab',
                    CurrentMask: '••',
                    InputState: '•cde•',
                    CaretPosition: 4
                });

                expect(res).toEqual({
                    CurrentValue: 'acdeb',
                    CurrentMask: '•••••',
                    InputState: '',
                    CaretPosition: 4,
                });
            })
        })
        describe('removing in between', () => {
            it('remove at start', () => {
                const res = MaskValue({
                    CurrentValue: 'ab',
                    CurrentMask: '••',
                    InputState: '•',
                    CaretPosition: 0
                });

                expect(res).toEqual({
                    CurrentValue: 'b',
                    CurrentMask: '•',
                    InputState: '',
                    CaretPosition: 0,
                });
            })
            it('remove in between', () => {
                const res = MaskValue({
                    CurrentValue: 'abc',
                    CurrentMask: '•••',
                    InputState: '••',
                    CaretPosition: 1
                });

                expect(res).toEqual({
                    CurrentValue: 'ac',
                    CurrentMask: '••',
                    InputState: '',
                    CaretPosition: 1,
                });
            })
            it('removing multiple in between', () => {
                const res = MaskValue({
                    CurrentValue: 'abcd',
                    CurrentMask: '••••',
                    InputState: '••',
                    CaretPosition: 1
                });

                expect(res).toEqual({
                    CurrentValue: 'ad',
                    CurrentMask: '••',
                    InputState: '',
                    CaretPosition: 1,
                });
            })
        })
    })
})
