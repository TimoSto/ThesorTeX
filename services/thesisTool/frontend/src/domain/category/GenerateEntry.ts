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

export function GenerateCiteForCategory(bibFields: Field[], citeFields: Field[], values: string[]): string {
    const filteredValues: string[] = [];

    const bibFieldNames = bibFields.map(f => f.Name);

    const filteredCiteFieldNames = citeFields.map(f => f.Name).filter(f => bibFieldNames.indexOf(f) === -1);

    citeFields.forEach((f, i) => {
        const indexInBib = bibFieldNames.indexOf(f.Name);
        if (indexInBib === -1) {
            const value = values[bibFieldNames.length + filteredCiteFieldNames.indexOf(f.Name)];
            filteredValues.push(value ? value : "");
        } else {
            const value = values[indexInBib];
            filteredValues.push(value ? value : "");
        }
    });

    return GenerateEntryForCategory(citeFields, filteredValues);
}
