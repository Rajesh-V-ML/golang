package math
import "testing"

func TestAdd(t *testing.T){
	got := Add(3,6)
	want := 10

	if got ! = want {
		t.Errorf("got %q, wanted %q",got,want)
	}
}