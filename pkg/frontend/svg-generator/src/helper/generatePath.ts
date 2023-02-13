import turnVector, {Arc, Vector} from "./turnVector";

export interface PathPart {
    vector?: Vector;
    arc?: Arc;
}

//TODO: optional param turnPoint (default 0,0)
export default function generatePath(points: PathPart[], angle: number) {
    const startPoint = turnVector(points[0].vector!, angle);

    points.shift();

    let p = `M${startPoint.x},${startPoint.y} `;

    // let prevVec = {
    //     x: startPoint.x,
    //     y: startPoint.y
    // };
    //
    // let prevVecTurned = {
    //     x: startPoint.x,
    //     y: startPoint.y
    // };

    const vectors = points.map((v) => {
        if (v.vector) {
            const relVec = {
                x: v.vector.x,
                y: v.vector.y
            };

            const turned = turnVector(relVec, angle);

            const result = {
                x: turned.x,
                y: turned.y
            };

            // prevVec.x = v.x;
            // prevVec.y = v.y;
            //
            // prevVecTurned.x = result.x;
            // prevVecTurned.y = result.y;

            return {
                vector: result
            };
        }
    });

    vectors.forEach(v => {
        if (v!.vector) {
            p += `L${v!.vector.x},${v!.vector.y} `;
        }
    });

    p += "z";

    return p;
}
