package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHash(t *testing.T) {

	// var hash string
	testCases := []struct {
		name    string
		text    string
		result  string
		isValid bool
	}{
		{
			name:    "valid",
			text:    "1",
			result:  "c4ca4238a0b923820dcc509a6f75849b",
			isValid: true,
		},
		{
			name:    "invalid",
			text:    "123",
			result:  "00000000000000000000000000000000",
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hash := GetHash(tc.text)
			if tc.isValid {
				assert.Equal(t, tc.result, hash)
			} else {
				assert.NotEqual(t, tc.result, hash)
			}
			assert.Len(t, hash, 32)
		})
	}
}

func TestGenerateRandom(t *testing.T) {

	testCases := []struct {
		name    string
		size    int
		result  []byte
		isValid bool
	}{
		{
			name:    "valid",
			size:    5,
			result:  []byte("12345"),
			isValid: true,
		},
		{
			name:    "invalid",
			size:    10,
			result:  []byte("12"),
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := GenerateRandom(tc.size)
			assert.NoError(t, err)
			if tc.isValid {
				assert.Len(t, tc.result, len(result))
			} else {
				assert.NotEqual(t, tc.result, result)
			}
		})
	}
}

func BenchmarkGetHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetHash("text")
	}
}

func BenchmarkGetShort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetShort("text")
	}
}

func BenchmarkGeneratorUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratorUUID()
	}
}

func BenchmarkGenerateRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateRandom(10)
	}
}

func BenchmarkAddSlash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddSlash("test")
	}
}

func TestAddSlash(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid",
			args: args{s: "slash"},
			want: "slash/",
		},
		{
			name: "valid",
			args: args{s: "slash/"},
			want: "slash/",
		},
		{
			name: "valid",
			args: args{s: ""},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, AddSlash(tt.args.s), "AddSlash(%v)", tt.args.s)
		})
	}
}

func TestGetShort(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid",
			args: args{text: "1"},
			want: "c4ca4238a0b923820dcc509a6f75849b",
		},
		{
			name: "valid",
			args: args{text: "777"},
			want: "f1c1592588411002af340cbaedd6fc33",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetShort(tt.args.text), "GetShort(%v)", tt.args.text)
		})
	}
}

func TestGeneratorUUID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "valid",
			want: GeneratorUUID(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotEqualf(t, tt.want, GeneratorUUID(), "GeneratorUUID()")
		})
	}
}

func TestEncrypted(t *testing.T) {
	type args struct {
		msg []byte
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "string",
			args: args{[]byte("1"), "1"},
			want: []byte{0x36, 0xda, 0xee, 0xb, 0x16, 0x82, 0x42, 0xc4, 0x58, 0xcc, 0xba, 0x3b, 0x1a, 0xcf, 0xc3, 0x4e, 0xde, 0x80, 0xb7, 0xc, 0x1d, 0xda, 0x4, 0xca, 0xbd, 0x17, 0xde, 0x31, 0x4f},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Encrypted(tt.args.msg, tt.args.key)

			assert.NotEqualf(t, tt.want, got, "Encrypted(%v, %v)", tt.args.msg, tt.args.key)
		})
	}
}

func TestDecrypted(t *testing.T) {
	type args struct {
		msg []byte
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		isValid bool
	}{
		{
			name:    "valid sting decode 1",
			args:    args{[]byte{0x36, 0xda, 0xee, 0xb, 0x16, 0x82, 0x42, 0xc4, 0x58, 0xcc, 0xba, 0x3b, 0x1a, 0xcf, 0xc3, 0x4e, 0xde, 0x80, 0xb7, 0xc, 0x1d, 0xda, 0x4, 0xca, 0xbd, 0x17, 0xde, 0x31, 0x4f}, "1"},
			want:    []byte("1"),
			isValid: true,
		},
		{
			name:    "invalid sting decode 1",
			args:    args{[]byte{0x36, 0xda, 0xee, 0xb, 0x16, 0x82, 0x42, 0xc4, 0x58, 0xcc, 0xba, 0x3b, 0x1a, 0xcf, 0xc3, 0x4e, 0xde, 0x80, 0xb7, 0xc, 0x1d, 0xda, 0x4, 0xca, 0xbd, 0x17, 0xde, 0x31, 0x4f}, "2"},
			want:    []byte("1"),
			isValid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Decrypted(tt.args.msg, tt.args.key)
			if tt.isValid {
				assert.Equalf(t, tt.want, got, "Decrypted(%v, %v)", tt.args.msg, tt.args.key)
			} else {
				assert.NotEqualf(t, tt.want, got, "Decrypted(%v, %v)", tt.args.msg, tt.args.key)

			}
		})
	}
}
