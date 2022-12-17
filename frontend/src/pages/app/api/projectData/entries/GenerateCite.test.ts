import {BibCategory} from "@/pages/app/api/projectData/categories/BibCategory";
import {BibEntry} from "@/pages/app/api/projectData/entries/BibEntry";
import GenerateCite from "@/pages/app/api/projectData/entries/GenerateCite";

describe('GenerateCite', () => {
  it('should give correct', () => {
    const c = [
      {
        Name: 'c',
        Fields: [
          {
            Field: 'f1',
            Suffix: ' ',
            Italic: true
          },
          {
            Field: 'f2',
            Prefix: '(',
            Suffix: ').',
            Italic: false
          }
        ],
        CiteFields: [
          {
            Field: 'f3',
            Suffix: ' ',
            Italic: true
          },
          {
            Field: 'f2',
            Prefix: '(',
            Suffix: ').',
            Italic: false
          }
        ]
      },
      {
        Name: 'c2',
        Fields: [
          {
            Field: 'f1',
            Suffix: ' - ',
            Italic: false
          },
          {
            Field: 'f2',
            Prefix: '(',
            Suffix: ').',
            Italic: true
          }
        ],
        CiteFields: [
          {
            Field: 'f2',
            Suffix: ' ',
            Italic: true
          },
          {
            Field: 'f3',
            Prefix: '(',
            Suffix: ').',
            Italic: false
          }
        ]
      },
    ] as BibCategory[];

    const e = {
      Key: 'test',
      Category: 'c2',
      Fields: [
        'mein',
        'eintrag',
        'test'
      ]
    } as BibEntry

    expect(GenerateCite(e, c)).toEqual('<i>eintrag</i> (test).')
  })

  it('should give correct', () => {
    const c = [
      {
        Name: 'c',
        Fields: [
          {
            Field: 'f1',
            Suffix: ' ',
            Italic: true
          },
          {
            Field: 'f2',
            Prefix: '(',
            Suffix: ').',
            Italic: false
          }
        ],
        CiteFields: [
          {
            Field: 'f3',
            Suffix: ' ',
            Italic: true
          },
          {
            Field: 'f2',
            Prefix: '(',
            Suffix: ').',
            Italic: false
          }
        ]
      },
      {
        Name: 'c2',
        Fields: [
          {
            Field: 'f1',
            Suffix: ' - ',
            Italic: false
          },
          {
            Field: 'f2',
            Prefix: '(',
            Suffix: ').',
            Italic: true
          }
        ],
        CiteFields: [
          {
            Field: 'f2',
            Suffix: ' ',
            Italic: true
          },
          {
            Field: 'f3',
            Prefix: '(',
            Suffix: ').',
            Italic: false
          }
        ]
      },
    ] as BibCategory[];

    const e = {
      Key: 'test',
      Category: 'c',
      Fields: [
        'mein',
        'eintrag',
        'test'
      ]
    } as BibEntry

    expect(GenerateCite(e, c)).toEqual('<i>test</i> (eintrag).')
  })
})
