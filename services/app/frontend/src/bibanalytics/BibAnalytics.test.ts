import AnalyseBibFile, {
    AttributeValue, CreateEntry,
    extractEntryAttributes,
    extractKey,
    findBibliographyCategory,
    getTypeOfEntry,
    separateEntries
} from "./BibAnalytics";
import {BibEntry} from "../api/projectData/entries/BibEntry";
import {BibCategory, Field} from "../api/projectData/categories/BibCategory";


describe('analyseFile', () => {

    describe('separateEntries', () => {
        it('should give one', () => {
            const file = `
@book{
field1="hallo",
field2="@test"
}`;
            expect(separateEntries(file)).toHaveLength(1);
        })
        it('should give two', () => {
            const file = `
@book{
field1="hallo",
field2="@test"
}
@wer{
field1="werlo",
field2="werest"
}`;
            expect(separateEntries(file)).toHaveLength(2);
        })
        it('should give two', () => {
            const file = `
@book{
    field1="hallo",
    field2="@test"
}

@wer{
field1="werlo",
field2="werest"
}`;
            expect(separateEntries(file)).toHaveLength(2);
        })
    });

    describe('getTypeOfEntry', () => {
        it('should give book', () => {
            expect(getTypeOfEntry('@book{\n\tt="hallo\n"}')).toEqual('book')
        })
        it('should give book', () => {
            expect(getTypeOfEntry('book{\n\tt="hallo\n"}')).toEqual('book')
        })
    });

    describe('extractKey', () => {
        it('should give test', () => {
            expect(extractKey('@book{test,a1="rotoe"}', [])).toEqual('test')
        })
        it('should give test', () => {
            expect(extractKey('@book{test,\na1="rotoe"}', [])).toEqual('test')
        })
        it('should give test', () => {
            expect(extractKey('@book{\ttest,a1="rotoe"}', [])).toEqual('test')
        })
        it('should give test-1', () => {
            expect(extractKey('@book{test,a1="rotoe"}', [{Key: 'test'} as BibEntry])).toEqual('test-1')
        })
    })
    describe('extracting attributes and values', () => {
        it('should work for unescaped', () => {
            expect(extractEntryAttributes('{herbert,test="hallo",rudi="sommer"}')).toEqual([
                {
                    Attribute: 'test',
                    Value: 'hallo'
                },
                {
                    Attribute: 'rudi',
                    Value: 'sommer'
                }
            ])
        })
        it('should work for escaped', () => {
            expect(extractEntryAttributes('{herbert,test={"hallo"},rudi="sommer"}')).toEqual([
                {
                    Attribute: 'test',
                    Value: 'hallo'
                },
                {
                    Attribute: 'rudi',
                    Value: 'sommer'
                }
            ])
        })
    })
    describe('findBibCategory', () => {
        it('should give book', () => {
            const aTypes = [
                {
                    Name: 'random1',
                    CitaviFilters: [],
                    CitaviCategory: ''
                },
                {
                    Name: 'buch',
                    CitaviCategory: 'book',
                    CitaviFilters: []
                },
                {
                    Name: 'random2',
                    CitaviFilters: ['doi'],
                    CitaviCategory: ''
                },
            ];
            expect(findBibliographyCategory('book', [{Attribute: 'doi', Value: '/teste/123'}], aTypes as unknown as BibCategory[]).Name).toEqual('buch')
        })
        it('should give buchDoi', () => {
            const aTypes = [
                {
                    Name: 'random1',
                    CitaviFilters: [],
                    CitaviCategory: ''
                },
                {
                    Name: 'buch',
                    CitaviCategory: 'book',
                    CitaviFilters: []
                },
                {
                    Name: 'buchDoi',
                    CitaviCategory: 'book',
                    CitaviFilters: ['doi']
                },
                {
                    Name: 'random2',
                    CitaviFilters: ['doi'],
                    CitaviCategory: ''
                },
            ];
            expect(findBibliographyCategory('book', [{Attribute: 'doi', Value: '/teste/123'}], aTypes as unknown as BibCategory[]).Name).toEqual('buchDoi')
        })
        it('should give buchDoi', () => {
            const aTypes = [
                {
                    Name: 'random1',
                    CitaviFilters: [],
                    CitaviCategory: ''
                },
                {
                    Name: 'buchDoi',
                    CitaviCategory: 'book',
                    CitaviFilters: ['doi']
                },
                {
                    Name: 'buch',
                    CitaviCategory: 'book',
                    CitaviFilters: []
                },
                {
                    Name: 'random2',
                    CitaviFilters: ['doi'],
                    CitaviCategory: ''
                },
            ];
            expect(findBibliographyCategory('book', [{Attribute: 'doi', Value: '/teste/123'}], aTypes as unknown as BibCategory[]).Name).toEqual('buchDoi')
        })
    });
    describe('createEntry', () => {
        it('should give correct with only bibfields', () => {
            const t: BibCategory = {
                Name: 't1',
                CitaviCategory: 'book',
                CitaviFilters: [],
                Fields: [
                    {
                        Field: 'f1',
                        CitaviAttributes: ['cf1'],
                    } as Field,
                    {
                        Field: 'f2',
                        CitaviAttributes: ['cf2'],
                    } as Field,
                    {
                        Field: 'f3',
                        CitaviAttributes: ['cf3'],
                    } as Field,
                ],
                CiteFields: [],
                Example: '',
            };
            const attributes: AttributeValue[] = [
                {
                    Attribute: 'cf2',
                    Value: 'test 1'
                },
                {
                    Attribute: 'cf1',
                    Value: 'test 2'
                },
                {
                    Attribute: 'cf3',
                    Value: 'test 3'
                }
            ];
            const entry = CreateEntry(t, attributes, 'testkey');
            expect(entry.Key).toEqual('testkey');
            expect(entry.Category).toEqual('t1');
            expect(entry.Fields).toEqual(['test 2', 'test 1', 'test 3'])
        })
        it('should give correct with additional citefields', () => {
            const t: BibCategory = {
                Name: 't1',
                CitaviCategory: 'book',
                CitaviFilters: [],
                Fields: [
                    {
                        Field: 'f1',
                        CitaviAttributes: ['cf1'],
                    } as Field,
                    {
                        Field: 'f2',
                        CitaviAttributes: ['cf2'],
                    } as Field,
                    {
                        Field: 'f3',
                        CitaviAttributes: ['cf3'],
                    } as Field,
                ],
                CiteFields: [
                    {
                        Field: 'f4',
                        CitaviAttributes: ['cf4'],
                    } as Field,
                ],
                Example: ''
            };
            const attributes: AttributeValue[] = [
                {
                    Attribute: 'cf2',
                    Value: 'test 1'
                },
                {
                    Attribute: 'cf4',
                    Value: 'test 4'
                },
                {
                    Attribute: 'cf3',
                    Value: 'test 3'
                }
            ];
            const entry = CreateEntry(t, attributes, 'testkey');
            expect(entry.Key).toEqual('testkey');
            expect(entry.Category).toEqual('t1');
            expect(entry.Fields).toEqual(['f1', 'test 1', 'test 3', 'test 4'])
        })
        it('should give correct with duplicate', () => {
            //e.g. author and publisher and both are set => only use first
            const t: BibCategory = {
                Name: 't1',
                CitaviCategory: 'book',
                CitaviFilters: [],
                Fields: [
                    {
                        Field: 'f1',
                        CitaviAttributes: ['cf1'],
                    } as Field,
                    {
                        Field: 'f2',
                        CitaviAttributes: ['cf2', 'cf3'],
                    } as Field,
                ],
                CiteFields: [],
                Example: ''
            };
            const attributes: AttributeValue[] = [
                {
                    Attribute: 'cf1',
                    Value: 'test 1'
                },
                {
                    Attribute: 'cf2',
                    Value: 'test 2'
                },
                {
                    Attribute: 'cf3',
                    Value: 'test 3'
                }
            ];
            const entry = CreateEntry(t, attributes, 'testkey');
            expect(entry.Key).toEqual('testkey');
            expect(entry.Category).toEqual('t1');
            expect(entry.Fields).toEqual(['test 1', 'test 2'])
        })
    })
    describe('analyseFile', () => {
        it('one entry', async () => {
            const filecontent = `@book{Thiemann2008,
title="Barrierefreiheit",
bookTitle="Benutzerfreundliche Online-Hilfen: Grundlagen und Umsetzung mit MadCap Flare",
year="2008",
publisher="Vieweg+Teubner",
address="Wiesbaden",
pages="143--157",
abstract="Barrierefreies Design bedeutet, dass Menschen mit Behinderungen ein elektronisches Angebot uneingeschr{\\"a}nkt und selbstst{\\"a}ndig nutzen k{\\"o}nnen. Im Sinne der Nutzbarkeit muss die Barrierefreiheit beispielsweise einer Online-Hilfe im Web einen Schritt {\\"u}ber die reine Zug{\\"a}nglichkeit hinausgehen: auch behinderte Nutzer sollen mit ihren F{\\"a}higkeiten und Hilfsmitteln elektronische Angebote nutzen k{\\"o}nnen. Barrierefreiheit ist somit Gebrauchstauglichkeit vor dem Hintergrund einer Behinderung oder Einschr{\\"a}nkung.",
isbn="978-3-8348-9483-0",
doi="10.1007/978-3-8348-9483-0_7",
url="https://doi.org/10.1007/978-3-8348-9483-0_7"
}`
            const types: BibCategory[] = [
                {
                    Name: 'buch',
                    CitaviCategory: 'book',
                    CitaviFilters: [],
                    Fields: [
                        {
                            Field: 't1',
                            CitaviAttributes: ['author', 'publisher']
                        } as Field,
                        {
                            Field: 't2',
                            CitaviAttributes: ['title']
                        } as Field,
                        {
                            Field: 't3',
                            CitaviAttributes: ['t44itle']
                        } as Field
                    ],
                    CiteFields: [],
                    Example: ''
                }
            ]
            const entries = await AnalyseBibFile(filecontent, types);
            expect(entries.Entries).toHaveLength(1);
            expect(entries.Entries[0].Key).toEqual('Thiemann2008')
            expect(entries.Entries[0].Fields).toEqual(['Vieweg+Teubner', 'Barrierefreiheit', 't3'])
        })

        it('unknown type', async () => {
            const filecontent = `@booki{Thiemann2008,
title="Barrierefreiheit",
bookTitle="Benutzerfreundliche Online-Hilfen: Grundlagen und Umsetzung mit MadCap Flare",
year="2008",
publisher="Vieweg+Teubner",
address="Wiesbaden",
pages="143--157",
abstract="Barrierefreies Design bedeutet, dass Menschen mit Behinderungen ein elektronisches Angebot uneingeschr{\\"a}nkt und selbstst{\\"a}ndig nutzen k{\\"o}nnen. Im Sinne der Nutzbarkeit muss die Barrierefreiheit beispielsweise einer Online-Hilfe im Web einen Schritt {\\"u}ber die reine Zug{\\"a}nglichkeit hinausgehen: auch behinderte Nutzer sollen mit ihren F{\\"a}higkeiten und Hilfsmitteln elektronische Angebote nutzen k{\\"o}nnen. Barrierefreiheit ist somit Gebrauchstauglichkeit vor dem Hintergrund einer Behinderung oder Einschr{\\"a}nkung.",
isbn="978-3-8348-9483-0",
doi="10.1007/978-3-8348-9483-0_7",
url="https://doi.org/10.1007/978-3-8348-9483-0_7"
}`
            const types: BibCategory[] = [
                {
                    Name: 'buch',
                    CitaviCategory: 'book',
                    CitaviFilters: [],
                    Fields: [
                        {
                            Field: 't1',
                            CitaviAttributes: ['author', 'publisher']
                        } as Field,
                        {
                            Field: 't2',
                            CitaviAttributes: ['title']
                        } as Field,
                        {
                            Field: 't3',
                            CitaviAttributes: ['t44itle']
                        } as Field
                    ],
                    CiteFields: [],
                    Example: ''
                }
            ]
            const entries = await AnalyseBibFile(filecontent, types);
            expect(entries.Entries).toHaveLength(0);
            expect(entries.Unknowns).toHaveLength(1)
            expect(entries.Unknowns).toEqual([{Key: 'Thiemann2008', Category: 'booki'}])
        })

        it('duplicate key', async () => {
            const filecontent = `@book{Thiemann2008,
title="Barrierefreiheit",
bookTitle="Benutzerfreundliche Online-Hilfen: Grundlagen und Umsetzung mit MadCap Flare",
year="2008",
publisher="Vieweg+Teubner",
address="Wiesbaden",
pages="143--157",
abstract="Barrierefreies Design bedeutet, dass Menschen mit Behinderungen ein elektronisches Angebot uneingeschr{\\"a}nkt und selbstst{\\"a}ndig nutzen k{\\"o}nnen. Im Sinne der Nutzbarkeit muss die Barrierefreiheit beispielsweise einer Online-Hilfe im Web einen Schritt {\\"u}ber die reine Zug{\\"a}nglichkeit hinausgehen: auch behinderte Nutzer sollen mit ihren F{\\"a}higkeiten und Hilfsmitteln elektronische Angebote nutzen k{\\"o}nnen. Barrierefreiheit ist somit Gebrauchstauglichkeit vor dem Hintergrund einer Behinderung oder Einschr{\\"a}nkung.",
isbn="978-3-8348-9483-0",
doi="10.1007/978-3-8348-9483-0_7",
url="https://doi.org/10.1007/978-3-8348-9483-0_7"
}
@book{Thiemann2008,
title="Barrierefreiheit",
bookTitle="Benutzerfreundliche Online-Hilfen: Grundlagen und Umsetzung mit MadCap Flare",
year="2008",
publisher="Vieweg+Teubner",
address="Wiesbaden",
pages="143--157",
abstract="Barrierefreies Design bedeutet, dass Menschen mit Behinderungen ein elektronisches Angebot uneingeschr{\\"a}nkt und selbstst{\\"a}ndig nutzen k{\\"o}nnen. Im Sinne der Nutzbarkeit muss die Barrierefreiheit beispielsweise einer Online-Hilfe im Web einen Schritt {\\"u}ber die reine Zug{\\"a}nglichkeit hinausgehen: auch behinderte Nutzer sollen mit ihren F{\\"a}higkeiten und Hilfsmitteln elektronische Angebote nutzen k{\\"o}nnen. Barrierefreiheit ist somit Gebrauchstauglichkeit vor dem Hintergrund einer Behinderung oder Einschr{\\"a}nkung.",
isbn="978-3-8348-9483-0",
doi="10.1007/978-3-8348-9483-0_7",
url="https://doi.org/10.1007/978-3-8348-9483-0_7"
}
@book{Thiemann2008,
title="Barrierefreiheit",
bookTitle="Benutzerfreundliche Online-Hilfen: Grundlagen und Umsetzung mit MadCap Flare",
year="2008",
publisher="Vieweg+Teubner",
address="Wiesbaden",
pages="143--157",
abstract="Barrierefreies Design bedeutet, dass Menschen mit Behinderungen ein elektronisches Angebot uneingeschr{\\"a}nkt und selbstst{\\"a}ndig nutzen k{\\"o}nnen. Im Sinne der Nutzbarkeit muss die Barrierefreiheit beispielsweise einer Online-Hilfe im Web einen Schritt {\\"u}ber die reine Zug{\\"a}nglichkeit hinausgehen: auch behinderte Nutzer sollen mit ihren F{\\"a}higkeiten und Hilfsmitteln elektronische Angebote nutzen k{\\"o}nnen. Barrierefreiheit ist somit Gebrauchstauglichkeit vor dem Hintergrund einer Behinderung oder Einschr{\\"a}nkung.",
isbn="978-3-8348-9483-0",
doi="10.1007/978-3-8348-9483-0_7",
url="https://doi.org/10.1007/978-3-8348-9483-0_7"
}`
            const types: BibCategory[] = [
                {
                    Name: 'buch',
                    CitaviCategory: 'book',
                    CitaviFilters: [],
                    Fields: [
                        {
                            Field: 't1',
                            CitaviAttributes: ['author', 'publisher']
                        } as Field,
                        {
                            Field: 't2',
                            CitaviAttributes: ['title']
                        } as Field,
                        {
                            Field: 't3',
                            CitaviAttributes: ['t44itle']
                        } as Field
                    ],
                    CiteFields: [],
                    Example: ''
                }
            ]
            const entries = await AnalyseBibFile(filecontent, types);
            expect(entries.Entries[0].Key).toEqual('Thiemann2008')
            expect(entries.Entries[1].Key).toEqual('Thiemann2008-1')
            expect(entries.Entries[2].Key).toEqual('Thiemann2008-2')
        })
    })
})
