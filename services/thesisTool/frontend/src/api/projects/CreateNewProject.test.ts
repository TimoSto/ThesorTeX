import {beforeEach, describe, expect, it} from "vitest";
import {SetResponse_CreateProject} from "../mocks/handlers";
import CreateNewProject from "./CreateNewProject";
import ProjectMetaData from "../../domain/projects/ProjectMetaData";
import {host} from "../config";

describe("CreateNewProject", () => {
    beforeEach(() => {
        location.replace(host);
    });

    it("should give 200 and meta data on success", async () => {
        SetResponse_CreateProject({
            Name: "test",
            Created: "2022-02-01",
            LastEdited: "2022-02-01",
            NumberOfEntries: "1"
        }, 200);

        const resp = await CreateNewProject("test");

        expect(resp.Success).toBe(true);
        expect(resp.Data).toEqual({
            Name: "test",
            Created: "2022-02-01",
            LastEdited: "2022-02-01",
            NumberOfEntries: "1"
        });
    });
    it("should give 500 and nothing on error", async () => {
        SetResponse_CreateProject({} as ProjectMetaData, 500);

        const resp = await CreateNewProject("test");

        expect(resp.Success).toBe(false);
        expect(resp.Data).toBe(undefined);
    });
});
