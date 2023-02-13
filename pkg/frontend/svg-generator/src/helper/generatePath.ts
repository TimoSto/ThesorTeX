import turnVector, {Vector} from "./turnVector";

//TODO: optional param turnPoint (default 0,0)
export default function generatePath(points: Vector[], angle: number) {
    const startPoint = turnVector(points[0], angle);

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

    const vectors = points.map(v => {
        const relVec = {
            x: v.x,
            y: v.y
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

        return result;
    });

    vectors.forEach(v => {
        p += `L${v.x},${v.y} `;
    });

    p += "z";

    return p;
}
