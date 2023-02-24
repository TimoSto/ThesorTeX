export interface VersionInfo {
    Name: string;
    Date: string;
}

export interface VersionData {
    Tool: VersionInfo[];
    ThesisTemplate: VersionInfo[],
    CvTemplate: VersionInfo[]
}

export default async function GetToolVersions(): Promise<VersionData> {
    const resp = await fetch("/versions");

    if (resp.ok) {
        return await resp.json();
    }

    return {} as VersionData;
}
