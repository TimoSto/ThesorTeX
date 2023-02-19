export interface AnalyseLineResult {
    type: "" | "H1" | "H2" | "H3" | "H4" | "H5" | "TEXT" | "EMPTY" | "START_CODE" | "END_CODE" | "START_LIST" | "END_LIST" | "LIST_ITEM" | "IMAGE" | "LINE" | "INFO";
    content?: string;
    codeType?: "TEX" | "JS";
}

interface TypeRegexMapping {
    type: AnalyseLineResult["type"],
    regex: RegExp,
    prefixToRemove: string
}

const typeMappings: TypeRegexMapping[] = [
    {
        type: "H1",
        regex: new RegExp("^# "),
        prefixToRemove: "# "
    },
    {
        type: "H2",
        regex: new RegExp("^## "),
        prefixToRemove: "## "
    },
    {
        type: "H3",
        regex: new RegExp("^### "),
        prefixToRemove: "### "
    },
    {
        type: "H4",
        regex: new RegExp("^#### "),
        prefixToRemove: "#### "
    },
    {
        type: "H5",
        regex: new RegExp("^##### "),
        prefixToRemove: "##### "
    },
    {
        type: "START_CODE",
        regex: new RegExp("^```[^\s]"),
        prefixToRemove: "```"
    },
    {
        type: "LIST_ITEM",
        regex: new RegExp("^- "),
        prefixToRemove: "- "
    },
];

export function analyseLine(line: string): AnalyseLineResult {
    const result: AnalyseLineResult = {
        type: ""
    };
    let type: TypeRegexMapping | undefined = undefined;
    for (const m of typeMappings) {
        if (m.regex.test(line)) {
            type = m;
            break;
        }
    }
    if (!!type) {
        result.type = type.type;
        let content = line.startsWith(type.prefixToRemove) ? line.slice(type.prefixToRemove.length) : line;
        result.content = content;
    } else {
        if (line.length > 0) {
            if (line.trim() === "---") {
                result.type = "LINE";
            } else if (line.trim() === "```") {
                result.type = "END_CODE";
            } else {
                result.type = "TEXT";
                result.content = line;
            }
        } else {
            result.type = "EMPTY";
        }
    }

    return result;
}
