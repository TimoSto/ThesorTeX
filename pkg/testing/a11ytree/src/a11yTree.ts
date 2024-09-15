import { AccessibilityTreeNode, AccessibilityTreeResult } from "./types";
import orderNodes from "./orderNodes";
import { nodesToA11yTree } from "./nodesToA11yTree";

export async function getAccessibilityTree(client: any): Promise<AccessibilityTreeResult> {
    const {
        nodes
    } = await client.send("Accessibility.getFullAXTree");

    const tree = analyseNodeTree(nodes);

    return tree;
}

function analyseNodeTree(nodes: any[] /*TODO: use Protocol.Accessibility.Node[] as type*/): AccessibilityTreeResult {
    let res: AccessibilityTreeResult = {
        Nodes: []
    };

    if (nodes[0].role?.value === "RootWebArea") {
        res.PageTitle = nodes[0].name.value;
    }

    let orderedNodes = orderNodes(nodes, 0);

    res.Nodes = [nodesToA11yTree(orderedNodes)];

    return res;
}

function analyseChildNodes(i: number, nodes: any[]): AccessibilityTreeNode[] {
    let node: AccessibilityTreeNode = {
        name: nodes[i].name?.value,
        role: nodes[i].role?.value,
        ignored: nodes[i].ignored.value,
        children: [],
    };

    // if (nodes[i].ignored) {
    //     if (nodes[i].childIds?.length > 0) {
    //         return analyseChildNodes(parseInt(nodes[i].childIds[0]), nodes);
    //     }
    // }

    nodes[i].childIds.forEach((cid: string) => {
        const n = parseInt(cid);

        if (!isNaN(n)) {
            node.children?.push(...analyseChildNodes(n - 1, nodes));
        }
    });

    return [node];


}

function visualizeNodeTree(tree: AccessibilityTreeResult) {
    console.log("----------------");
    console.log(tree.PageTitle);
    console.log("----");
    visualizeChildNodes(tree.Nodes, 1);
}

function visualizeChildNodes(nodes: AccessibilityTreeNode[], level: number) {
    nodes.forEach(n => {
        console.log(`${"\t".repeat(level)} ${n.ignored ? "(ignored)" : ""} Name: ${n.name}, Role: ${n.role}, ${!!n.hasPopup ? `hasPopup: ${n.hasPopup}, ` : ""} Hidden: ${n.hidden}`);
        if (n.children && n.children?.length > 0) {
            visualizeChildNodes(n.children, level + 1);
        }
    });
}