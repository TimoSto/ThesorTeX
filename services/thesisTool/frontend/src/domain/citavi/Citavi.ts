export const categories: string[] = [
    "article",
    "book",
    "booklet",
    "collection",
    "conference",
    "inbook",
    "incollection",
    "inproceedings",
    "manual",
    "masterthesis",
    "misc",
    "phdthesis",
    "proceedings",
    "techreport",
    "unpublished"
];

export const attributes: string[] = [
    "address",
    "author",
    "booktitle",
    "doi",
    "edition",
    "editor",
    "journal",
    "pages",
    "title",
    "publisher",
    "volume",
    "year",
];

export interface AttributeValue {
    Attr: string;
    Value: string;
}

export interface CitaviEntry {
    Key: string;
    Category: string;
    Attributes: AttributeValue[];
}
