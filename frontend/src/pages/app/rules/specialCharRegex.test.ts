import StringContainsSpecialChars from "@/pages/app/rules/specialCharRegex";

describe('StringContainsSpecialChar', () => {
  it('should be true 1', () => {
    expect(StringContainsSpecialChars('a_b')).toBe(true);
  })
  it('should be true 2', () => {
    expect(StringContainsSpecialChars('a!b')).toBe(true);
  })
  it('should be true 3', () => {
    expect(StringContainsSpecialChars('@b')).toBe(true);
  })
  it('should be true 4', () => {
    expect(StringContainsSpecialChars('a#')).toBe(true);
  })
  it('should be true 5', () => {
    expect(StringContainsSpecialChars('a$b')).toBe(true);
  })
  it('should be true 6', () => {
    expect(StringContainsSpecialChars('a(b)')).toBe(true);
  })
  it('should be true 7', () => {
    expect(StringContainsSpecialChars('a/b')).toBe(true);
  })
  it('should be true 8', () => {
    expect(StringContainsSpecialChars('a[]b')).toBe(true);
  })
  it('should be true 9', () => {
    expect(StringContainsSpecialChars('a=b')).toBe(true);
  })
  it('should be true 10', () => {
    expect(StringContainsSpecialChars('a=b')).toBe(true);
  })
  it('should be true 11', () => {
    expect(StringContainsSpecialChars('a:b')).toBe(true);
  })
  it('should be true 12', () => {
    expect(StringContainsSpecialChars('a;b')).toBe(true);
  })
  it('should be true 13', () => {
    expect(StringContainsSpecialChars('a<>b')).toBe(true);
  })
  it('should be true 14', () => {
    expect(StringContainsSpecialChars('a"b')).toBe(true);
  })
  it('should be true 15', () => {
    expect(StringContainsSpecialChars('a?b')).toBe(true);
  })

  it('should be false with a dot', () => {
    expect(StringContainsSpecialChars('a.b')).toBe(false);
  })
})
