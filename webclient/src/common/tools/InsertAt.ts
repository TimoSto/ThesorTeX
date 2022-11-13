
export default function InsertAt(str: string, value: string, index: number): string {
    return str.substring(0, index) + value + str.substring(index)
}
