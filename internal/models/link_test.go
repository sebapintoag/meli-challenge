package models

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestGenerateShortUrlKey_Result_NonEmptyString(t *testing.T) {
	var link Link
	link.GenerateShortUrlKey()

	if link.ShortUrl == "" {
		t.Errorf("got \"\"; expected a non-empty string")
	}
}

func TestGenerateShortUrlKey_ExpectedLenght_6(t *testing.T) {
	var link Link
	link.GenerateShortUrlKey()
	keyLength := len(link.ShortUrl)

	if keyLength != 6 {
		t.Errorf("len(link.GenerateShortUrlKey()) = %d; expected 6", keyLength)
	}
}

func TestRemoveUrlPrefix_WithPrefix_RemovesIt(t *testing.T) {
	prefix := os.Getenv("MELI_REDIRECT_URL")
	link := &Link{
		ShortUrl: fmt.Sprintf("%s/XXYYZZ", prefix),
	}
	link.RemoveUrlPrefix()

	if strings.Contains(link.ShortUrl, prefix) {
		t.Errorf("got %s; expected XXYYZZ", link.ShortUrl)
	}
}

func TestRemoveUrlPrefix_WithoutPrefix_SameUrl(t *testing.T) {
	link := &Link{
		ShortUrl: "XXYYZZ",
	}
	link.RemoveUrlPrefix()

	if link.ShortUrl != "XXYYZZ" {
		t.Errorf("got %s; expected XXYYZZ", link.ShortUrl)
	}
}
