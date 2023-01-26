import path from "path";
import fs from "fs";

export function ReadFile(file: string): string {
    const filePath = path.join(__dirname, file);

    return fs.readFileSync(filePath, {encoding: "utf-8"});
}
