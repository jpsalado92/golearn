package my_project

import "testing"

func TestSayHello(t *testing.T) {
	want := "Hello, test!."
	got := Say([]string{"test"})

	if got != want {
		t.Errorf("SayHello() = %q, wanted %q", got, want)
	}
}

func TestMultipleSayHello(t *testing.T) {
	subtests := []struct {
		names []string
		want  string
	}{
		{[]string{"test"}, "Hello, test!."},
		{names: []string{"test1", "test2"}, want: "Hello, test1, test2!."},
		{[]string{"test1", "test2", "test3"}, "Hello, test1, test2, test3!."},
		{[]string{}, "Hello, world!"},
	} 

	for _, tt := range subtests {
		t.Run(tt.want, func(t *testing.T) {
			got := Say(tt.names)
			if got != tt.want {
				t.Errorf("SayHello() = %q, wanted %q", got, tt.want)
			}
		})
	}
}
