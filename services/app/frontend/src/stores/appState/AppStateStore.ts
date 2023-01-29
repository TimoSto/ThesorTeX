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
        sidebarOpen: false,
        navigatingBack: false,
        currentProject: "",
        currentItem: "",
        unsavedChanges: false,
        unsavedDialogTriggered: false,
        unsavedDialogCallback: () => {
        },
    } as AppState),

    getters: {
        currentPage: (state: AppState): string => {
            return state.history[state.history.length - 1];
        }
    },

    actions: {
        navToPage(name: string) {
            this.history.push(name);
        },
        goBack() {
            if (!this.unsavedChanges) {
                this.navigatingBack = true;
            } else {
                this.unsavedDialogTriggered = true;
                this.unsavedDialogCallback = () => {
                    this.navigatingBack = true;
                };
            }
        },
        finishGoBack() {
            this.navigatingBack = false;
            this.history.pop();
            this.currentItem = "";
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
        },
        setItem(id: string) {
            this.currentItem = id;
        },
        switchToProject(name: string) {
            if (!this.unsavedChanges) {
                this.history = this.history.slice(0, 2);
                this.setProject(name);
            } else {
                this.unsavedDialogTriggered = true;
                this.unsavedDialogCallback = () => {
                    this.history = this.history.slice(0, 2);
                    this.setProject(name);
                };
            }
        },
        resolveCallback(accept: boolean) {
            this.unsavedDialogTriggered = false;
            if (accept) {
                this.unsavedDialogCallback();
                this.unsavedChanges = false;
            } else {
                this.unsavedDialogCallback = () => {
                };
            }
        }
    }
});