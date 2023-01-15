import {defineStore} from "pinia";
import ProjectMetaData from "../../domain/projects/ProjectMetaData";

export interface ProjectsListState {
    projects: ProjectMetaData[]
}

export const useProjectsListStore = defineStore({
    id: "projects-store",
    state: () => ({
        projects: []
    } as ProjectsListState),

    actions: {
        getList() {

        }
    }
})