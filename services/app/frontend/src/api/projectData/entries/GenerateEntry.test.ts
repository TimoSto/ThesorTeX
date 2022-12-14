import {BibCategory} from "../categories/BibCategory";
import GenerateEntry from "./GenerateEntry";
import {BibEntry} from "./BibEntry";

describe('GenerateEntryExample', () => {
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
        ]
      },
    ] as BibCategory[];

    const e = {
      Key: 'test',
      Category: 'c2',
      Fields: [
        'mein',
        'eintrag'
      ]
    } as BibEntry

    expect(GenerateEntry(e, c)).toEqual('mein - (<i>eintrag</i>).')
  })
})
