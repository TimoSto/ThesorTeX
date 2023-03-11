export function getToolDownloadLink(version: string, os: string): string {
    version = version.split(" ")[0];
    return `https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/thesisTool/${version}/${os}/ThesorTeX.zip`;
}

export function getThesisTemplateDownloadLink(version: string): string {
    version = version.split(" ")[0];
    return `https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/thesisTemplate/${version}/ThesisTemplate.zip`;
}

export function getCVTemplateDownloadLink(version: string): string {
    version = version.split(" ")[0];
    return `https://thesortex-artifacts.s3.eu-central-1.amazonaws.com/cvTemplate/${version}/CVTemplate.zip`;
}