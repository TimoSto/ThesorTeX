import {describe, expect, it} from "vitest";
import GetProjectsMetaData from "./GetProjectsMetaData";
import {SetResponse} from "../mocks/handlers";

describe("GetProjectsMetaData", async () => {
    it("should give ok and empty", async () => {
        const resp = await GetProjectsMetaData();

        expect(resp.Ok).toBeTruthy();
        expect(resp.Projects).toEqual([]);
    })
    it("should give ok and two projects", async () => {
        SetResponse([
            {Name: "test", Created: "", LastEdited: "", NumberOfEntries: ""},
            {Name: "test2", Created: "", LastEdited: "", NumberOfEntries: ""},
        ], 200)
        const resp = await GetProjectsMetaData();

        expect(resp.Ok).toBeTruthy();
        expect(resp.Projects.length).toEqual(2);
    })
    it("should give not ok and zero projects", async () => {
        SetResponse([
            {Name: "test", Created: "", LastEdited: "", NumberOfEntries: ""},
            {Name: "test2", Created: "", LastEdited: "", NumberOfEntries: ""},
        ], 500)
        const resp = await GetProjectsMetaData();

        expect(resp.Ok).toBeFalsy();
        expect(resp.Projects.length).toEqual(0);
    })
})
