import {CitaviEntry} from "./Citavi";
import {Category} from "../category/Category";
import {Entry} from "../entry/Entry";
import {trimAndParseValue} from "./ParseBibValues";

interface Unknown {
    Key: string,
    Category: string
}

export interface AnalyseResult {
    Entries: Entry[],
    Unknown: Unknown[]
}

export function AnalyseBibFile(file: string, categories: Category[]): AnalyseResult {
    const citaviEntries = GetEntries(file);

    let entries: Entry[] = [];
    let unknowns: Unknown[] = [];

    citaviEntries.forEach(e => {
        const entry = AssignCategory(e, categories);
        if (entry) {
            entries.push(entry);
        } else {
            unknowns.push({Key: e.Key, Category: e.Category});
        }
    });

    return {
        Entries: entries,
        Unknown: unknowns
    };
}

export function GetEntries(file: string): CitaviEntry[] {

    const entries: CitaviEntry[] = [];

    // split at @ in newline
    const parts = file.split(/^@/m);

    for (let i = 0; i < parts.length; i++) {
        if (parts[i] === "") {
            continue;
        }
        // rm all newlines to have a consistent pattern
        parts[i] = parts[i].replaceAll("\n", "");

        //string until first { is category
        const category = parts[i].match(/^[^{]*/);
        if (!category) {
            console.warn("cloud not read category");
            continue;
        }
        // between first { and first , is the key
        const key = parts[i].match(/(?<={)[^,]*/);
        if (!key) {
            console.warn("cloud not read key");
            continue;
        }

        const entry: CitaviEntry = {
            Key: key[0],
            Category: category[0],
            Attributes: []
        };

        // rm category, key and surrounding brackets
        parts[i] = parts[i].replace(`${category}{${key},`, "");
        const lastIndex = parts[i].lastIndexOf("}");
        if (lastIndex > 0) {
            parts[i] = parts[i].slice(0, lastIndex);
        }

        //TODO: take enclosure in {} into account

        // split at comma outside quotes
        const attributeValuePairs = parts[i].split(/,(?=(?:[^"]*"[^"]*")*[^"]*$)/g);

        for (let j = 0; j < attributeValuePairs.length; j++) {
            attributeValuePairs[j] = attributeValuePairs[j].trim();

            const attrValue = attributeValuePairs[j].split(/=(.*)/s);
            entry.Attributes.push({
                Attr: attrValue[0].trim(),
                Value: attrValue[1].trim()
            });
        }

        entries.push(entry);
    }

    return entries;
}

export function getCategoryScore(entry: CitaviEntry, category: Category): number {
    let score = 0;

    if (entry.Category !== category.CitaviCategory) {
        return -1;
    }

    const attributes = entry.Attributes.map(a => a.Attr);

    for (let i = 0; i < category.CitaviFilter.length; i++) {
        if (attributes.indexOf(category.CitaviFilter[i]) >= 0) {
            score++;
        } else {
            return -1;
        }
    }

    return score;
}

export function AssignCategory(entry: CitaviEntry, categories: Category[]): Entry | undefined {
    let max = -1;
    let category: Category = {} as Category;

    categories.forEach(c => {
        const score = getCategoryScore(entry, c);
        if (score >= 0 && score > max) {
            category = c;
            max = score;
        }
    });

    console.log(max);

    if (max === -1) {
        return undefined;
    }

    const parsedEntry: Entry = {
        Key: entry.Key,
        Category: category.Name,
        Fields: []
    };

    category.BibFields.forEach(f => {
        let val = "";

        const attribute = entry.Attributes.find(a => f.CitaviMapping.indexOf(a.Attr) >= 0);

        if (attribute) {
            val = trimAndParseValue(attribute.Value);
        }

        parsedEntry.Fields.push(val);
    });

    return parsedEntry;

}
