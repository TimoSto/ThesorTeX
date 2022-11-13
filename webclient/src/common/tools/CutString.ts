
export default function CutString(str: string, index: number, length: number): string {
    return str.substring(0, index) + str.substring(index + length);
}
