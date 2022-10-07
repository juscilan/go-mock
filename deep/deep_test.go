package deep

import "testing"

func TestExec(t *testing.T) {
	t.Run("Should return a speficic string", func(t *testing.T) {
		d := new(Deep)
		r, _ := d.Exec()
		e := "It was returned from deep level"
		if r != e {
			t.Errorf("Error, expected %s but got %s", e, r)
		}
	})
}
