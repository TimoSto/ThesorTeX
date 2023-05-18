import {ChildProcessWithoutNullStreams} from "child_process";
import {StartupBrowser, StartupSutIfExists} from "../helpers/startup";
import {ChromiumBrowser} from "playwright";

export let sut: ChildProcessWithoutNullStreams;
export let browser: ChromiumBrowser;

export default async function globalSetup() {
    browser = await StartupBrowser();
    sut = await StartupSutIfExists();
};