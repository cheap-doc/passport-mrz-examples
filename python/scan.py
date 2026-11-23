"""Send one document image to the doc.cheap API and print the main fields.

Standard library only: python3 scan.py path/to/image.jpg
"""
import base64
import json
import os
import sys
import urllib.error
import urllib.request

API = "https://api.doc.cheap/v1/scans"
# The public sandbox key from the docs. Set DOC_CHEAP_API_KEY to use your own.
KEY = os.environ.get("DOC_CHEAP_API_KEY", "sk_sandbox_public")


def scan(path):
    with open(path, "rb") as f:
        image = base64.b64encode(f.read()).decode("ascii")
    request = urllib.request.Request(
        API,
        data=json.dumps({"image": image}).encode("utf-8"),
        headers={"Authorization": f"Bearer {KEY}", "Content-Type": "application/json"},
        method="POST",
    )
    try:
        with urllib.request.urlopen(request, timeout=60) as response:
            return json.load(response)
    except urllib.error.HTTPError as err:
        # Every error body carries a stable code; branch on it, not on the status.
        error = json.load(err)["error"]
        sys.exit(f"{err.code} {error['code']}: {error['message']} (request {error['request_id']})")


def main():
    if len(sys.argv) != 2:
        sys.exit("usage: python3 scan.py <image.jpg|image.png>")
    result = scan(sys.argv[1])
    meta = result["meta"]
    document = result["document"] or {}
    holder = result["holder"] or {}
    # A photo that holds no readable document is a 200 too: meta.status says so.
    print("status:     ", meta["status"])
    print("billed:     ", meta["billed"])
    print("kind:       ", document.get("kind"))
    print("country:    ", document.get("country"))
    print("number:     ", document.get("number"))
    print("expiry_date:", document.get("expiry_date"))
    print("surname:    ", holder.get("surname"))
    print("given_names:", holder.get("given_names"))
    print("birth_date: ", holder.get("birth_date"))
    print("mrz:        ", result["mrz"]["status"])


if __name__ == "__main__":
    main()
