import {defineStore} from "pinia";
import {ref} from "vue";
import ProjectMetaData from "../api/projects/ProjectMetaData";


export const useProjectsStore = defineStore('projects-store', () => {
  const projects = ref([] as ProjectMetaData[])

  function setProjects(data: ProjectMetaData[]) {
    projects.value = data;
  }

  function addProject(data: ProjectMetaData) {
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

  function updateLastEditedOnProject(project: string, date: string) {
    console.log(date, project, projects.value);
    projects.value.forEach((p,i) => {
      if( p.Name === project ) {
        projects.value[i].LastModified = date;
      }
    })
  }

  return {projects, setProjects, addProject, rmProject, updateLastEditedOnProject}
})
