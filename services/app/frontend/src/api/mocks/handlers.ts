import { rest } from "msw";
import ProjectMetaData from "../../domain/projects/ProjectMetaData";

const getAllProjectsData = {
    data: [] as ProjectMetaData[],
    code: 200
}

const createProjectData = {
    data: {} as ProjectMetaData,
    code: 200
}

export function SetResponse_GetAllProjects(d: ProjectMetaData[], code: number) {
    getAllProjectsData.data = d;
    getAllProjectsData.code = code;
}

export function SetResponse_CreateProject(d: ProjectMetaData, code: number) {
    createProjectData.data = d;
    createProjectData.code = code;
}

export const handlers = [
    rest.get("http://localhost:2345/getAllProjects", (req, res, ctx) => {

        return res(
            // Respond with a 200 status code
            ctx.status(getAllProjectsData.code),
            ctx.json(getAllProjectsData.data)
        )
    }),
    rest.put("http://localhost:2345/createNewProject", (req, res, ctx) => {

        return res(
            // Respond with a 200 status code
            ctx.status(createProjectData.code),
            ctx.json(createProjectData.data)
        )
    })
]