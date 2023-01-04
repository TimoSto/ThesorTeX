import {BibEntry} from "./BibEntry";
import CheckEntryChanged from "./CheckEntryChanged";

describe("checkEntryChanged", function () {
  it('no changes', () => {
    const initial = {
      Key: 'test',
      Category: 'tc',
      Fields: ['a', 'b', 'cde']
    } as BibEntry
    expect(CheckEntryChanged(initial, 'test', 'tc', ['a', 'b', 'cde'])).toBe(false);
  });
  it('key changed', () => {
    const initial = {
      Key: 'test',
      Category: 'tc',
      Fields: ['a', 'b', 'cde']
    } as BibEntry
    expect(CheckEntryChanged(initial, 'test1', 'tc', ['a', 'b', 'cde'])).toBe(true);
  });
  it('category changed', () => {
    const initial = {
      Key: 'test',
      Category: 'tc',
      Fields: ['a', 'b', 'cde']
    } as BibEntry
    expect(CheckEntryChanged(initial, 'test', 'tce', ['a', 'b', 'cde'])).toBe(true);
  });
  it('field changed', () => {
    const initial = {
      Key: 'test',
      Category: 'tc',
      Fields: ['a', 'b', 'cde']
    } as BibEntry
    expect(CheckEntryChanged(initial, 'test', 'tc', ['a', 'b', 'cd'])).toBe(true);
  });
  it('field added', () => {
    const initial = {
      Key: 'test',
      Category: 'tc',
      Fields: ['a', 'b', 'cde']
    } as BibEntry
    expect(CheckEntryChanged(initial, 'test', 'tc', ['a', 'b', 'cde', 'fg'])).toBe(true);
  });
  it('field removed', () => {
    const initial = {
      Key: 'test',
      Category: 'tc',
      Fields: ['a', 'b', 'cde']
    } as BibEntry
    expect(CheckEntryChanged(initial, 'test', 'tc', ['a', 'b'])).toBe(true);
  });
});
