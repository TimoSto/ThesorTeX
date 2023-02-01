import {Entry} from "./Entry";
import {Field} from "../category/Category";

type Pair = {
    String: string,
    TeX: string
}

const pairs: Pair[] = [
    {
        String: "_",
        TeX: "{{\\_}}"
    },
    {
        String: ";",
        TeX: "{{;}}",
    },
    {//TODO: ggf vorher amp ersetzen
        String: "&",
        TeX: "{{\\&}}",
    },
    {
        String: "$",
        TeX: "{{\\$}}",
    },
    {
        String: "#",
        TeX: "{{\\#}}",
    },
    {
        String: "%",
        TeX: "{{\\%}}",
    },
    {
        String: "á",
        TeX: "{{\\'{a}}}",
    },
    {
        String: "â",
        TeX: "{{\\u{a}}}",
    },
    {
        String: "å",
        TeX: "{{\\r{a}}}",
    },
    {
        String: "é",
        TeX: "{{\\'{e}}}",
    },
    {
        String: "è",
        TeX: "{{\\`{e}}}",
    },
    {
        String: "<",
        TeX: "{{\\textless}}",
    },
    {
        String: ">",
        TeX: "{{\\textgreater}}",
    },
    {
        String: "°",
        TeX: "{{\\degree}}",
    },
    {
        String: "š",
        TeX: "{{\\v{s}}}",
    },
    {
        String: "č",
        TeX: "{{\\v{c}}}",
    },
];

export function EscapeEntryFields(entry: Entry, fields: Field[]): string[] {
    let escapedValues: string[] = [];

    entry.Fields.forEach((f, i) => {
        if (i < fields.length && !fields[i].Format.Preformatted) {
            pairs.forEach(p => {
                if (p.TeX === `{{\\${p.String}}}`) {
                    //search for all values that should be surrounded by {{\\}}
                    const matches = [...f.matchAll(new RegExp(p.String === "$" ? /\$/ : p.String, "g"))];
                    let added = 0;
                    matches.forEach(m => {
                        if (m.index !== undefined) {
                            const i = m.index + added;
                            //check if they are not surrounded
                            if (i < 3 || (f.substring(i - 3, i) !== "{{\\" && f.substring(i + p.String.length, i + p.String.length + 2) !== "}}")) {
                                f = f.substring(0, i) + p.TeX + f.substring(i + p.String.length);
                                added += p.TeX.length - p.String.length;
                            }
                        }
                    });
                } else {
                    f = f.replaceAll(p.String, p.TeX);
                }
            });
            escapedValues.push(f);
        } else {
            escapedValues.push(f);
        }
    });

    return escapedValues;
}

export function UnEscapeEntryFields(entry: Entry, fields: Field[]): string[] {
    let unescapedValues: string[] = [];
    entry.Fields.forEach((f, i) => {
        if (!fields[i].Format.Preformatted) {
            pairs.forEach(p => {
                unescapedValues.push(f.replaceAll(p.TeX, p.String));
            });
        } else {
            unescapedValues.push(f);
        }
    });

    return unescapedValues;
}
