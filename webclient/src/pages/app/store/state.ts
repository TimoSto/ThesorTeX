
export type AppState = {
    actionResult: {
        success: string,
        error: string
    },
    projects: ProjectData[]
}

export type ProjectData = {
    Name: string
}

export const appState: AppState = {
    actionResult: {
        success: "",
        error: ""
    },
    projects: []
}