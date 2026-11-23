// MRZ check digits as ICAO Doc 9303 defines them. No dependencies.

const WEIGHTS = [7, 3, 1];

/** The value of one MRZ character: 0-9 as themselves, A-Z as 10-35, "<" as 0. */
export function charValue(c) {
  if (c >= "0" && c <= "9") return c.charCodeAt(0) - 48;
  if (c >= "A" && c <= "Z") return c.charCodeAt(0) - 55; // A = 10
  if (c === "<") return 0;
  throw new Error(`not an MRZ character: ${JSON.stringify(c)}`);
}

/** The check digit of a field: each value times 7, 3, 1 repeating, summed, mod 10. */
export function checkDigit(data) {
  let sum = 0;
  for (let i = 0; i < data.length; i++) sum += charValue(data[i]) * WEIGHTS[i % 3];
  return sum % 10;
}
