
export type AppState = {
    projects: ProjectData[]
}

export type ProjectData = {
    Name: string
}

export const appState: AppState = {
    projects: []
}