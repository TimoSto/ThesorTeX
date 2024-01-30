import {AxNode, OrderedAxNode} from "./types";

export default function orderNodes(nodes: AxNode[], i: number): OrderedAxNode {
    let res: OrderedAxNode = {
        node: nodes[i],
        children: nodes[i].childIds && nodes[i].childIds!.length > 0 ? [] : undefined
    };

    nodes[i].childIds?.forEach((childId) => {
        for (let n = i + 1; n < nodes.length; n++) {
            if (nodes[n].nodeId === childId) {
                const childResult = orderNodes(nodes, n);
                res.children?.push(childResult);
            }
        }
    });


    return res;
}
