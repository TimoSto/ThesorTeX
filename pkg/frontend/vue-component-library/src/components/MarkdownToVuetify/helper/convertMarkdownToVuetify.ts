import {analyseLine} from "./analyseLine";

export function ConvertMarkdownToVuetify(file: string): VuetifyComponent[] {
    const components: VuetifyComponent[] = [];

    const lines = file.split("\n");

    let insideBlock: boolean | "CODE" | "INFO" | "LIST" | "PLAIN_TEXT" = false;

    let newLine = false;

    lines.forEach(l => {
        const result = analyseLine(l);

        switch (result.type) {
            case "TEXT":
                if (insideBlock === "CODE") {
                    (components[components.length - 1].Content as string[]).push(result.content as string + "\n");
                } else {
                    if (components.length > 0 && components[components.length - 1].Tag === "TEXT" && !newLine) {
                        components[components.length - 1].Content += parseItalicAndBold(result.content as string);
                    } else {
                        components.push({Tag: "TEXT", Content: parseItalicAndBold(result.content as string)});
                        newLine = false;
                    }
                }
                break;
            case "EMPTY":
                newLine = true;
                break;
            case "H1":
                components.push({Tag: "H1", Content: result.content as string});
                break;
            case "H2":
                components.push({Tag: "H2", Content: result.content as string});
                break;
            case "H3":
                components.push({Tag: "H3", Content: result.content as string});
                break;
            case "LIST_ITEM":
                if (components[components.length - 1].Tag !== "LIST") {
                    components.push({Tag: "LIST", Content: []});
                }
                (components[components.length - 1].Content as string[]).push(parseItalicAndBold(result.content as string));
                break;
            case "LINE":
                components.push({Tag: "HR", Content: ""});
                break;
            case "START_CODE":
                components.push({Tag: "CODE", Content: []});
                insideBlock = "CODE";
                break;
            case "END_CODE":
                insideBlock = false;
                break;
        }
    });

    return components;
}

export interface VuetifyComponent {
    Tag: string;
    Content: String | string[];
}

const boldRegex = new RegExp(/\*\*(.*?)\*\*/gs);
const italicRegex = new RegExp(/\*(.*?)\*/gs);

function parseItalicAndBold(line: string): string {
    const boldResult = line!.match(boldRegex);
    if (boldResult) {
        boldResult.forEach(v => {
            line = line?.replace(v, `<b>${v.substring(2, v.length - 2)}</b>`);
        });
    }
    const italicResult = line!.match(italicRegex);
    if (italicResult) {
        italicResult.forEach(v => {
            line = line?.replace(v, `<i>${v.substring(1, v.length - 1)}</i>`);
        });
    }
    return line;
}
