package generics

import "testing"

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		stack := &Stack[int]{}

		AssertTrue(t, stack.IsEmpty())

		stack.Push(42)
		AssertFalse(t, stack.IsEmpty())

		stack.Push(124)
		AssertFalse(t, stack.IsEmpty())

		el, ok := stack.Pop()
		AssertFalse(t, stack.IsEmpty())
		AssertTrue(t, ok)
		AssertEqual(t, el, 124)

		el, ok = stack.Pop()
		AssertTrue(t, stack.IsEmpty())
		AssertTrue(t, ok)
		AssertEqual(t, el, 42)

		_, ok = stack.Pop()
		AssertTrue(t, stack.IsEmpty())
		AssertFalse(t, ok)
	})

	t.Run("string stack", func(t *testing.T) {
		stack := &Stack[string]{}

		AssertTrue(t, stack.IsEmpty())

		stack.Push("hello")
		AssertFalse(t, stack.IsEmpty())

		stack.Push("world")
		AssertFalse(t, stack.IsEmpty())

		el, ok := stack.Pop()
		AssertFalse(t, stack.IsEmpty())
		AssertTrue(t, ok)
		AssertEqual(t, el, "world")

		el, ok = stack.Pop()
		AssertTrue(t, stack.IsEmpty())
		AssertTrue(t, ok)
		AssertEqual(t, el, "hello")

		_, ok = stack.Pop()
		AssertTrue(t, stack.IsEmpty())
		AssertFalse(t, ok)
	})
}

func AssertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("wanted %v not to equal %v", got, want)
	}
}

func AssertTrue(t testing.TB, got bool) {
	t.Helper()

	if !got {
		t.Error("expected true, but got false")
	}
}

func AssertFalse(t testing.TB, got bool) {
	t.Helper()

	if got {
		t.Error("expected false, but got true")
	}
}
