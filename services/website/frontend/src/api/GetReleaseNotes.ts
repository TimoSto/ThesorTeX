export default async function GetReleaseNotes(target: string): Promise<string> {
    const parts = target.split(" - ");
    const resp = await fetch(`https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/${parts[0]}/${parts[1]}/ReleaseNotes.md`);

    if (!resp.ok) {
        return "# Could not read release notes";
    }

    return resp.text();
}
