export default async function GetToolVersions(): Promise<string[]> {
    const resp = await fetch("/versions/tool");

    if (resp.ok) {
        return await resp.json();
    }

    return [];
}
