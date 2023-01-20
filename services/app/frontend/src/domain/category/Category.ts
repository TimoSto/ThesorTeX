export interface Category {
    Name: string
    CtiaviCategory: string
    BibFields: Field[]
    CiteFields: Field[]
    CitaviFilter: string[]
}

export interface Field {
    Name: string
    Format: format
    CitaviMapping: string[]
}

export interface format {
    Prefix: string
    Suffix: string
    Style: string
    Preformatted: boolean
}
