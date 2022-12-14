import {ConfigData} from "./ConfigData";

export default async function SaveConfigData(data: ConfigData): Promise<boolean> {
  const resp = await fetch('/app/config', {
    method: 'POST',
    body: JSON.stringify(data)
  });

  return resp.ok;
}
