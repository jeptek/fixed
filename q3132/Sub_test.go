package q3132

import "testing"

func TestSub(t *testing.T) {
	// tested already in Add, just check that y is negated:

	if got, want := Sub(Two, One), One; got != want {
		t.Errorf("Sub(%s,%s): got [%d](%s) want [%d](%s)",
			Two,
			One,
			got,
			got,
			want,
			want,
		)
	}

}
