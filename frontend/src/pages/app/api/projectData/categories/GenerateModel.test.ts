import {BibCategory} from "@/pages/app/api/projectData/categories/BibCategory";
import GenerateModelForCategory from "@/pages/app/api/projectData/categories/GenerateModel";

describe('GenerateModel', () => {
  it('should work with no styling or prefixes/suffixes', () => {
    const c = {
      Name: 'c',
      Fields: [
        {
          Field: 'f1'
        },
        {
          Field: 'f2'
        }
      ]
    } as BibCategory;

    const model = GenerateModelForCategory(c)
    expect(model).toEqual('f1f2')
  })
  it('should work with no styling but prefixes/suffixes', () => {
    const c = {
      Name: 'c',
      Fields: [
        {
          Field: 'f1',
          Suffix: ' '
        },
        {
          Field: 'f2',
          Prefix: '(',
          Suffix: ').'
        }
      ]
    } as BibCategory;

    const model = GenerateModelForCategory(c)
    expect(model).toEqual('f1 (f2).')
  })
  it('should work with no prefixes/suffixes and styling', () => {
    const c = {
      Name: 'c',
      Fields: [
        {
          Field: 'f1',
          Suffix: ' ',
          Italic: 'italic'
        },
        {
          Field: 'f2',
          Prefix: '(',
          Suffix: ').',
          Italic: 'normal'
        }
      ]
    } as BibCategory;

    const model = GenerateModelForCategory(c)
    expect(model).toEqual('<i>f1</i> (f2).')
  })
  it('should work with styling but no prefixes/suffixes', () => {
    const c = {
      Name: 'c',
      Fields: [
        {
          Field: 'f1'
        },
        {
          Field: 'f2',
          Italic: 'italic'
        }
      ]
    } as BibCategory;

    const model = GenerateModelForCategory(c)
    expect(model).toEqual('f1<i>f2</i>')
  })
})
