import { describe, expect, it } from "vitest";
import { AccessibilityTreeNode, OrderedAxNode } from "./types";
import { nodesToA11yTree } from "./nodesToA11yTree";

describe("nodesToA11yTree", () => {
    it("simple", () => {
        let nodes: OrderedAxNode = {
            node: {
                nodeId: "1",
                name: {value: "TestPage"},
                childIds: ["2", "3"]
            },
            children: [
                {
                    node: {
                        nodeId: "2",
                        name: {value: "Header"},
                    },
                    children: [
                        {
                            node: {
                                nodeId: "5",
                                name: {value: "Titel der Seite"}
                            }
                        },
                        {
                            node: {
                                nodeId: "6",
                                name: {value: "Home button"}
                            }
                        }
                    ]
                },
                {
                    node: {
                        nodeId: "3",
                        name: {value: "Main"},
                        childIds: ["4"]
                    },
                    children: [
                        {
                            node: {
                                nodeId: "4",
                                name: {value: "Test Button"},
                            }
                        }
                    ]
                },
            ]
        };

        let exp = {
            name: "TestPage",
            children: [
                {
                    name: "Header",
                    children: [
                        {
                            name: "Titel der Seite"
                        },
                        {
                            name: "Home button"
                        }
                    ]
                },
                {
                    name: "Main",
                    children: [
                        {
                            name: "Test Button"
                        }
                    ]
                }
            ]
        };

        let res = nodesToA11yTree(nodes);

        expect(res).toEqual(exp);
    });
    it("map attributes", () => {
        let nodes: OrderedAxNode = {
            node: {
                nodeId: "1",
                name: {value: "TestPage"},
                role: {type: "internalRole", value: "RootWebArea"},
                childIds: ["2", "3"]
            },
            children: [
                {
                    node: {
                        nodeId: "t",
                        ignored: true
                    },
                    children: [
                        {
                            node: {
                                nodeId: "2",
                                role: {type: "role", value: "banner"},
                                name: {value: "Header"},
                            },
                            children: [
                                {
                                    node: {
                                        nodeId: "5",
                                        role: {type: "role", value: "heading"},
                                        properties: [{name: "level", value: {type: "integer", value: 1}}],
                                        name: {value: "Titel der Seite"}
                                    }
                                },
                                {
                                    node: {
                                        nodeId: "6",
                                        role: {type: "role", value: "button"},
                                        properties: [{
                                            name: "focusable",
                                            value: {type: "booleanOrUndefined", value: true}
                                        },
                                            {
                                                name: "focused",
                                                value: {type: "booleanOrUndefined", value: true}
                                            }
                                        ],
                                        name: {value: "Home button"}
                                    }
                                }
                            ]
                        },
                        {
                            node: {
                                nodeId: "3",
                                name: {value: "Main"},
                                childIds: ["4"],
                                role: {type: "role", value: "main"},
                            },
                            children: [
                                {
                                    node: {
                                        nodeId: "4",
                                        name: {value: "Test Button"},
                                        role: {type: "role", value: "button"},
                                        properties: [
                                            {
                                                name: "disabled",
                                                value: {type: "booleanOrUndefined", value: true}
                                            }
                                        ]
                                    }
                                }
                            ]
                        },
                    ]
                }
            ]
        };

        let exp: AccessibilityTreeNode = {
            name: "TestPage",
            role: "RootWebArea",
            children: [
                {
                    ignored: true,
                    children: [
                        {
                            name: "Header",
                            role: "banner",
                            children: [
                                {
                                    name: "Titel der Seite",
                                    role: "heading",
                                    level: 1
                                },
                                {
                                    name: "Home button",
                                    role: "button",
                                    focusable: true,
                                    focused: true
                                }
                            ]
                        },
                        {
                            name: "Main",
                            role: "main",
                            children: [
                                {
                                    name: "Test Button",
                                    role: "button",
                                    disabled: true,
                                }
                            ]
                        }
                    ]
                }
            ]
        };

        let res = nodesToA11yTree(nodes);

        expect(res).toEqual(exp);
    });
});
