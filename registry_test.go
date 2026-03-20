package schemata

import (
	"encoding/json"
	"testing"
)

func TestGetKind1Schema(t *testing.T) {
	raw, ok := Get("kind1Schema")
	if !ok {
		t.Fatal("expected kind1Schema to exist")
	}
	if !json.Valid(raw) {
		t.Fatal("kind1Schema is not valid JSON")
	}
}

func TestGetNoteSchema(t *testing.T) {
	raw, ok := Get("noteSchema")
	if !ok {
		t.Fatal("expected noteSchema to exist")
	}
	if !json.Valid(raw) {
		t.Fatal("noteSchema is not valid JSON")
	}
}

func TestGetNonexistent(t *testing.T) {
	_, ok := Get("nonexistent")
	if ok {
		t.Fatal("expected nonexistent key to return false")
	}
}

func TestKeysCount(t *testing.T) {
	keys := Keys()
	if len(keys) <= 100 {
		t.Fatalf("expected >100 keys, got %d", len(keys))
	}
}
