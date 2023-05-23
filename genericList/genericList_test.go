package genericList

import "testing"

func TestGenericListInsert(t *testing.T) {
	glist := New[string]()

	glist.Insert("bob") // 0
	glist.Insert("foo") // 1
	glist.Insert("bar") // 2

	glist.Remove(1)
	glist.RemoveByValue("bar")

}
