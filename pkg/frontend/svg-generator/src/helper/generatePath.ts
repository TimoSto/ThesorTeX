import turnVector, {Vector} from "./turnVector";

export default function generatePath(points: Vector[], angle: number) {
    const startPoint = points[0];

    points.shift();

    let p = `M${startPoint.x},${startPoint.y} `;

    let prevVec = {
        x: startPoint.x,
        y: startPoint.y
    };

    let prevVecTurned = {
        x: startPoint.x,
        y: startPoint.y
    };

    const vectors = points.map(v => {
        const relVec = {
            x: v.x - prevVec.x,
            y: v.y - prevVec.y
        };

        const turned = turnVector(relVec, angle);

        const result = {
            x: prevVecTurned.x + turned.x,
            y: prevVecTurned.y + turned.y
        };

        prevVec.x = v.x;
        prevVec.y = v.y;

        prevVecTurned.x = result.x;
        prevVecTurned.y = result.y;

        return result;
    });

    vectors.forEach(v => {
        p += `L${v.x},${v.y} `;
    });

    p += "z";

    return p;
}
