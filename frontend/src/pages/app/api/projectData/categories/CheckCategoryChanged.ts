import {BibCategory} from "@/pages/app/api/projectData/categories/BibCategory";

export default function CheckCategoryChanged(initial: BibCategory, name: string, citaviCategory: string, citaviFilter: string[]) {
  return name !== initial.Name ||
    citaviCategory !== initial.CitaviType ||
    citaviFilter.length !== initial.CitaviNecessaryFields.length ||
    !citaviFilter.every((v, i) => v === initial.CitaviNecessaryFields[i]);
}
