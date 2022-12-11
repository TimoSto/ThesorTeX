import {BibCategory} from "@/pages/app/api/projectData/categories/BibCategory";
import CheckCategoryChanged from "@/pages/app/api/projectData/categories/CheckCategoryChanged";

describe('checkCategoryChanged', () => {
  it('no changes', () => {
    const initial = {
      Name: 'test',
      CitaviType: 'tt',
      CitaviNecessaryFields: [],
      Fields: [],
      CiteFields: [],
      Example: ''
    } as BibCategory;
    expect(CheckCategoryChanged(initial, 'test', 'tt', [])).toBe(false)
  });
  it('name changed', () => {
    const initial = {
      Name: 'test',
      CitaviType: 'tt',
      CitaviNecessaryFields: [],
      Fields: [],
      CiteFields: [],
      Example: ''
    } as BibCategory;
    expect(CheckCategoryChanged(initial, 'testi', 'tt', [])).toBe(true)
  });
  it('citavi category changed', () => {
    const initial = {
      Name: 'test',
      CitaviType: 'tt',
      CitaviNecessaryFields: [],
      Fields: [],
      CiteFields: [],
      Example: ''
    } as BibCategory;
    expect(CheckCategoryChanged(initial, 'test', 'tti', [])).toBe(true)
  });
  it('citavi filter changed', () => {
    const initial = {
      Name: 'test',
      CitaviType: 'tt',
      CitaviNecessaryFields: [],
      Fields: [],
      CiteFields: [],
      Example: ''
    } as BibCategory;
    expect(CheckCategoryChanged(initial, 'testi', 'tt', ['te'])).toBe(true)
  });
})
