import fs from "fs";
import {test} from "@playwright/test";

const storyName = process.env.STORY_NAME;
const dist = process.env.STORY_DIST;

test(`exracting png and svg for stoy ${storyName}`, async ({page}) => {
    await page.goto(`http://localhost:6006/iframe.html?args=&id=${storyName}&viewMode=story`);
    await new Promise(resolve => {
        setTimeout(resolve, 1500);
    });
    const content = await page.content();

    const svgHtml = content.substring(content.indexOf("<svg"), content.indexOf("</svg>") + 6);

    fs.writeFile(dist, svgHtml, (err) => {
        if (err) {
            throw err;
        }
    });
})
