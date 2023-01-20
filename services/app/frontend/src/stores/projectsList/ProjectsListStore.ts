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
        setProjects(projects: ProjectMetaData[]) {
            this.projects = projects;
        },
        addProject(p: ProjectMetaData) {
            this.projects.push(p);
            this.projects.sort((p1, p2) => {
                return p1.Name.toLowerCase() > p2.Name.toLowerCase() ? 1 : -1;
            })
        },
        removeProject(p: string) {
            const i = this.projects.map(p => p.Name).indexOf(p);
            if (i >= 0) {
                this.projects.splice(i, 1);
            }
        }
    }
})