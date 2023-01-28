export default function getFullUri(path: string): string {
    return new URL(path, "http://localhost:8448").href;
}
