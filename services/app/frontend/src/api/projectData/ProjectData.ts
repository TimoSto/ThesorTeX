import {BibEntry} from "./entries/BibEntry";
import {BibCategory} from "./categories/BibCategory";


export interface ProjectData {
  Entries: BibEntry[]
  Categories: BibCategory[]
}
