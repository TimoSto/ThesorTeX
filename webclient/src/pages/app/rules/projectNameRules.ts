
export function CheckProjectName(name: string, existing: string[]): boolean|string {
    if( existing.indexOf(name) > -1 ) {
        return 'NAME_EXISTS'
    }

    return true
}
