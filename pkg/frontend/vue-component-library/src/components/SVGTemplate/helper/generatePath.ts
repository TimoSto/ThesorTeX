import {PathPart, SVGPartial, Vector} from "./SVG";

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

export function generateTransform(partial: SVGPartial): string {
    let angle = partial.angle ? partial.angle : 0;
    let scale = partial.scale ? partial.scale : "1";
    let translate = "0 0";
    if (partial.translate) {
        const rotated = turnVector(partial.translate, angle);
        translate = `${rotated.x} ${rotated.y}`;
    }

    let tr = `scale(${scale}) rotate(${angle}, 0, 0) translate(${translate})`;

    console.log(tr);
    return tr;
}

export function turnVector(vector: Vector, deg: number): Vector {
    // see: https://studyflix.de/mathematik/drehmatrix-3814

    const rads = deg * Math.PI / 180;
    const turned: Vector = {
        x: Math.cos(rads) * vector.x - Math.sin(rads) * vector.y,
        y: Math.sin(rads) * vector.x + Math.cos(rads) * vector.y,
    };

    turned.x = Math.round(turned.x * 1000) / 1000;
    turned.y = Math.round(turned.y * 1000) / 1000;

    return turned;
}
