import {defineStore} from "pinia";

export const pageNames = [
    "main",
    "project"
];

export interface AppState {
    history: string[],
    sidebarOpen: boolean,
}

export const useAppStateStore = defineStore( {
    id: 'app-state',
    state: () => ({
        history: [pageNames[0]],
        sidebarOpen: false,
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
          this.history.pop();
          if( this.history.length === 1 ) {
              this.sidebarOpen = false;
          }
        },
        setSidebarOpened(v: boolean) {
            this.sidebarOpen = v;
        }
    }
})