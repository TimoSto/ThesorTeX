const unparseValues = [
    ["{\\\"a}", "ä"],
    ["{\\\"o}", "ö"],
    ["{\\\"u}", "ü"],
    ["{{\\ss}}", "ß"],
    ["{{\\textunderscore}}", "{{\\_}}"],
];

export function trimAndParseValue(v: string): string {
    // this parsing happens always
    v = v.trim();
    if (v.charAt(0) === "{" && v.charAt(v.length - 1) === "}") {
        v = v.substring(1, v.length - 1);
    }
    v = v.trim();
    if (v.charAt(0) === "\"" && v.charAt(v.length - 1) === "\"") {
        v = v.substring(1, v.length - 1);
    }
    unparseValues.forEach(val => {
        v = v.replaceAll(val[0], val[1]);
    });

    return v;
}
