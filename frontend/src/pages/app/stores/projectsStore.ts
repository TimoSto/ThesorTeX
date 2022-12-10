import {defineStore} from "pinia";
import {ref} from "vue";
import ProjectOverviewData from "../api/projects/ProjectOverviewData";


export const useProjectsStore = defineStore('projects-store', () => {
  const projects = ref([] as ProjectOverviewData[])

  function setProjects(data: ProjectOverviewData[]) {
    projects.value = data;
  }

  function addProject(data: ProjectOverviewData) {
    projects.value.push(data);

    projects.value.sort((p1, p2) => {
      return p1.Name.toUpperCase() > p2.Name.toUpperCase() ? 1 : -1;
    })
  }

  function rmProject(name: string) {
    const i = projects.value.map(p => p.Name).indexOf(name);
    if( i >= 0) {
      projects.value.splice(i, 1);
    }
  }

  return {projects, setProjects, addProject, rmProject}
})
