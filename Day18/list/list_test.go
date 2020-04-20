package list

import (
	"testing"
)

type assert testing.T

func TestAddFirst(t *testing.T) {
	t.Run("EmptyList_AddNode", func(t *testing.T) {
		l := &List{}
		l.AddFirst(&Node{Value: 1})

		assertAreEqual(t, 1, l.head.Value)
	})

	t.Run("NonEmptyList_AddNode", func(t *testing.T) {
		l := &List{}
		l.AddFirst(&Node{Value: 2})
		l.AddFirst(&Node{Value: 1})

		assertAreEqual(t, 1, l.head.Value)
		assertAreEqual(t, 2, l.head.next.Value)
	})
}

func TestAddLast(t *testing.T) {
	t.Run("EmptyList_AddNode", func(t *testing.T) {
		l := &List{}
		l.AddLast(&Node{Value: 1})

		assertAreEqual(t, 1, l.head.Value)
	})

	t.Run("NonEmptyList_AddNode", func(t *testing.T) {
		l := &List{}
		l.AddLast(&Node{Value: 1})
		l.AddLast(&Node{Value: 2})

		assertAreEqual(t, 1, l.head.Value)
		assertAreEqual(t, 2, l.head.next.Value)
	})
}

func TestRemoveFirst(t *testing.T) {
	t.Run("EmptyList_ReturnError", func(t *testing.T) {
		l := &List{}
		err := l.RemoveFirst()

		assertError(t, err, EmptyListError)
	})

	t.Run("NonEmptyList_RemoveFirstNode", func(t *testing.T) {
		l := &List{}
		l.AddFirst(&Node{Value: 1})
		l.AddLast(&Node{Value: 2})

		l.RemoveFirst()

		assertAreEqual(t, 2, l.head.Value)
	})
}

func TestRemoveLast(t *testing.T) {
	t.Run("EmptyList_ReturnError", func(t *testing.T) {
		l := &List{}
		err := l.RemoveLast()

		assertError(t, err, EmptyListError)
	})

	t.Run("SingleNode_EmptyList", func(t *testing.T) {
		l := &List{}
		l.AddFirst(&Node{Value: 1})
		l.RemoveLast()
		err := l.RemoveLast()

		assertError(t, err, EmptyListError)
	})

	t.Run("NonEmptyList_RemoveLastNode", func(t *testing.T) {
		l := &List{}
		l.AddFirst(&Node{Value: 1})
		l.AddLast(&Node{Value: 2})

		l.RemoveLast()
		l.RemoveLast()
		err := l.RemoveLast()

		assertError(t, err, EmptyListError)
	})
}

func assertAreEqual(t *testing.T, expected, actual int) {
	t.Helper()

	if expected != actual {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}

func assertError(t *testing.T, expected, actual error) {
	t.Helper()

	if expected == nil {
		t.Fatal("Expected error but was nil")
	}

	if expected != actual {
		t.Errorf("Error message expected to contain %q but was %q", expected, actual)
	}
}
