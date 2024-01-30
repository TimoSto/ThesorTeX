import {AccessibilityTreeNode, AxNode, OrderedAxNode} from "./types";

export function nodesToA11yTree(tree: OrderedAxNode): AccessibilityTreeNode {
    // if( !tree.children || tree.children?.length === 0 ) {
    //     return undefined
    // }
    let node = mapNode(tree.node);

    tree.children?.forEach((c: OrderedAxNode) => {
        let childNode = nodesToA11yTree(c);
        if (!node.children) {
            node.children = [];
        }
        node.children.push(childNode);
    });

    console.log(node, tree);

    return node;
}

function mapNode(node: AxNode): AccessibilityTreeNode {
    return {
        name: node.name?.value
    };
}
