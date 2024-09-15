import { describe, expect, it } from "vitest";
import { AxNode, OrderedAxNode } from "./types";
import orderNodes from "./orderNodes";

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
