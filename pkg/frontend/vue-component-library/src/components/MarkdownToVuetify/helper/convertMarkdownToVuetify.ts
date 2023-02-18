import {analyseLine} from "./analyseLine";

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

        if (result.type === "H1") {
            html += `<h1 class="text-h1">${result.content}</h1>` + "\n";
        } else if (result.type === "H2") {
            html += `<h2 class="text-h2">${result.content}</h2>` + "\n";
        } else if (result.type === "H3") {
            html += `<h3 class="text-h3">${result.content}</h3>` + "\n";
        } else if (result.type === "H4") {
            html += `<h4 class="text-h4">${result.content}</h4>` + "\n";
        } else if (result.type === "H5") {
            html += `<h5 class="text-h5">${result.content}</h5>` + "\n";
        } else if (result.type === "TEXT") {
            if (insideBlock !== "PLAIN_TEXT") {
                insideBlock = "PLAIN_TEXT";
                html += "<p class=\"text-body-1\">";
            }
            let contentToAdd = result.content;
            if (result.content?.endsWith("  ")) {
                contentToAdd?.trimEnd();
                contentToAdd += "<br>";
            }
            html += contentToAdd;
        } else if (result.type === "EMPTY") {
            // if (insideBlock === "PLAIN_TEXT") {
            //     html += "</p>\n";
            //     insideBlock = false;
            // }
        }
    });

    return html;
}
