import {browser, sut} from "./setup";

export default async function globalTeardown() {
    sut?.kill();
    await browser.close();
};