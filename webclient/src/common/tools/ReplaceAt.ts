
export default function ReplaceAt(str: string, replacement: string, index: number): string {
    return str.substring(0, index) + replacement + str.substring(index + replacement.length);
}
