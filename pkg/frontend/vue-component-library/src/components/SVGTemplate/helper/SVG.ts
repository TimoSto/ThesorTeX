export interface TemplateSVG {
    partials: SVGPartial[],
    viewport: string,
}

export interface SVGPartial {
    parts: PathPart[],
    strokeColor: string,
    strokeWidth: string,
    fillColor: string,
    angle: number,
    scale?: number
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
    xEnd: number;
    yEnd: number;
}
