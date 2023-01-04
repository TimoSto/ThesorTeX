import {BibEntry} from "./BibEntry";
import SaveResponse from "../SaveResponse";

export interface EntrySaveData extends BibEntry {
  InitialKey: string
  Project: string
}

export default async function SaveEntry(entry: BibEntry, initialKey: string, project: string): Promise<SaveResponse> {
  const obj: EntrySaveData = {
    ...entry,
    InitialKey: initialKey,
    Project: project
  }

  const data = JSON.stringify(obj);

  const resp = await fetch('/app/saveEntry', {
    method: 'PUT',
    body: data
  });

  const respData = await resp.json();

  return {
    Ok: resp.ok,
    Data: respData
  };
}
