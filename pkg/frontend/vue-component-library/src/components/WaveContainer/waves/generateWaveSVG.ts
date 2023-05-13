export default function generateWaveSVG(width: number, height: number, fn: (x: number) => number): string {
    let points = "";
    for (let x = -width / 2; x < width / 2; x += 10) {
        points += `L ${x} ${fn(x)} `;
    }

    points += `L ${width / 2} ${height / 2} L ${-width / 2} ${height / 2} `;
    let svg = `<svg xmlns="http://www.w3.org/2000/svg" fill="green" width="${width}px" height="${height}px" viewBox="${-width / 2} ${-height / 2} ${width} ${height}">
        <path d="M ${-width / 2} ${fn(-width / 2)} ${points}Z" stroke="black" />
</svg>`;

    return svg;
}
