import {defineStore} from "pinia";

export const pageNames = [
    "main",
    "project",
    "category",
    "entry"
];

export interface AppState {
    history: string[],
    sidebarOpen: boolean,
    currentProject: string,
    navigatingBack: boolean,
    currentItem: string,//entry key or category name
    unsavedChanges: boolean,
    unsavedDialogTriggered: boolean,
    unsavedDialogCallback: () => void
}

export const useAppStateStore = defineStore({
    id: "app-state",
    state: () => ({
        history: [pageNames[0]],
        currentProject: "",
        currentItem: "",
    } as AppState),

    actions: {
        setProject(name: string) {
            this.currentProject = name;
        },
        setItem(id: string) {
            this.currentItem = id;
        },
        switchToProject(name: string) {
            this.history = this.history.slice(0, 2);
            this.setProject(name);
        },
    }
});