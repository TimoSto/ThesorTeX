import {BibCategory} from "../categories/BibCategory";
import {BibEntry} from "./BibEntry";
import GenerateCite from "./GenerateCite";

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
