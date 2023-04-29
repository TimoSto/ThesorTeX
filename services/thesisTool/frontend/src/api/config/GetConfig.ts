import Config from "../../domain/config/Config";
import {host} from "../config";

export default async function GetConfig(): Promise<Config> {
    const resp = await fetch(`${host}/getConfig`);

    if (resp.ok) {
        const data = await resp.json();
        return data as Config;
    }

    return {} as Config;
}
