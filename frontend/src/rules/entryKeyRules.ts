
export const existingKeyError = 'Rules.KeyAlreadyExists';
export default function GetEntryKeyRules(existing: string[], initialKey: string): ((key: string) => boolean|string)[] {
  const fn = (key: string) => {
    const i = existing.indexOf(key)
    if( i >= 0 && existing[i] != initialKey ) {
      return existingKeyError
    }
    return true
  }

  return [fn]
}
