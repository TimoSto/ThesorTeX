
export default function CheckPortValid(port: string): boolean|string {

    if( !/^\d+$/.test(port) ) {
        return "Only numbers"
    }

    return true
}
