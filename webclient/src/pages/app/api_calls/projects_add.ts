import {ApiResponse} from "@/pages/app/api_calls/response";

export default async function AddProject(name: string): Promise<ApiResponse> {
    const resp = await fetch("/app/createProject", {
        method: "PUT",
        body: name
    })

    return {
        Success: resp.ok
    }
}
