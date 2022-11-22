import { MutationTree } from 'vuex';
import { MutationTypes } from './mutation-types';
import {AppState, ProjectData} from "@/pages/app/store/state";

export type Mutations<S = AppState> = {
    [MutationTypes.SET_SUCCESS](state: AppState, message: string): void
    [MutationTypes.SET_ERROR](state: AppState, message: string): void
    [MutationTypes.SET_PROJECTS](state: AppState, projects: ProjectData[]): void
    [MutationTypes.ADD_PROJECT](state: AppState, project: ProjectData): void
};

export const mutations: MutationTree<AppState> & Mutations = {
    [MutationTypes.SET_SUCCESS](state: AppState, message: string) {
        state.actionResult.success = message;
    },
    [MutationTypes.SET_ERROR](state: AppState, message: string) {
        state.actionResult.error = message;
    },
    [MutationTypes.SET_PROJECTS](state, payload: ProjectData[]) {
        if( !payload ) {
            payload = [];
        }
        state.projects = payload
    },
    [MutationTypes.ADD_PROJECT](state: AppState, project: ProjectData) {
        state.projects.push(project);
        state.projects.sort((a: ProjectData, b: ProjectData) => {
            return a.Name.toUpperCase() > b.Name.toUpperCase() ? 1 : -1;
        });
    }
};