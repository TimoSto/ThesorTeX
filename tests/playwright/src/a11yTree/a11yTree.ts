import { CDPSession } from "@playwright/test";
import { AccessibilityTreeNode, AccessibilityTreeResult } from "./types";
import orderNodes from "./orderNodes";
import { nodesToA11yTree } from "./nodesToA11yTree";

export async function getAccessibilityTree(client: CDPSession): Promise<AccessibilityTreeResult> {
    const {
        nodes
    } = await client.send("Accessibility.getFullAXTree");

    const tree = analyseNodeTree(nodes);

    visualizeNodeTree(tree);

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

    console.log(nodes[i].childIds, nodes[i].childIds?.length, nodes[i].ignored);
    // if (nodes[i].childIds?.length > 0 && nodes[i].ignored) {
    //     console.log(nodes[i]);
    //     // ignore this level
    //     return node.children!;
    // } else {
    //
    // }
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

type AccessibilityTree = {
    PageTitle: string
}

type A11yTreeNode = {}

class CRAXNode {
    private _payload: any;
    private _children: any[];
    private _richlyEditable: boolean;
    private _editable: boolean;
    private _focusable: boolean;
    private _expanded: boolean;
    private _hidden: boolean;
    private _name: any;
    private _role: any;
    private _cachedHasFocusableChild: any;
    private _client: any;

    constructor(client: CDPSession, payload: any) {
        this._payload = void 0;
        this._children = [];
        this._richlyEditable = false;
        this._editable = false;
        this._focusable = false;
        this._expanded = false;
        this._hidden = false;
        this._name = void 0;
        this._role = void 0;
        this._cachedHasFocusableChild = void 0;
        this._client = void 0;
        this._client = client;
        this._payload = payload;
        this._name = this._payload.name ? this._payload.name.value : "";
        this._role = this._payload.role ? this._payload.role.value : "Unknown";

        for (const property of this._payload.properties || []) {
            if (property.name === "editable") {
                this._richlyEditable = property.value.value === "richtext";
                this._editable = true;
            }

            if (property.name === "focusable") this._focusable = property.value.value;
            if (property.name === "expanded") this._expanded = property.value.value;
            if (property.name === "hidden") this._hidden = property.value.value;
        }
    }

    _isPlainTextField() {
        if (this._richlyEditable) return false;
        if (this._editable) return true;
        return this._role === "textbox" || this._role === "ComboBox" || this._role === "searchbox";
    }

    _isTextOnlyObject() {
        const role = this._role;
        return role === "LineBreak" || role === "text" || role === "InlineTextBox" || role === "StaticText";
    }

    _hasFocusableChild() {
        if (this._cachedHasFocusableChild === undefined) {
            this._cachedHasFocusableChild = false;

            for (const child of this._children) {
                if (child._focusable || child._hasFocusableChild()) {
                    this._cachedHasFocusableChild = true;
                    break;
                }
            }
        }

        return this._cachedHasFocusableChild;
    }

    children() {
        return this._children;
    }

    async _findElement(element: any) {
        const objectId = element._objectId;
        const {
            node: {
                backendNodeId
            }
        } = await this._client.send("DOM.describeNode", {
            objectId
        });
        const needle = this.find((node: any) => node._payload.backendDOMNodeId === backendNodeId);
        return needle || null;
    }

    find(predicate: any) {
        if (predicate(this)) return this;

        for (const child of this._children) {
            const result = child.find(predicate);
            if (result) return result;
        }

        return null;
    }

    isLeafNode() {
        if (!this._children.length) return true; // These types of objects may have children that we use as internal
        // implementation details, but we want to expose them as leaves to platform
        // accessibility APIs because screen readers might be confused if they find
        // any children.

        if (this._isPlainTextField() || this._isTextOnlyObject()) return true; // Roles whose children are only presentational according to the ARIA and
        // HTML5 Specs should be hidden from screen readers.
        // (Note that whilst ARIA buttons can have only presentational children, HTML5
        // buttons are allowed to have content.)

        switch (this._role) {
            case "doc-cover":
            case "graphics-symbol":
            case "img":
            case "Meter":
            case "scrollbar":
            case "slider":
            case "separator":
            case "progressbar":
                return true;

            default:
                break;
        } // Here and below: Android heuristics


        if (this._hasFocusableChild()) return false;
        if (this._focusable && this._role !== "WebArea" && this._role !== "RootWebArea" && this._name) return true;
        if (this._role === "heading" && this._name) return true;
        return false;
    }

    isControl() {
        switch (this._role) {
            case "button":
            case "checkbox":
            case "ColorWell":
            case "combobox":
            case "DisclosureTriangle":
            case "listbox":
            case "menu":
            case "menubar":
            case "menuitem":
            case "menuitemcheckbox":
            case "menuitemradio":
            case "radio":
            case "scrollbar":
            case "searchbox":
            case "slider":
            case "spinbutton":
            case "switch":
            case "tab":
            case "textbox":
            case "tree":
                return true;

            default:
                return false;
        }
    }

    isInteresting(insideControl: any) {
        const role = this._role;
        if (role === "Ignored" || this._hidden) return false;
        if (this._focusable || this._richlyEditable) return true; // If it's not focusable but has a control role, then it's interesting.

        if (this.isControl()) return true; // A non focusable child of a control is not interesting

        if (insideControl) return false;
        return this.isLeafNode() && !!this._name;
    }

    normalizedRole() {
        switch (this._role) {
            case "RootWebArea":
                return "WebArea";

            case "StaticText":
                return "text";

            default:
                return this._role;
        }
    }

    serialize() {
        const properties = new Map();

        for (const property of this._payload.properties || []) properties.set(property.name.toLowerCase(), property.value.value);

        if (this._payload.description) properties.set("description", this._payload.description.value);
        const node: any = {
            role: this.normalizedRole(),
            name: this._payload.name ? this._payload.name.value || "" : ""
        };
        const userStringProperties = ["description", "keyshortcuts", "roledescription", "valuetext"];

        for (const userStringProperty of userStringProperties) {
            if (!properties.has(userStringProperty)) continue;
            node[userStringProperty] = properties.get(userStringProperty);
        }

        const booleanProperties = ["disabled", "expanded", "focused", "modal", "multiline", "multiselectable", "readonly", "required", "selected"];

        for (const booleanProperty of booleanProperties) {
            // WebArea's treat focus differently than other nodes. They report whether their frame  has focus,
            // not whether focus is specifically on the root node.
            if (booleanProperty === "focused" && (this._role === "WebArea" || this._role === "RootWebArea")) continue;
            const value = properties.get(booleanProperty);
            if (!value) continue;
            node[booleanProperty] = value;
        }

        const numericalProperties = ["level", "valuemax", "valuemin"];

        for (const numericalProperty of numericalProperties) {
            if (!properties.has(numericalProperty)) continue;
            node[numericalProperty] = properties.get(numericalProperty);
        }

        const tokenProperties = ["autocomplete", "haspopup", "invalid", "orientation"];

        for (const tokenProperty of tokenProperties) {
            const value = properties.get(tokenProperty);
            if (!value || value === "false") continue;
            node[tokenProperty] = value;
        }

        const axNode = node;

        if (this._payload.value) {
            if (typeof this._payload.value.value === "string") axNode.valueString = this._payload.value.value;
            if (typeof this._payload.value.value === "number") axNode.valueNumber = this._payload.value.value;
        }

        if (properties.has("checked")) axNode.checked = properties.get("checked") === "true" ? "checked" : properties.get("checked") === "false" ? "unchecked" : "mixed";
        if (properties.has("pressed")) axNode.pressed = properties.get("pressed") === "true" ? "pressed" : properties.get("pressed") === "false" ? "released" : "mixed";
        return axNode;
    }

    static createTree(client: CDPSession, payloads: any) {
        const nodeById = new Map();

        for (const payload of payloads) nodeById.set(payload.nodeId, new CRAXNode(client, payload));

        for (const node of nodeById.values()) {
            for (const childId of node._payload.childIds || []) node._children.push(nodeById.get(childId));
        }

        return nodeById.values().next().value;
    }

}