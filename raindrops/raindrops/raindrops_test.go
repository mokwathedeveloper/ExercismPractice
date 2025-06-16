package raindrops

import "testing"

func TestConvert(t *testing.T) {
    tests := []struct {
        input    int
        expected string
    }{
        {28, "Plong"},
        {30, "PlingPlang"},
        {34, "34"},
        {105, "PlingPlangPlong"},
        {3, "Pling"},
        {5, "Plang"},
        {7, "Plong"},
        {1, "1"},
        {0, "PlingPlangPlong"},
        {-21, "PlingPlong"},
        {-34, "-34"},
    }

    for _, test := range tests {
        got := Convert(test.input)
        if got != test.expected {
            t.Errorf("Convert(%d) = %q, want %q", test.input, got, test.expected)
        }
    }
}

func BenchmarkConvert(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Convert(i)
    }
}