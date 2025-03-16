package hello

import "testing"

func TestSayHello(t *testing.T) {
	subTests := []struct{
		items []string 
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items: []string{"Sam"},
			result: "Hello, Sam!",
		},
		{
			items: []string{"Matt", "Anne"},
			result: "Hello, Matt, Anne!",
		},
	}
	
	for _, st  := range subTests{
		if s := Say(st.items); s != st.result{
			t.Errorf("Wanted %s (%v), got %s", st.result,
		st.items, s)
		}
	}
}
