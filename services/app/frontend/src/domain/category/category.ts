
export interface Category {
    Name: string
    CtiaviCategory: string
}

export interface field {
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
