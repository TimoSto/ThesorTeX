import {BibCategory} from "./BibCategory";
import GenerateModelForFields from "./GenerateModel";

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

    const model = GenerateModelForFields(c.Fields)
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

    const model = GenerateModelForFields(c.Fields)
    expect(model).toEqual('f1 (f2).')
  })
  it('should work with no prefixes/suffixes and styling', () => {
    const c = {
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
    } as BibCategory;

    const model = GenerateModelForFields(c.Fields)
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
          Italic: true
        }
      ]
    } as BibCategory;

    const model = GenerateModelForFields(c.Fields)
    expect(model).toEqual('f1<i>f2</i>')
  })
})
