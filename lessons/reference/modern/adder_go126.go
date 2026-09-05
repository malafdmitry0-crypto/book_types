//go:build go1.26

package modern

type Adder[A Adder[A]] interface { // Go 1.26+
	Add(A) A
}
