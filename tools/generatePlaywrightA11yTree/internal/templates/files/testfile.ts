import { test } from '@playwright/test';

test('test', async ({page}) => {
    await page.goto('about:blank');
    await page.goto('chrome-error://chromewebdata/');
    await page.goto('http://localhost:8449/');
    await page.goto('http://localhost:8449/#/');
    await page.getByText('Vorlage für eine Abschluss- / Hausarbeit').first().click();
    await page.getByText('Vorlage für eine Abschluss- / Hausarbeit').first().click();
    await page.getByRole('link', {name: 'Zu den Downloads'}).first().click();
    await page.locator('td:nth-child(4) > .v-btn').first().click();
});