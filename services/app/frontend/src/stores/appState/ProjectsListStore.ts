import {defineStore} from "pinia";
import ProjectMetaData from "../../domain/projects/ProjectMetaData";
import GetProjectsMetaData from "../../api/projects/GetProjectsMetaData";

export interface ProjectsListState {
    projects: ProjectMetaData[]
}

export const useProjectsListStore = defineStore({
    id: "projects-store",
    state: () => ({
        projects: []
    } as ProjectsListState),

    actions: {
        async readAllProjectsFromServer() {
            const resp = await GetProjectsMetaData();
            if( resp.Ok ) {
                this.projects = resp.Projects;
            } else {
                // todo handle error
            }
        }
    }
})