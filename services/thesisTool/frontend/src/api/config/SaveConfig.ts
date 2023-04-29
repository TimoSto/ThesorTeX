import {Config} from "../../domain/config/Config";
import {host} from "../config";

export default async function SaveConfig(cfg: Config): Promise<boolean> {
    const resp = await fetch(`${host}/saveConfig`, {
        method: "POST",
        body: JSON.stringify(cfg)
    });

    return resp.ok;
}
