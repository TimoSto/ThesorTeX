export default function generateWaveSVG(width: number, height: number, bg: string, waveColor: string, fn: (x: number) => number): string {
    let points = "";
    for (let x = -width / 2; x < width / 2 + 10; x += 10) {
        points += `L ${x} ${fn(x)} `;
    }

    points += `L ${width / 2} ${height / 2 + 10} L ${-width / 2} ${height / 2 + 10} `;
    let svg = `<svg xmlns="http://www.w3.org/2000/svg" fill="${waveColor}" width="${width}px" height="${height}px" viewBox="${-width / 2} ${-height / 2} ${width} ${height}">
        <rect x="${-width / 2}" y="${-height / 2}" width="${width}" height="${height}" fill="${bg}"></rect>
        <path d="M ${-width / 2} ${fn(-width / 2)} ${points}Z" stroke="black" />
</svg>`;

    return svg;
}
