
export interface BibCategory {
  Name:                  string
  CitaviCategory:            string
  CitaviFilters: string[] //z.B. nur dieser Typ wenn doi existiert
  Fields:                Field[]
  CiteFields:            Field[]
  Example: string
}

export interface Field {
  Field:           string
  Italic:            boolean
  Prefix:           string
  Suffix:           string
  TexValue:         boolean
  CitaviAttributes: string[]
}
