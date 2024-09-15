import { AccessibilityTreeNode, AccessibilityTreeResult } from "./types";
import orderNodes from "./orderNodes";
import { nodesToA11yTree } from "./nodesToA11yTree";

export async function getAccessibilityTree(client: any): Promise<AccessibilityTreeResult> {
    const {
        nodes
    } = await client.send("Accessibility.getFullAXTree");

    const tree = analyseNodeTree(nodes);

    // visualizeNodeTree(tree);

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

export function visualizeChildNodes(nodes: AccessibilityTreeNode[], level: number) {
    nodes.forEach(n => {
        console.log(`${"\t".repeat(level)} ${n.ignored ? "(ignored)" : ""} Name: ${n.name}, Role: ${n.role}, ${!!n.hasPopup ? `hasPopup: ${n.hasPopup}, ` : ""} Hidden: ${n.hidden}`);
        if (n.children && n.children?.length > 0) {
            visualizeChildNodes(n.children, level + 1);
        }
    });
}