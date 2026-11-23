import assert from "node:assert/strict";
import { test } from "node:test";
import { checkDigit } from "./mrz.mjs";

// Fields of the ICAO Doc 9303 specimen passport line.
const LINE = "L898902C36UTO7408122F1204159ZE184226B<<<<<10";

test("fields of the ICAO specimen", () => {
  assert.equal(checkDigit("L898902C3"), 6); // document number
  assert.equal(checkDigit("740812"), 2); // birth date
  assert.equal(checkDigit("120415"), 9); // expiry date
  assert.equal(checkDigit("ZE184226B<<<<<"), 1); // personal number
});

test("composite of the ICAO specimen", () => {
  assert.equal(checkDigit(LINE.slice(0, 10) + LINE.slice(13, 20) + LINE.slice(21, 43)), 0);
});

test("fillers count as zero", () => {
  assert.equal(checkDigit("<<<<<<"), 0);
});

test("one misread character changes the digit", () => {
  assert.notEqual(checkDigit("L898902C4"), 6);
});

test("refuses a character outside the MRZ alphabet", () => {
  assert.throws(() => checkDigit("l898902c3"));
});
