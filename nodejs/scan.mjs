// Send one document image to the doc.cheap API and print the main fields.
// Node.js 18 or later, no packages: node scan.mjs path/to/image.jpg
import { readFile } from "node:fs/promises";

const API = "https://api.doc.cheap/v1/scans";
// The public sandbox key from the docs. Set DOC_CHEAP_API_KEY to use your own.
const KEY = process.env.DOC_CHEAP_API_KEY ?? "sk_sandbox_public";

const path = process.argv[2];
if (!path) {
  console.error("usage: node scan.mjs <image.jpg|image.png>");
  process.exit(2);
}

const image = (await readFile(path)).toString("base64");
const response = await fetch(API, {
  method: "POST",
  headers: { Authorization: `Bearer ${KEY}`, "Content-Type": "application/json" },
  body: JSON.stringify({ image }),
  signal: AbortSignal.timeout(60_000),
});
const result = await response.json();

if (!response.ok) {
  // Every error body carries a stable code; branch on it, not on the status.
  const { code, message, request_id } = result.error;
  console.error(`${response.status} ${code}: ${message} (request ${request_id})`);
  process.exit(1);
}

// A photo that holds no readable document is a 200 too: meta.status says so.
const { meta, mrz } = result;
const document = result.document ?? {};
const holder = result.holder ?? {};
console.log("status:     ", meta.status);
console.log("billed:     ", meta.billed);
console.log("kind:       ", document.kind ?? null);
console.log("country:    ", document.country ?? null);
console.log("number:     ", document.number ?? null);
console.log("expiry_date:", document.expiry_date ?? null);
console.log("surname:    ", holder.surname ?? null);
console.log("given_names:", holder.given_names ?? null);
console.log("birth_date: ", holder.birth_date ?? null);
console.log("mrz:        ", mrz.status);
