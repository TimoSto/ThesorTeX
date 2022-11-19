import { MutationTree } from 'vuex';
import { MutationTypes } from './mutation-types';
import {AppState} from "@/pages/app/store/state";

export type Mutations<S = AppState> = {
    [MutationTypes.SET_PROJECTS](state: AppState, projects: string[]): void
};

export const mutations: MutationTree<AppState> & Mutations = {
    [MutationTypes.SET_PROJECTS](state, payload: string[]) {
        state.projects = payload
    },
};