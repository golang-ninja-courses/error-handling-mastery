package index

import "testing"

func BenchmarkGetIndexFromFileName(b *testing.B) {
	cases := []string{
		"parsed_page",
		"parsedpage",
		"parsed_page_",
		"parsed_page_100_suffix",
		"parsed_page_-1",
		"parsed_page_0",
		"parsed_page_15.5",
		"parsed_page_1000",
		"absolutely incorrect file name",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := range cases {
			_, _ = GetIndexFromFileName(cases[i])
		}
	}
}
