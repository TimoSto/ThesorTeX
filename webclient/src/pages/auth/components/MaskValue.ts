/**
 * Function to update the value of a masked input field
 * @param value value before input and unmasked
 * @param mask value before input but masked
 * @param inputValue value after input, initial value is masked, now entered is unmasked
 */
export default function MaskValue(value: string, mask: string, inputValue: string) {

}

interface StringDifference {
    Removed: number
    Changed: CharChanged[]
    Added: string[]
}

interface CharChanged {
    Index: number
    Old: string
    New: string
}

function GetStringDifference(oldStr: string, newStr: string): StringDifference {
    let diff = {
        Removed: 0,
        Changed: [],
        Added: []
    } as StringDifference

    return diff
}
