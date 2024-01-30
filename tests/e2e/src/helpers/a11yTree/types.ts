// given by Accessibility api
export type AxNode = {
    nodeId: string
    ignored?: boolean
    ignoredReasons?: ValueProperty[]
    role?: TypeValue
    chromeRole?: TypeValue
    name?: TypeValueWithSource
    properties?: ValueProperty[]
    childIds?: string[]
    parentId?: string
}

export type TypeValue = {
    type: string
    value: string
}

export type TypeValueWithSource = {
    type?: string
    value: string
    sources?: TypeValue[]
}

export type ValueProperty = {
    type: string
    value: TypeValue
}

// middle result
export type OrderedAxNode = {
    node: AxNode
    children?: OrderedAxNode[]
}

// result
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
    children?: AccessibilityTreeNode[],
    hasFocusableChild?: boolean,
    isTextOnly?: boolean,
    isControl?: boolean
}
