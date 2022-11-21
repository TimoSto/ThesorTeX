import { MutationTree } from 'vuex';
import { MutationTypes } from './mutation-types';
import {AppState, ProjectData} from "@/pages/app/store/state";

export type Mutations<S = AppState> = {
    [MutationTypes.SET_PROJECTS](state: AppState, projects: ProjectData[]): void
};

export const mutations: MutationTree<AppState> & Mutations = {
    [MutationTypes.SET_PROJECTS](state, payload: ProjectData[]) {
        console.log(payload);
        state.projects = payload
    },
};