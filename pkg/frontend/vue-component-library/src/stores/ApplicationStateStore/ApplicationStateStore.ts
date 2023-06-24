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
}

const MainPage = "main";

export const useApplicationStateStore = defineStore({
    id: "application-state",
    state: () => ({
        history: [MainPage],
        navState: NavState.None
    } as ApplicationState),
    actions: {
        openPage(page: string) {
            this.$state.history.push(page);
            this.$state.navState = NavState.Forth;
            setTimeout(() => {
                this.$state.navState = NavState.None;
            }, 750);
        },
        goBack(n: number) {
            this.$state.navState = n === 1 ? NavState.Back : NavState.BackMultiple;
            setTimeout(() => {//TODO: unit test
                this.$state.navState = NavState.None;
                this.$state.history = this.$state.history.slice(0, -1);
            }, 750);
        }
    }
});
