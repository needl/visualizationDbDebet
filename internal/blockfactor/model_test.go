package blockfactor

import (
	"encoding/json"
	"testing"
)

func TestBoolIntScan(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		want      bool
		expectErr bool
	}{
		{name: "nil", input: nil, want: false},
		{name: "int64 one", input: int64(1), want: true},
		{name: "int64 zero", input: int64(0), want: false},
		{name: "int one", input: 1, want: true},
		{name: "bool true", input: true, want: true},
		{name: "unsupported", input: "x", expectErr: true},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var b boolInt
			err := (&b).Scan(tc.input)
			if tc.expectErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if bool(b) != tc.want {
				t.Fatalf("expected %v, got %v", tc.want, bool(b))
			}
		})
	}
}

func TestBoolIntValue(t *testing.T) {
	vTrue := boolInt(true)
	valTrue, err := (&vTrue).Value()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if valTrue != int64(1) {
		t.Fatalf("expected 1, got %v", valTrue)
	}

	vFalse := boolInt(false)
	valFalse, err := (&vFalse).Value()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if valFalse != int64(0) {
		t.Fatalf("expected 0, got %v", valFalse)
	}
}

func TestBoolIntMarshalJSON(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		v := boolInt(true)
		raw, err := (&v).MarshalJSON()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if string(raw) != "true" {
			t.Fatalf("expected true, got %s", string(raw))
		}

		var decoded bool
		if err := json.Unmarshal(raw, &decoded); err != nil {
			t.Fatalf("invalid json: %v", err)
		}
		if !decoded {
			t.Fatal("expected decoded true")
		}
	})

	t.Run("false", func(t *testing.T) {
		v := boolInt(false)
		raw, err := (&v).MarshalJSON()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if string(raw) != "false" {
			t.Fatalf("expected false, got %s", string(raw))
		}
	})
}
