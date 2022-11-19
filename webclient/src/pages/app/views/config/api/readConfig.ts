
export interface AppConfiguration {
    Port: string
    ProjectsDir: string
    Error: string
}

export default async function GetConfig(): Promise<AppConfiguration> {
    const resp = await fetch("/app/config");

    if( !resp.ok ) {
        return {
            Error: resp.status.toString(),
            Port: "",
            ProjectsDir: ""
        }
    }

    return await resp.json()
}
