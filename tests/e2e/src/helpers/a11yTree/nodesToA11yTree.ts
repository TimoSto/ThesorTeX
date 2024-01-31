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

    return node;
}

function mapNode(node: AxNode): AccessibilityTreeNode {
    let res: AccessibilityTreeNode = {};

    if (node.name) {
        res.name = node.name.value;
    }

    if (node.role) {
        res.role = node.role?.value as string;
    }

    if (node.ignored) {
        res.ignored = true;
    }

    node.properties?.forEach(prop => {
        if (prop.name === "level") {
            res.level = prop.value.value as number;
        } else if (prop.name === "focusable") {
            res.focusable = prop.value.value as boolean;
        } else if (prop.name === "focused") {
            res.focused = prop.value.value as boolean;
        } else if (prop.name === "disabled") {
            res.disabled = prop.value.value as boolean;
        } else if (prop.name === "hidden") {
            res.hidden = prop.value.value as boolean;
        } else if (prop.name === "level") {
            res.level = prop.value.value as number;
        } else if (prop.name === "hasPopup") {
            res.hasPopup = prop.value.value as string;
        }
    });

    return res;
}
