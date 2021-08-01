package oop

import "testing"

func TestAB(t *testing.T) {
	t.Run("empty init", func(t *testing.T) {
		var man AB
		ab := "ab"
		man.name = ab
		a := "a"
		man.A.name = a
		b := "b"
		man.B.name = b

		if man.name != ab {
			t.Errorf("name of AB should be '%s' but got '%s'", ab, man.name)
		}

		if man.A.name != a {
			t.Errorf("name of A should be '%s' but got '%s'", a, man.A.name)
		}

		if man.B.name != b {
			t.Errorf("name of B should be '%s' but got '%s'", b, man.B.name)
		}
	})

	t.Run("directly inti", func(t *testing.T) {
		ab := "ab"
		a := "a"
		b := "b"
		var man AB = AB{
			A{name: a},
			B{name: b},
			ab,
		}

		if man.name != ab {
			t.Errorf("name of AB should be '%s' but got '%s'", ab, man.name)
		}

		if man.A.name != a {
			t.Errorf("name of A should be '%s' but got '%s'", a, man.A.name)
		}

		if man.B.name != b {
			t.Errorf("name of B should be '%s' but got '%s'", b, man.B.name)
		}
	})
}
