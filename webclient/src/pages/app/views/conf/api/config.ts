
export interface AppConfiguration {
    Port: string
    ProjectsDir: string
    Error: string
}

export async function GetConfig(): Promise<AppConfiguration> {
    const resp = await fetch("/app/config");

    if( !resp.ok ) {
        return {
            Error: resp.status.toString(),
            Port: "",
            ProjectsDir: ""
        }
    }

    const obj = await resp.json()

    obj.Error = '';

    return obj;
}

export async function SaveConfig(config: AppConfiguration): Promise<boolean> {
    const resp = await fetch("/app/config", {
        method: "POST",
        body: JSON.stringify(config)
    });

    return resp.ok
}
