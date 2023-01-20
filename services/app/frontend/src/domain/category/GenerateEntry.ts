import {Category} from "./Category";

export default function GenerateEntryForCategory(category: Category | undefined, fields: string[]): string {
    let entry = "";

    if (!category) {
        return entry;
    }

    for (let i = 0; i < category.BibFields.length; i++) {
        entry += category.BibFields[i].Format.Prefix;

        let value = i < fields.length ? fields[i] : "";
        switch (category.BibFields[i].Format.Style) {
            case "italic":
                value = `<i>${value}</i>`
        }
        entry += value;

        entry += category.BibFields[i].Format.Suffix;
    }

    return entry;
}
