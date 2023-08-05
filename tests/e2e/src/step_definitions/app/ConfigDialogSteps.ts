import {Then, When} from "@cucumber/cucumber";
import {OurWorld} from "../../types";
import {expect} from "@playwright/test";

When("{string} is entered as custom port", async function (this: OurWorld, v: string) {
    await this.page.locator(".v-input", {has: this.page.locator(".v-field-label", {hasText: "Port"})}).locator("input").fill(v);
});

Then("the value for the custom port is {string}", async function (this: OurWorld, v: string) {
    const val = await this.page.locator(".v-input", {has: this.page.locator(".v-field-label", {hasText: "Port"})}).locator("input").inputValue();
    await expect(val).toEqual(v);
});

When("{string} is entered as custom projects path", async function (this: OurWorld, v: string) {
    await this.page.locator(".v-input", {has: this.page.locator(".v-field-label", {hasText: "Directory for projects"})}).locator("input").fill(v);
});

Then("the value for the custom projects path is {string}", async function (this: OurWorld, v: string) {
    const val = await this.page.locator(".v-input", {has: this.page.locator(".v-field-label", {hasText: "Directory for projects"})}).locator("input").inputValue();
    await expect(val).toEqual(v);
});

When("auto open is activated", async function (this: OurWorld) {
    await this.page.locator(".v-selection-control", {has: this.page.locator(".v-label", {hasText: "Open browser on startup"})}).locator("input").click();
});

Then("auto open is enabled", async function (this: OurWorld) {
    const checked = await this.page.locator(".v-selection-control", {has: this.page.locator(".v-label", {hasText: "Open browser on startup"})}).locator("input").inputValue();
    await expect(checked).toEqual("true");
});

When("{string} is set as log level", async function (this: OurWorld, v: string) {
    await this.page.locator(".v-input", {has: this.page.locator(".v-field-label", {hasText: "log level"})}).click();
    await this.page.locator(".v-list-item", {has: this.page.locator(".v-list-item-title", {hasText: v})}).click();
});

When("the value for the log level is {string}", async function (this: OurWorld, v: string) {
    const val = await this.page.locator(".v-input", {has: this.page.locator(".v-field-label", {hasText: "log level"})}).locator("input").inputValue();
    expect(val).toEqual(v);
});
