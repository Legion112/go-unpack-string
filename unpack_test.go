package unpack

import "testing"

func TestUnpack(t *testing.T) {
	tests := []struct {
		name    string
		arg   	string
		want    string
		wantErr bool
	}{
		{name: "SimpleStringWithoutUnpacking", arg: "abcd", want: "abcd", wantErr: false},
		{name: "Base", arg: "a4bc2d5e", want: "aaaabccddddde", wantErr: false},
		{name: "ErrorNotCorrectStringIfTwoDigit", arg: "45", want: "", wantErr: true},
		{name: "EscapeDigit", arg: `qw\4\5`, want: `qw45`, wantErr: false},
		{name: "EscapeDigitWithRepeat", arg: `qwe\45`, want: `qwe44444`, wantErr: false},
		{name: "EscapeBackSlash", arg: `qwe\\5`, want: `qwe\\\\\`, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Unpack(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unpack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Unpack() got = %v, want %v", got, tt.want)
			}
		})
	}
}