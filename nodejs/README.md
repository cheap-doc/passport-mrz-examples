# Passport MRZ example – Node.js

Sends one document image to the [doc.cheap API](https://doc.cheap/docs) with the public sandbox key and prints the main fields. Also holds a dependency-free MRZ check-digit function (ICAO Doc 9303, weights 7-3-1) with a unit test.

Requires: Node.js 18 or later. No packages.

## Run

```bash
node scan.mjs path/to/image.jpg
```

Use a specimen page or another test image, never a real person's document. The public sandbox key keeps a record of every scan for one year ([details](https://doc.cheap/docs/concepts/data-retention-and-privacy#the-public-sandbox-key-keeps-every-scan-for-a-year)). Set `DOC_CHEAP_API_KEY` to use your own key; nothing else changes.

It prints `status` (`recognized`, `no_document_found`, `unreadable`, `unsupported_document` or `rejected`), `billed`, the document fields (`kind`, `country`, `number`, `expiry_date`), the holder fields (`surname`, `given_names`, `birth_date`) and the MRZ verdict (`passed`, `failed` or `absent`). An image that holds no readable document is an HTTP 200 with a `status` other than `recognized`, so check `status` before reading the fields.

## Test the check digits

```bash
node --test mrz.test.mjs
```

## Files

- `scan.mjs` – the API call, with the built-in `fetch`.
- `mrz.mjs` – `checkDigit(data)`.
- `mrz.test.mjs` – the unit test, with the built-in `node:test`.
