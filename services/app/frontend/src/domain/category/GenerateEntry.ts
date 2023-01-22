import {Field} from "./Category";

export default function GenerateEntryForCategory(fields: Field[], values: string[]): string {
    let entry = "";

    if (!fields) {
        return entry;
    }

    for (let i = 0; i < fields.length; i++) {
        entry += fields[i].Format.Prefix;

        let value = i < values.length ? values[i] : "";
        if (fields[i].Format.Italic) {
            value = `<i>${value}</i>`;
        }
        entry += value;

        entry += fields[i].Format.Suffix;
    }

    return entry;
}
