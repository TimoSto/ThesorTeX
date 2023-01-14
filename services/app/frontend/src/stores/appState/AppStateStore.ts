import {defineStore} from "pinia";

export const pageNames = [
    "main"
];

export interface AppState {
    currentPage: string,
    sidebarOpen: boolean,
}

export const useAppStateStore = defineStore( {
    id: 'app-state',
    state: () => ({
        currentPage: pageNames[0],
        sidebarOpen: false,
    } as AppState),

    actions: {
        setPage(name: string) {
            this.currentPage = name;
            if( name === pageNames[0] ) {
                this.sidebarOpen = false;
            }
        },
        setSidebarOpened(v: boolean) {
            this.sidebarOpen = v;
        }
    }
})