import {BibEntry} from "../api/projectData/entries/BibEntry";
import GeneralizeTeXEscaping from "./BibParser";
import {BibCategory, Field} from "../api/projectData/categories/BibCategory";

export interface AttributeValue {
    Attribute: string
    Value: string
}

export interface BibAnalyticsResult {
    Entries: BibEntry[]
    Unknowns: BibEntry[]
}

export default function AnalyseBibFile(file: string, categories: BibCategory[]): BibAnalyticsResult {
    //step 1: separate entries
    //TODO: find a better way with regex
    const entries_raw = separateEntries(file);

    const entries: BibEntry[] = [];
    const unknown: BibEntry[] = [];

    entries_raw.forEach((e: string) => {
        //step 2: get citavi-type of entry
        const cType = getTypeOfEntry(e);

        //step 3: get key
        let key = extractKey(e, entries)

        //step 4: extract attributes and values
        const attributePairs = extractEntryAttributes(e)

        //step 5: find associated bibliography-type
        const bType = findBibliographyCategory(cType, attributePairs, categories);

        if( bType.Name && bType.Name !== '' ) {
            //step 6: create entry of type with values
            entries.push(CreateEntry(bType, attributePairs, key))
        } else {
            unknown.push({Key: key, Category: cType} as BibEntry)
        }
    })

    return {
        Entries: entries,
        Unknowns: unknown
    }
}

export function separateEntries(file: string): string[] {
    file = file.substring(file.indexOf('@'))

    file = file.replace(/\t/g, '');

    return file.split(/^@/gm).filter(e => e.length > 0);
}

export function getTypeOfEntry(e: string): string {
    if( e.charAt(0) === '@' ) {
        e = e.substring(1)
    }
    return e.substring(0, e.indexOf('{'))
}

export function extractKey(e: string, existing: BibEntry[]): string {
    let k = e.substring(e.indexOf('{') + 1, e.indexOf(','));
    while(k.charAt(0) === '\t' ) {
        k = k.substring(1);
    }

    const existingKeys = existing.map(e => e.Key);
    let addition = 0;
    let newKey = k;
    while(existingKeys.indexOf(newKey) !== -1 ) {
        addition++;
        newKey = k + `-${addition}`
    }
    k = newKey;

    return k
}

export function extractEntryAttributes(e: string): AttributeValue[] {
    //only have arguments
    e = e.substring(e.indexOf('{') + 1, e.lastIndexOf('}'));

    //rm all newlines and indents to be equal in all export formats (springer vs ieee vs ...)
    e = e.replaceAll('\n', '');
    e = e.replaceAll('\t', '');

    //split at comma not in quotes
    const re = /,(?=(?:(?:[^"]*"){2})*[^"]*$)/
    const lines = e.split(re);
    //ignore first because first is key
    lines.shift()

    const attributes: AttributeValue[] = []

    lines.forEach(l => {
        const eqIndex = l.indexOf('=');
        if( eqIndex > 0 ) {
            attributes.push({
                Attribute: l.substring(0, eqIndex),
                Value: GeneralizeTeXEscaping(l.substring(eqIndex + 1))
            })
        }
    })

    return attributes
}

export function findBibliographyCategory(citaviType: string, attributes: AttributeValue[], types: BibCategory[]): BibCategory {

    let typeToReturn: BibCategory = {} as BibCategory;
    let maxFoundFilters = 0;

    types.forEach(t => {
        if( t.CitaviCategory.toLowerCase() === citaviType.toLowerCase() ) {
            let allFound = true;

            const attrNames = attributes.map(a => a.Attribute);
            t.CitaviFilters.forEach(f => {
                if( attrNames.indexOf(f) === -1 ) {
                    allFound = false;
                }
            })

            //choose the type with the most found filters
            if( allFound && t.CitaviFilters.length >= maxFoundFilters ) {
                typeToReturn = t;
                maxFoundFilters = t.CitaviFilters.length;
            }
        }
    })

    return typeToReturn;
}

export function CreateEntry(type: BibCategory, attributes: AttributeValue[], key: string): BibEntry {
    const entry: BibEntry = {
        Key: key,
        Category: type.Name,
        CiteNumber: 0,
        Fields: [],
        Example: '',
        Comment: ''
    }

    //fill fields with attribute names
    type.Fields.forEach((f: Field) => {
        //setting attribute name as default
        entry.Fields.push(f.Field)
        //looking for citavi-attribute matching the configured citavi-attributes of field
        attributes.forEach(a => {
            if( f.CitaviAttributes.indexOf(a.Attribute) >= 0 ) {
                const i = entry.Fields.indexOf(f.Field);
                if( i >= 0 ) {
                    entry.Fields[i] = a.Value;
                }
            }
        })
    });

    const bibFields = type.Fields.map((f: Field) => f.Field);

    type.CiteFields.forEach((f: Field) => {
        if( bibFields.indexOf(f.Field) === -1 ) {
            //setting attribute name as default
            entry.Fields.push(f.Field)
            //looking for citavi-attribute matching the configured citavi-attributes of field d
            attributes.forEach(a => {
                if( f.CitaviAttributes.indexOf(a.Attribute) >= 0 ) {
                    entry.Fields[entry.Fields.indexOf(f.Field)] = a.Value;
                }
            })
        }
    });

    for( let i = 0 ; i < attributes.length ; i++ ) {
        const index = entry.Fields.indexOf(attributes[i].Attribute)
        if (index > -1 ) {
            entry.Fields[index] = attributes[i].Value;
        }
    }

    return entry
}
