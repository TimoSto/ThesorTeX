export interface VersionInfo {
    Name: string;
    Date: string;
}

export default async function GetToolVersions(): Promise<VersionInfo[]> {
    const resp = await fetch("/versions/tool");

    if (resp.ok) {
        return await resp.json();
    }

    return [];
}
