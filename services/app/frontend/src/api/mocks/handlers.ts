import { rest } from "msw";
import ProjectMetaData from "../../domain/projects/ProjectMetaData";

const data = {
    data: [] as ProjectMetaData[],
    code: 200
}

export function SetResponse(d: ProjectMetaData[], code: number) {
    data.data = d;
    data.code = code;
}

export const handlers = [
    rest.get("http://localhost:2345/getAllProjects", (req, res, ctx) => {

        return res(
            // Respond with a 200 status code
            ctx.status(data.code),
            ctx.json(data.data)
        )
    })
]