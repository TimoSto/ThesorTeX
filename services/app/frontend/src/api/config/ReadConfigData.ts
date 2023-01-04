import {ConfigData} from "./ConfigData";

export default async function ReadConfigData(): Promise<ConfigData> {
  const resp = await fetch('/app/config');

  return await resp.json() as ConfigData
}
