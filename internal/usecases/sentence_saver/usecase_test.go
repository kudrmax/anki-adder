package sentence_saver

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestClearString(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "",
			in:   "some sentence",
			out:  "some sentence",
		},
		{
			name: "",
			in:   "",
			out:  "",
		},
		{
			name: "",
			in:   "\n\n\n",
			out:  "",
		},
		{
			name: "",
			in:   "some\n\n",
			out:  "some",
		},
		{
			name: "",
			in:   "some\nsentence",
			out:  "some sentence",
		},
		{
			name: "",
			in:   "\nsome\n\nsentence",
			out:  "some sentence",
		},
		{
			name: "",
			in:   "  \nsome\n\nsentence\n \n   \n",
			out:  "some sentence",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClearString(tt.in)
			assert.Equal(t, got, tt.out)
		})
	}
}
