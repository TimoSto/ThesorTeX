import {analyseLine} from "./analyseLine";

const boldRegex = new RegExp(/\*\*(.*?)\*\*/gs);
const italicRegex = new RegExp(/\*(.*?)\*/gs);

export function convertMarkdownToVuetify(file: string): string {
    let html = "";

    const lines = file.split("\n");

    let insideBlock: boolean | "CODE" | "INFO" | "LIST" | "PLAIN_TEXT" = false;

    lines.forEach(line => {
        const result = analyseLine(line);

        if (insideBlock === "PLAIN_TEXT" && result.type !== "TEXT") {
            html += "</p>\n";
            insideBlock = false;
        }
        if (insideBlock === "LIST" && result.type !== "LIST_ITEM") {
            html += "</ul>\n";
            insideBlock = false;
        }

        if (result.type === "H1") {
            html += `<h1 class="text-h4">${result.content}</h1>` + "\n";
        } else if (result.type === "H2") {
            html += `<h2 class="text-h5">${result.content}</h2>` + "\n";
        } else if (result.type === "H3") {
            html += `<h3 class="text-h5">${result.content}</h3>` + "\n";
        } else if (result.type === "H4") {
            html += `<h4 class="text-h6">${result.content}</h4>` + "\n";
        } else if (result.type === "H5") {
            html += `<h5 class="text-h6">${result.content}</h5>` + "\n";
        } else if (result.type === "TEXT") {
            if (insideBlock !== "PLAIN_TEXT") {
                insideBlock = "PLAIN_TEXT";
                html += "<p class=\"text-body-1\">";
            }
            let contentToAdd = parseItalicAndBold(result.content!);

            if (result.content?.endsWith("  ")) {
                contentToAdd?.trimEnd();
                contentToAdd += "<br>";
            }
            html += contentToAdd;
        } else if (result.type === "LIST_ITEM") {
            if (insideBlock !== "LIST") {
                insideBlock = "LIST";
                html += "<ul class=\"text-body-1\">\n";
            }
            html += `    <li>${parseItalicAndBold(result.content!)}</li>\n`;
        }
    });

    return html;
}

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
