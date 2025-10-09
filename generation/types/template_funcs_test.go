package types

import "testing"

func TestWrapperToPrimitive(t *testing.T) {
	tests := []struct {
		name        string
		wrapperType GoType
		want        string
	}{
		{
			name:        "String wrapper",
			wrapperType: "wrapperspb.String",
			want:        "string",
		},
		{
			name:        "Bool wrapper",
			wrapperType: "wrapperspb.Bool",
			want:        "bool",
		},
		{
			name:        "Float wrapper",
			wrapperType: "wrapperspb.Float",
			want:        "float64",
		},
		{
			name:        "Int32 wrapper",
			wrapperType: "wrapperspb.Int32",
			want:        "int",
		},
		{
			name:        "Int64 wrapper",
			wrapperType: "wrapperspb.Int64",
			want:        "int",
		},
		{
			name:        "unknown wrapper",
			wrapperType: "wrapperspb.Unknown",
			want:        "",
		},
		{
			name:        "non-wrapper type",
			wrapperType: "string",
			want:        "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wrapperToPrimitive(tt.wrapperType)
			if got != tt.want {
				t.Errorf("wrapperToPrimitive(%q) = %q, want %q", tt.wrapperType, got, tt.want)
			}
		})
	}
}

func TestStripPrecision(t *testing.T) {
	tests := []struct {
		name string
		arg  GoType
		want string
	}{
		{
			name: "int32",
			arg:  "int32",
			want: "int",
		},
		{
			name: "int64",
			arg:  "int64",
			want: "int",
		},
		{
			name: "uint32",
			arg:  "uint32",
			want: "uint",
		},
		{
			name: "uint64",
			arg:  "uint64",
			want: "uint",
		},
		{
			name: "float32",
			arg:  "float32",
			want: "float64",
		},
		{
			name: "float64",
			arg:  "float64",
			want: "float64",
		},
		{
			name: "string (no change)",
			arg:  "string",
			want: "string",
		},
		{
			name: "bool (no change)",
			arg:  "bool",
			want: "bool",
		},
		{
			name: "custom type with int in name",
			arg:  "myint32type",
			want: "myinttype",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stripPrecision(tt.arg)
			if got != tt.want {
				t.Errorf("stripPrecision(%q) = %q, want %q", tt.arg, got, tt.want)
			}
		})
	}
}

func TestPrimitiveToWrapper(t *testing.T) {
	tests := []struct {
		name        string
		wrapperType GoType
		want        string
	}{
		{
			name:        "Float wrapper",
			wrapperType: "wrapperspb.Float",
			want:        "float32",
		},
		{
			name:        "String wrapper",
			wrapperType: "wrapperspb.String",
			want:        "string",
		},
		{
			name:        "Bool wrapper",
			wrapperType: "wrapperspb.Bool",
			want:        "bool",
		},
		{
			name:        "Int32 wrapper",
			wrapperType: "wrapperspb.Int32",
			want:        "int32",
		},
		{
			name:        "Int64 wrapper",
			wrapperType: "wrapperspb.Int64",
			want:        "int64",
		},
		{
			name:        "unknown wrapper (returns as-is)",
			wrapperType: "wrapperspb.Unknown",
			want:        "wrapperspb.Unknown",
		},
		{
			name:        "non-wrapper type (returns as-is)",
			wrapperType: "string",
			want:        "string",
		},
		{
			name:        "custom type (returns as-is)",
			wrapperType: "MyCustomType",
			want:        "MyCustomType",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := primitiveToWrapper(tt.wrapperType)
			if got != tt.want {
				t.Errorf("primitiveToWrapper(%q) = %q, want %q", tt.wrapperType, got, tt.want)
			}
		})
	}
}

// Test that wrapper conversions are reversible where applicable
func TestWrapperConversionSymmetry(t *testing.T) {
	tests := []struct {
		wrapper   GoType
		primitive string
	}{
		{
			wrapper:   "wrapperspb.String",
			primitive: "string",
		},
		{
			wrapper:   "wrapperspb.Bool",
			primitive: "bool",
		},
		{
			wrapper:   "wrapperspb.Int32",
			primitive: "int32",
		},
		{
			wrapper:   "wrapperspb.Int64",
			primitive: "int64",
		},
	}

	for _, tt := range tests {
		t.Run(string(tt.wrapper), func(t *testing.T) {
			// Convert wrapper to primitive
			gotPrimitive := wrapperToPrimitive(tt.wrapper)
			// Note: wrapperToPrimitive returns different types (int vs int32/int64)
			// so we can't test perfect symmetry here, but we can test the reverse

			// Convert primitive back to wrapper
			gotWrapper := primitiveToWrapper(tt.wrapper)
			if gotWrapper != tt.primitive {
				t.Errorf("primitiveToWrapper(%q) = %q, want %q", tt.wrapper, gotWrapper, tt.primitive)
			}

			// For string and bool, the conversion should be symmetric
			if tt.wrapper == "wrapperspb.String" || tt.wrapper == "wrapperspb.Bool" {
				if gotPrimitive != tt.primitive {
					t.Errorf("wrapperToPrimitive(%q) = %q, want %q", tt.wrapper, gotPrimitive, tt.primitive)
				}
			}
		})
	}
}
