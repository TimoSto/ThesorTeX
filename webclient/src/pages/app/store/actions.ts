import { ActionTree, ActionContext } from 'vuex';
import {AppState} from './state';
import { Mutations } from './mutations';
import {ActionTypes} from "@/pages/app/store/action-types";
import {MutationTypes} from "@/pages/app/store/mutation-types";
import GetProjects from "@/pages/app/api_calls/projects_get";
import AddProject from "@/pages/app/api_calls/projects_add";

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
        const resp = await GetProjects();

        if( resp.Success ) {
            commit(MutationTypes.SET_PROJECTS, resp.Projects)
        }
    },
    async [ActionTypes.ADD_PROJECT]({ commit }, name: string) {
        const resp = await AddProject(name);

        if( resp.Success ) {
            commit(MutationTypes.ADD_PROJECT, resp.Project);
        }
    },
};