export interface Vector {
    x: number;
    y: number;
}

export default function turnVector(vector: Vector, deg: number): Vector {
    // see: https://studyflix.de/mathematik/drehmatrix-3814
    
    const rads = deg * Math.PI / 180;
    return {
        x: Math.cos(rads) * vector.x - Math.sin(rads) * vector.y,
        y: Math.sin(rads) * vector.x + Math.cos(rads) * vector.y,
    };
}
