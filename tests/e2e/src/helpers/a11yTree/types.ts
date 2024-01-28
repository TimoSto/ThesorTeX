export type AccessibilityTreeResult = {
    PageTitle?: string
    Nodes: AccessibilityTreeNode[]
}

export type AccessibilityTreeNode = {
    id?: number,
    ignored?: boolean,
    role?: string,
    name?: string,
    description?: string,
    focusable?: boolean,
    focused?: boolean,
    disabled?: boolean,
    invalid?: boolean,
    children?: AccessibilityTreeNode[]
}
