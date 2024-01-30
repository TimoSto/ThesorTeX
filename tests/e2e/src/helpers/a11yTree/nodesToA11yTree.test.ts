import {describe, expect, it} from "vitest";
import {OrderedAxNode} from "./types";
import {nodesToA11yTree} from "./nodesToA11yTree";

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
});
