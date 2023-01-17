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
            // todo externalise
            const resp = await GetProjectsMetaData();
            if( resp.Ok ) {
                this.projects = resp.Projects;
            } else {
                // todo handle error
            }
        },
        addProject(p: ProjectMetaData) {
            this.projects.push(p);
            this.projects.sort((p1, p2) => {
                return p1.Name.toLowerCase() > p2.Name.toLowerCase() ? 1 : -1;
            })
        }
    }
})