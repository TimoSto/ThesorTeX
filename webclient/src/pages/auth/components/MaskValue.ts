/**
 * Function to update the value of a masked input field
 * @param CurrentValue value before input and unmasked
 * @param CurrentMask value before input but masked
 * @param InputState value inside input, initial value is masked, now entered is unmasked
 */

export interface MaskState {
    CurrentValue: string
    CurrentMask: string
    InputState: string
}

const maskChar = '#';

export default function MaskValue(state: MaskState): MaskState {
    const newMask: MaskState = {
        CurrentValue: state.CurrentValue,
        CurrentMask: state.CurrentMask,
        InputState: '',
    }
    const diff_value = GetStringDifference(state.CurrentValue, state.InputState);

    for( let i = 0; i < diff_value.Added.length ; i++ ) {
        newMask.CurrentValue += diff_value.Added[i];
        newMask.CurrentMask += maskChar;
    }

    newMask.CurrentValue = newMask.CurrentValue.substring(0, newMask.CurrentValue.length - diff_value.Removed);
    newMask.CurrentMask = newMask.CurrentMask.substring(0, newMask.CurrentMask.length - diff_value.Removed);

    return newMask;
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

export function GetStringDifference(oldStr: string, newStr: string): StringDifference {
    let diff = {
        Removed: 0,
        Changed: [],
        Added: []
    } as StringDifference

    const lengthDiff = oldStr.length - newStr.length;
    if( lengthDiff > 0 ) {
        diff.Removed = lengthDiff;
    }

    for( let i = 0 ; i < newStr.length ; i++ ) {
        if( i < oldStr.length ) {
            if( oldStr.charAt(i) !== newStr.charAt(i) ) {
                diff.Changed.push({
                    Index: i,
                    Old: oldStr.charAt(i),
                    New: newStr.charAt(i)
                })
            }
        } else {
            diff.Added.push(newStr.charAt(i))
        }
    }

    return diff
}
