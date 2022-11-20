import { ActionTree, ActionContext } from 'vuex';
import {AppState} from './state';
import { Mutations } from './mutations';
import {ActionTypes} from "@/pages/app/store/action-types";

export type AugmentedActionContext = {
    commit<K extends keyof Mutations>(
        key: K,
        payload: Parameters<Mutations[K]>[1]
    ): ReturnType<Mutations[K]>;
} & Omit<ActionContext<AppState, AppState>, 'commit'>;

export interface Actions {
    [ActionTypes.GET_PROJECTS](
        { commit }: AugmentedActionContext,
        payload: undefined
    ): Promise<void>;
}

export const actions: ActionTree<AppState, AppState> & Actions = {
    async [ActionTypes.GET_PROJECTS]({ commit }) {
        //Reading projects from local server
    },
};