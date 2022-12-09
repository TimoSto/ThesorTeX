import {BibEntry} from "@/pages/app/api/projectData/entries/BibEntry";

export interface EntrySaveData extends BibEntry {
  InitialKey: string
  Project: string
}

export default async function SaveEntry(entry: BibEntry, initialKey: string, project: string): Promise<boolean> {
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

  return resp.ok;
}
