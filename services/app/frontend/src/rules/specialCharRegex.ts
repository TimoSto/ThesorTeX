
const specialCharsRegex = /[`!@#$%^&*()_+\-=\[\]{};':"\\|,<>\/?~]/;
export default function StringContainsSpecialChars(str: string): boolean {
  return specialCharsRegex.test(str);
}
