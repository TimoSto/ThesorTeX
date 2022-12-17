import {ConfigData} from "@/pages/app/api/config/ConfigData";

export default async function ReadConfigData(): Promise<ConfigData> {
  const resp = await fetch('/app/config');

  return await resp.json() as ConfigData
}
