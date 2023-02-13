export interface Vector {
    x: number;
    y: number;
}

export interface Arc {
    radius: number;
    rotation: number;
    clockwise: boolean;
    xEnd: number;
    yEnd: number;
}

export default function turnVector(vector: Vector, deg: number): Vector {
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
