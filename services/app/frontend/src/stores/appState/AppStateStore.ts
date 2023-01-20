import {defineStore} from "pinia";

export const pageNames = [
    "main",
    "project",
    "category"
];

export interface AppState {
    history: string[],
    sidebarOpen: boolean,
    currentProject: string,
    navigatingBack: boolean
}

export const useAppStateStore = defineStore({
    id: "app-state",
    state: () => ({
        history: [pageNames[0]],
        sidebarOpen: false,
        navigatingBack: false,
    } as AppState),

    getters: {
        currentPage: (state: AppState): string => {
            return state.history[state.history.length - 1]
        }
    },

    actions: {
        navToPage(name: string) {
            this.history.push(name);
        },
        goBack() {
            this.navigatingBack = true;
        },
        finishGoBack() {
            this.navigatingBack = false;
            this.history.pop();
            if (this.history.length === 1) {
                this.sidebarOpen = false;
                this.currentProject = "";
            }
        },
        setSidebarOpened(v: boolean) {
            this.sidebarOpen = v;
        },
        setProject(name: string) {
            this.currentProject = name;
        }
    }
})