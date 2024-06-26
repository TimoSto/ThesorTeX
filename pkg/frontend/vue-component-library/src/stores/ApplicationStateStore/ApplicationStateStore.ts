import {defineStore} from "pinia";

export enum NavState {
    None,
    Forth,
    Back = -1,
    BackMultiple = -2
}

export interface ApplicationState {
    history: string[],
    navState: NavState,
    instantNav: boolean,
    unsavedChanges: boolean,
    unsavedDialogTriggered: boolean,
    unsavedDialogCallback: () => void,
    sidebarOpened: boolean
}

const MainPage = "main";

export const useApplicationStateStore = defineStore({
    id: "application-state",
    state: () => ({
        history: [MainPage],
        navState: NavState.None,
        instantNav: false,
        unsavedChanges: false,
        unsavedDialogTriggered: false,
        unsavedDialogCallback: () => {
        },
        sidebarOpened: false,
    } as ApplicationState),
    getters: {
        currentPage: (state: ApplicationState): string => {
            return state.history[state.history.length - 1];
        }
    },
    actions: {
        openPage(page: string) {
            this.$state.history.push(page);
            this.$state.navState = NavState.Forth;
            setTimeout(() => {
                this.$state.navState = NavState.None;
            }, 750);
        },
        goBack(n: number, instantNav?: boolean, cb?: () => void) {
            this.$state.instantNav = !!instantNav;
            if (this.unsavedChanges) {
                this.unsavedDialogTriggered = true;
                this.unsavedDialogCallback = () => {
                    this.unsavedChanges = false;
                    this.goBack(n, instantNav, cb);
                };
            } else {
                this.$state.navState = n === 1 ? NavState.Back : NavState.BackMultiple;
                setTimeout(() => {//TODO: unit test
                    this.$state.navState = NavState.None;
                    this.$state.history = this.$state.history.slice(0, -1);
                    this.$state.instantNav = false;
                }, instantNav ? 0 : 750);
                if (!!cb) {
                    cb();
                }
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
