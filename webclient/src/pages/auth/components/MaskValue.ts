/**
 * Function to update the value of a masked input field
 * @param CurrentValue value before input and unmasked
 * @param CurrentMask value before input but masked
 * @param InputState value inside input, initial value is masked, now entered is unmasked
 */
import InsertAt from "@/common/tools/InsertAt";
import CutString from "@/common/tools/CutString";

export interface MaskState {
    CurrentValue: string
    CurrentMask: string
    InputState: string
    CaretPosition: number
}

export const maskChar = '•';

export default function MaskValue(state: MaskState): MaskState {
    const newMask: MaskState = {
        CurrentValue: state.CurrentValue,
        CurrentMask: state.CurrentMask,
        InputState: '',
        CaretPosition: state.CaretPosition
    }
    const diff_value = GetStringDifference(state.CurrentValue, state.CurrentMask, state.InputState, state.CaretPosition);

    newMask.CurrentValue = InsertAt(state.CurrentValue, diff_value.Added, state.CaretPosition - diff_value.Added.length);
    newMask.CurrentMask += maskChar.repeat(diff_value.Added.length);

    newMask.CurrentValue = CutString(newMask.CurrentValue, state.CaretPosition, diff_value.Removed);
    newMask.CurrentMask = newMask.CurrentMask.substring(0, newMask.CurrentMask.length - diff_value.Removed);

    return newMask;
}

interface StringDifference {
    Removed: number
    Added: string
}

//TODO: mask-param deletable?
export function GetStringDifference(currentValue: string, currentMask: string, currentInput: string, caret: number): StringDifference {
    const diff = {
        Removed: 0,
        Added: ''
    } as StringDifference

    const lengthDiff = currentValue.length - currentInput.length;
    if( lengthDiff > 0 ) {
        diff.Removed = lengthDiff;
    } else {
        diff.Added = currentInput.substring(caret, caret + lengthDiff)
    }

    return diff
}
