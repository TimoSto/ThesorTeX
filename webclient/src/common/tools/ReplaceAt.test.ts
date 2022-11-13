import ReplaceAt from "@/common/tools/ReplaceAt";

describe('RepacleAt', () => {
    it('at end', () => {
        expect(ReplaceAt('hallo', 'i', 4)).toEqual('halli')
    })
    it('at start', () => {
        expect(ReplaceAt('hallo', 'i', 0)).toEqual('iallo')
    })
    it('in middle', () => {
        expect(ReplaceAt('hallo', 'i', 2)).toEqual('hailo')
    })
})
