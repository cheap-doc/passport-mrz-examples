# Passport MRZ examples

Small, runnable programs that read a passport or ID card photo with the [doc.cheap API](https://doc.cheap/docs), one folder per language. Each folder holds:

- a program that sends one image (its path given as an argument) with the public sandbox key `sk_sandbox_public` and prints the main fields;
- a dependency-free MRZ check-digit function – ICAO Doc 9303, weights 7-3-1 – with a small unit test;
- a README with the run and test commands.

| Folder | Language | Run | Test |
|---|---|---|---|
| [python](python) | Python 3.8+ | `python3 scan.py image.jpg` | `python3 -m unittest test_mrz.py` |
| [nodejs](nodejs) | Node.js 18+ | `node scan.mjs image.jpg` | `node --test mrz.test.mjs` |
| [csharp](csharp) | C# / .NET 8 | `dotnet run --project PassportScan -- image.jpg` | `dotnet test PassportScan.Tests` |
| [php](php) | PHP 8 | `php scan.php image.jpg` | `php mrz_test.php` |
| [java](java) | Java 17+ | `mvn -q compile exec:java -Dexec.args="image.jpg"` | `mvn -q test` |
| [go](go) | Go 1.21+ | `go run . image.jpg` | `go test ./...` |
| [ruby](ruby) | Ruby 2.6+ | `ruby scan.rb image.jpg` | `ruby mrz_test.rb` |

## The sandbox key

`sk_sandbox_public` is published in the docs and needs no account. It runs the same recognition as a live key and charges nothing: 10 recognised documents per IP address in total, and at most 10 requests an hour. Use a specimen page or another test image, never a real person's document: the public sandbox key keeps a record of every scan for one year ([details](https://doc.cheap/docs/concepts/data-retention-and-privacy#the-public-sandbox-key-keeps-every-scan-for-a-year)).

A free account adds 100 free recognitions every month. Set `DOC_CHEAP_API_KEY` to your own key and every program uses it instead; nothing else changes.

## The request

Every program makes the same call:

```http
POST https://api.doc.cheap/v1/scans
Authorization: Bearer sk_sandbox_public
Content-Type: application/json

{ "image": "<the JPEG or PNG file, base64-encoded>" }
```

The answer comes back in the same response. A photo that holds no readable document is still HTTP 200, with `meta.status` saying what happened, so the programs check `meta.status` before reading `document` and `holder`. An error is a 4xx or 5xx whose body carries a stable `error.code`. The full response is described in the [response reference](https://doc.cheap/docs/reference/response).

## Check digits without an API

The check-digit functions need no network. To check a whole MRZ in the browser, paste it into the free [MRZ checker](https://doc.cheap/mrz-checker).

## Licence

[MIT](LICENSE).
