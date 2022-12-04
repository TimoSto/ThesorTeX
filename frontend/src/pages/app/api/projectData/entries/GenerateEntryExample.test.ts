import {BibCategory} from "@/pages/app/api/projectData/categories/BibCategory";
import GenerateEntry from "@/pages/app/api/projectData/entries/GenerateEntryExample";
import {BibEntry} from "@/pages/app/api/projectData/entries/BibEntry";

describe('GenerateEntryExample', () => {
  it('should give correct', () => {
    const c = [
      {
        Name: 'c',
        Fields: [
          {
            Field: 'f1',
            Suffix: ' ',
            Style: 'italic'
          },
          {
            Field: 'f2',
            Prefix: '(',
            Suffix: ').',
            Style: 'normal'
          }
        ]
      },
      {
        Name: 'c2',
        Fields: [
          {
            Field: 'f1',
            Suffix: ' - ',
            Style: 'normal'
          },
          {
            Field: 'f2',
            Prefix: '(',
            Suffix: ').',
            Style: 'italic'
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
