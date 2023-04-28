export default function getFullUri(path: string): string {
    return new URL(path, process.env.SYSTEM_BASE_URL).href;
}
