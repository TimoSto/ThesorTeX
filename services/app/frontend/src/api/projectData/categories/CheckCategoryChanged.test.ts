import {BibCategory} from "./BibCategory";
import CheckCategoryChanged from "./CheckCategoryChanged";

describe('checkCategoryChanged', () => {
  it('no changes', () => {
    const initial = {
      Name: 'test',
      CitaviCategory: 'tt',
      CitaviFilters: [],
      Fields: [{Field: 'test', Italic: false, Prefix: '', Suffix: '', TexValue: false, CitaviAttributes: []}],
      CiteFields: [],
      Example: ''
    } as BibCategory;
    expect(CheckCategoryChanged(initial, 'test', 'tt', [], [{Field: 'test', Italic: false, Prefix: '', Suffix: '', TexValue: false, CitaviAttributes: []}], [])).toBe(false)
  });
  it('name changed', () => {
    const initial = {
      Name: 'test',
      CitaviCategory: 'tt',
      CitaviFilters: [],
      Fields: [],
      CiteFields: [],
      Example: ''
    } as BibCategory;
    expect(CheckCategoryChanged(initial, 'testi', 'tt', [], [], [])).toBe(true)
  });
  it('citavi category changed', () => {
    const initial = {
      Name: 'test',
      CitaviCategory: 'tt',
      CitaviFilters: [],
      Fields: [],
      CiteFields: [],
      Example: ''
    } as BibCategory;
    expect(CheckCategoryChanged(initial, 'test', 'tti', [], [], [])).toBe(true)
  });
  it('citavi filter changed', () => {
    const initial = {
      Name: 'test',
      CitaviCategory: 'tt',
      CitaviFilters: [],
      Fields: [],
      CiteFields: [],
      Example: ''
    } as BibCategory;
    expect(CheckCategoryChanged(initial, 'testi', 'tt', ['te'], [], [])).toBe(true)
  });
  it('field changed', () => {
    const initial = {
      Name: 'test',
      CitaviCategory: 'tt',
      CitaviFilters: [],
      Fields: [{Field: 'test', Italic: false, Prefix: '', Suffix: '', TexValue: false, CitaviAttributes: []}],
      CiteFields: [],
      Example: ''
    } as BibCategory;
    expect(CheckCategoryChanged(initial, 'testi', 'tt', ['te'], [{Field: 'test1', Italic: false, Prefix: '', Suffix: '', TexValue: false, CitaviAttributes: []}], [])).toBe(true)
  });
  it('field added', () => {
    const initial = {
      Name: 'test',
      CitaviCategory: 'tt',
      CitaviFilters: [],
      Fields: [],
      CiteFields: [],
      Example: ''
    } as BibCategory;
    expect(CheckCategoryChanged(initial, 'testi', 'tt', ['te'], [{Field: 'test1', Italic: false, Prefix: '', Suffix: '', TexValue: false, CitaviAttributes: []}], [])).toBe(true)
  });
  it('field removed', () => {
    const initial = {
      Name: 'test',
      CitaviCategory: 'tt',
      CitaviFilters: [],
      Fields: [{Field: 'test', Italic: false, Prefix: '', Suffix: '', TexValue: false, CitaviAttributes: []}],
      CiteFields: [],
      Example: ''
    } as BibCategory;
    expect(CheckCategoryChanged(initial, 'testi', 'tt', ['te'], [], [])).toBe(true)
  });
})
