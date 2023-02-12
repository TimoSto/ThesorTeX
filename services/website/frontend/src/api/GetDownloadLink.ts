export default function getDownloadLink(version: string, os: string): string {
    version = version.split(" ")[0];
    return `https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/${version}/${os}/ThesorTeX.zip`;
}