import {PathPart} from "./SVG";

export function generatePath(parts: PathPart[]): string {
    if (!parts[0].vector) {
        throw Error("first point has to be a vector");
    }

    let path = `M ${parts[0].vector.x} ${parts[0].vector.y} `;

    parts.shift();

    parts.forEach(p => {
        if (p.vector) {
            path += `L${p.vector.x},${p!.vector.y} `;
        } else if (p!.arc) {
            if (!p.arc.radiusY) {
                p.arc.radiusY = p.arc.radius;
            }
            path += `A${p.arc.radius},${p.arc.radiusY} ${p.arc.rotation} 0 ${p.arc.clockwise ? "0" : "1"} ${p.arc.end.x} ${p.arc.end.y} `;
        }
    });

    path += "z";


    return path;
}
