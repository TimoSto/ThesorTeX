import {test} from "@playwright/test";
import fs from "fs";

test("get mascot looking left", async ({page}) => {
    await page.goto("http://localhost:6006/iframe.html?args=&id=mascot--left&viewMode=story");
    await new Promise(resolve => {
        setTimeout(resolve, 1500);
    });
    const content = await page.content();

    const svgHtml = content.substring(content.indexOf("<svg"), content.indexOf("</svg>") + 6);

    fs.writeFile("../mascot_looking_left.svg", svgHtml, (err) => {
        if (err) {
            throw err;
        }
    });
});

test("get mascot looking right", async ({page}) => {
    await page.goto("http://localhost:6006/iframe.html?args=&id=mascot--right&viewMode=story");
    await new Promise(resolve => {
        setTimeout(resolve, 1500);
    });
    const content = await page.content();

    const svgHtml = content.substring(content.indexOf("<svg"), content.indexOf("</svg>") + 6);

    fs.writeFile("../mascot_looking_right.svg", svgHtml, (err) => {
        if (err) {
            throw err;
        }
    });
});
