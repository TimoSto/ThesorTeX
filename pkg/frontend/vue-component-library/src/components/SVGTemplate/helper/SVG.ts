export interface TemplateSVG {
    partials: SVGPartial[],
    viewBox: string,
    width: number,
    height: number,
}

export interface SVGPartial {
    parts: PathPart[][],
    strokeColor: string,
    strokeWidth: string,
    fillColor: string,
    angle?: number,
    scale?: number,
    translate?: Vector,
}

export interface PathPart {
    vector?: Vector;
    arc?: Arc;
}

export interface Vector {
    x: number;
    y: number;
}

export interface Arc {
    radius: number;
    radiusY?: number;
    rotation: number;
    clockwise: boolean;
    end: Vector;
}

export function MirrorVectorsY(parts: PathPart[]): PathPart[] {
    parts.forEach(p => {
        if (p.vector) {
            p.vector.x *= -1;
        }
    });

    return parts;
}
