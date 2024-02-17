import {sut} from "./global-setup";

export default function globalTeardown() {
    sut?.kill();
};
