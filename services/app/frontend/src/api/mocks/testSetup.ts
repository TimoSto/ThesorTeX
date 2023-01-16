import {server} from "./browser";
import {afterAll, afterEach, beforeAll} from "vitest";
import { fetch } from 'cross-fetch';

// Add `fetch` polyfill.
global.fetch = fetch

beforeAll(async () => {
    await server.listen();
});

afterEach(async () => {
    await server.resetHandlers();
})

afterAll(async () => {
    await server.close();
})
