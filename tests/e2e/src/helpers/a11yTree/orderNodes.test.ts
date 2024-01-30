import {describe, expect, it} from "vitest";
import {AxNode, OrderedAxNode} from "./types";
import orderNodes from "./orderNodes";

let nodes = [
    {
        nodeId: "1",
        ignored: false,
        role: {type: "internalRole", value: "RootWebArea"},
        name: {
            type: "computedString",
            value: "Test page",
            sources: [{type: "relatedElement", "attribute": "aria-labelledby"}, {
                type: "attribute",
                "attribute": "aria-label"
            }, {type: "attribute", "attribute": "aria-label", "superseded": true}, {
                type: "relatedElement",
                value: {type: "computedString", value: "Test page"},
                "nativeSource": "title"
            }, {type: "attribute", "attribute": "title", "superseded": true}]
        },
        properties: [{
            name: "focusable",
            value: {type: "booleanOrUndefined", value: true}
        }, {name: "focused", value: {type: "booleanOrUndefined", value: true}}],
        childIds: ["2"],
        backendDOMNodeId: 2,
        "frameId": "E9F7BBB9721202A9B38DA1AC503EF39C"
    }, {
        nodeId: "2",
        ignored: true,
        ignoredReasons: [{name: "uninteresting", value: {type: "boolean", value: true}}],
        role: {type: "role", value: "none"},
        chromeRole: {type: "internalRole", value: 0},
        parentId: "1",
        childIds: ["3"],
        backendDOMNodeId: 3
    }, {
        nodeId: "3",
        ignored: true,
        ignoredReasons: [{name: "uninteresting", value: {type: "boolean", value: true}}],
        role: {type: "role", value: "none"},
        chromeRole: {type: "internalRole", value: 0},
        parentId: "2",
        childIds: ["4", "5"],
        backendDOMNodeId: 4
    }, {
        nodeId: "4",
        ignored: false,
        role: {type: "role", value: "banner"},
        chromeRole: {type: "internalRole", value: 94},
        name: {
            type: "computedString",
            value: "",
            sources: [{type: "relatedElement", "attribute": "aria-labelledby"}, {
                type: "attribute",
                "attribute": "aria-label"
            }, {type: "attribute", "attribute": "title"}]
        },
        properties: [],
        parentId: "3",
        childIds: ["6"],
        backendDOMNodeId: 5
    }, {
        nodeId: "5",
        ignored: false,
        role: {type: "role", value: "main"},
        chromeRole: {type: "internalRole", value: 118},
        name: {
            type: "computedString",
            value: "",
            sources: [{type: "relatedElement", "attribute": "aria-labelledby"}, {
                type: "attribute",
                "attribute": "aria-label"
            }, {type: "attribute", "attribute": "title"}]
        },
        properties: [],
        parentId: "3",
        childIds: ["7"],
        backendDOMNodeId: 6
    }, {
        nodeId: "6",
        ignored: false,
        role: {type: "role", value: "heading"},
        chromeRole: {type: "internalRole", value: 96},
        name: {
            type: "computedString",
            value: "Test Page For A11y tree",
            sources: [{type: "relatedElement", "attribute": "aria-labelledby"}, {
                type: "attribute",
                "attribute": "aria-label"
            }, {
                type: "contents",
                value: {type: "computedString", value: "Test Page For A11y tree"}
            }, {type: "attribute", "attribute": "title", "superseded": true}]
        },
        properties: [{name: "level", value: {type: "integer", value: 3}}],
        parentId: "4",
        childIds: ["8"],
        backendDOMNodeId: 7
    }, {
        nodeId: "7",
        ignored: false,
        role: {type: "role", value: "paragraph"},
        chromeRole: {type: "internalRole", value: 133},
        name: {
            type: "computedString",
            value: "",
            sources: [{type: "relatedElement", "attribute": "aria-labelledby"}, {
                type: "attribute",
                "attribute": "aria-label"
            }, {type: "attribute", "attribute": "title"}]
        },
        properties: [],
        parentId: "5",
        childIds: ["9"],
        backendDOMNodeId: 8
    }, {
        nodeId: "8",
        ignored: false,
        role: {type: "internalRole", value: "StaticText"},
        chromeRole: {type: "internalRole", value: 158},
        name: {
            type: "computedString",
            value: "Test Page For A11y tree",
            sources: [{
                type: "contents",
                value: {type: "computedString", value: "Test Page For A11y tree"}
            }]
        },
        properties: [],
        parentId: "6",
        childIds: [],
        backendDOMNodeId: 9
    }, {
        nodeId: "9",
        ignored: false,
        role: {type: "internalRole", value: "StaticText"},
        chromeRole: {type: "internalRole", value: 158},
        name: {
            type: "computedString",
            value: "This is some content",
            sources: [{type: "contents", value: {type: "computedString", value: "This is some content"}}]
        },
        properties: [],
        parentId: "7",
        childIds: [],
        backendDOMNodeId: 10
    }];

describe("OrderNodes", () => {
    it("simple", () => {
        let n: AxNode[] = [
            {
                nodeId: "1",
                childIds: ["2", "3"]
            },
            {
                nodeId: "2",
            },
            {
                nodeId: "3",
                childIds: ["4"]
            },
            {
                nodeId: "4",
            }
        ];

        let res = orderNodes(n, 0);

        let exp: OrderedAxNode = {
            node: {
                nodeId: "1",
                childIds: ["2", "3"]
            },
            children: [
                {
                    node: {
                        nodeId: "2",
                    }
                },
                {
                    node: {
                        nodeId: "3",
                        childIds: ["4"]
                    },
                    children: [
                        {
                            node: {
                                nodeId: "4"
                            }
                        }
                    ]
                },
            ]
        };

        expect(res).toEqual(exp);
    });
    it("more complex", () => {
        let n: AxNode[] = [
            {
                nodeId: "1",
                childIds: ["2", "3"]
            },
            {
                nodeId: "2",
                childIds: ["4", "5"]
            },
            {
                nodeId: "3",
                childIds: ["6"]
            },
            {
                nodeId: "4",
            },
            {
                nodeId: "5",
                childIds: ["7", "8"]
            },
            {
                nodeId: "6",
            },
            {
                nodeId: "7",
            },
            {
                nodeId: "8",
            }
        ];

        let res = orderNodes(n, 0);

        let exp: OrderedAxNode = {
            node: {
                nodeId: "1",
                childIds: ["2", "3"]
            },
            children: [
                {
                    node: {
                        nodeId: "2",
                        childIds: ["4", "5"]
                    },
                    children: [
                        {
                            node: {
                                nodeId: "4"
                            }
                        },
                        {
                            node: {
                                nodeId: "5",
                                childIds: ["7", "8"]
                            },
                            children: [
                                {
                                    node: {
                                        nodeId: "7"
                                    }
                                },
                                {
                                    node: {
                                        nodeId: "8"
                                    }
                                }
                            ]
                        }
                    ]
                },
                {
                    node: {
                        nodeId: "3",
                        childIds: ["6"]
                    },
                    children: [
                        {
                            node: {
                                nodeId: "6"
                            }
                        }
                    ]
                },
            ]
        };

        expect(res).toEqual(exp);
    });
});
