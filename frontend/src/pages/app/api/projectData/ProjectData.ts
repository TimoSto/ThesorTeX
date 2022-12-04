import {BibEntry} from "@/pages/app/api/projectData/entries/BibEntry";
import {BibCategory} from "@/pages/app/api/projectData/categories/BibCategory";

export interface ProjectData {
  Entries: BibEntry[]
  Categories: BibCategory[]
}
