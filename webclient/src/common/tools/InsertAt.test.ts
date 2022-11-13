import InsertAt from "@/common/tools/InsertAt";

describe('InsertAt', () => {
    it('at end', () => {
        expect(InsertAt('hallo', 'i', 5)).toEqual('halloi')
    })
    it('at start', () => {
        expect(InsertAt('hallo', 'i', 0)).toEqual('ihallo')
    })
    it('in middle', () => {
        expect(InsertAt('hallo', 'i', 2)).toEqual('haillo')
    })
    it('on empty', () => {
        expect(InsertAt('', 'i', 0)).toEqual('i')
    })
})
