// Send one document image to the doc.cheap API and print the main fields.
// Standard library only: go run . path/to/image.jpg
package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const api = "https://api.doc.cheap/v1/scans"

type scanResult struct {
	Meta struct {
		Status string `json:"status"`
		Billed bool   `json:"billed"`
	} `json:"meta"`
	Document *struct {
		Kind       string  `json:"kind"`
		Country    *string `json:"country"`
		Number     *string `json:"number"`
		ExpiryDate *string `json:"expiry_date"`
	} `json:"document"`
	Holder *struct {
		Surname    *string `json:"surname"`
		GivenNames *string `json:"given_names"`
		BirthDate  *string `json:"birth_date"`
	} `json:"holder"`
	Mrz struct {
		Status string `json:"status"`
	} `json:"mrz"`
}

type errorBody struct {
	Error struct {
		Code      string `json:"code"`
		Message   string `json:"message"`
		RequestID string `json:"request_id"`
	} `json:"error"`
}

func str(s *string) string {
	if s == nil {
		return "null"
	}
	return *s
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: go run . <image.jpg|image.png>")
		os.Exit(2)
	}
	// The public sandbox key from the docs. Set DOC_CHEAP_API_KEY to use your own.
	key := os.Getenv("DOC_CHEAP_API_KEY")
	if key == "" {
		key = "sk_sandbox_public"
	}

	raw, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	body, _ := json.Marshal(map[string]string{"image": base64.StdEncoding.EncodeToString(raw)})
	req, _ := http.NewRequest(http.MethodPost, api, bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("Content-Type", "application/json")

	resp, err := (&http.Client{Timeout: 60 * time.Second}).Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Every error body carries a stable code; branch on it, not on the status.
		var e errorBody
		_ = json.NewDecoder(resp.Body).Decode(&e)
		fmt.Fprintf(os.Stderr, "%d %s: %s (request %s)\n", resp.StatusCode, e.Error.Code, e.Error.Message, e.Error.RequestID)
		os.Exit(1)
	}

	var r scanResult
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// A photo that holds no readable document is a 200 too: meta.status says so.
	fmt.Println("status:     ", r.Meta.Status)
	fmt.Println("billed:     ", r.Meta.Billed)
	if r.Document != nil {
		fmt.Println("kind:       ", r.Document.Kind)
		fmt.Println("country:    ", str(r.Document.Country))
		fmt.Println("number:     ", str(r.Document.Number))
		fmt.Println("expiry_date:", str(r.Document.ExpiryDate))
	}
	if r.Holder != nil {
		fmt.Println("surname:    ", str(r.Holder.Surname))
		fmt.Println("given_names:", str(r.Holder.GivenNames))
		fmt.Println("birth_date: ", str(r.Holder.BirthDate))
	}
	fmt.Println("mrz:        ", r.Mrz.Status)
}
