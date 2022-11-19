import { ActionTree, ActionContext } from 'vuex';
import {AppState} from './state';
import { Mutations } from './mutations';

export type AugmentedActionContext = {
    commit<K extends keyof Mutations>(
        key: K,
        payload: Parameters<Mutations[K]>[1]
    ): ReturnType<Mutations[K]>;
} & Omit<ActionContext<AppState, AppState>, 'commit'>;

export interface Actions {

}

export const actions: ActionTree<AppState, AppState> & Actions = {

};