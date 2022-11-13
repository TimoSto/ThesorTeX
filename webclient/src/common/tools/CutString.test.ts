import CutString from "@/common/tools/CutString";

describe('CutString', () => {
    it('one from end', () => {
        expect(CutString('hallo', 4, 1)).toEqual('hall')
    })
    it('two from between', () => {
        expect(CutString('hallo', 1, 2)).toEqual('hlo')
    })
    it('one from start', () => {
        expect(CutString('hallo', 0, 1)).toEqual('allo')
    })
})
