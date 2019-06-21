package otus_lesson5

import (
	"strconv"
	"testing"
)

func TestNewList(t *testing.T) {
	testSuite := createTestSuite()
	list := NewList()

	//Test init state
	if list == nil || list.Len() != 0 {
		t.Fatal("List is nil or already contains items")
	}

	var count int = 0

	//Add objects from test suite to list & check that list items count is grown
	for _, v := range testSuite {
		list.PushBack(v)
		count++

		if list.Len() != count {
			t.Fatalf("Unexpected list length: expected = %d, actual = %d", count, list.Len())
		}
	}

	//Check first & last items in list and test suite
	if testSuite[0] != list.First().Value() {
		t.Fatalf("Illegal first list element: expected = %v, actual = %v", testSuite[0], list.First().Value())
	}

	lastTestElement := testSuite[len(testSuite) - 1]

	if lastTestElement != list.Last().Value() {
		t.Fatalf("Illegal last list element: expected = %v, actual = %v", lastTestElement, list.Last().Value())
	}

	//Check order of elements in list & test suite starting from index = 1 (first element already checked)
	element := list.First()
	count = 1

	for element = element.Next(); element != nil; element = element.Next() {
		if testSuite[count] != element.Value() {
			t.Fatalf("Illegal list element: expected = %v, actual = %v", testSuite[count], element.Value())
		}
		count++
	}

	//Check that we already scan all elements
	if count != len(testSuite) {
		t.Fatalf("Not all elements founded")
	}

	//Test delete first
	element = list.First()
	element.Remove()

	if element == list.First() || len(testSuite) - 1 != list.Len() {
		t.Fatalf("Element not deleted. Removed value: %v, first element = %v. Expected length = %d, actual = %d",
			element.Value(), list.First().Value(), len(testSuite) - 1, list.Len())
	}

	//Test delete after first element
	element = list.First().Next()
	rightElement := element.Next()

	if rightElement == list.First() || rightElement.Prev() == list.First() {
		t.Fatalf("Incorrect params. Element is same or already references")
	}

	element.Remove()

	if list.First().Next() != rightElement || rightElement.Prev() != list.First() {
		t.Fatalf("Elements is not linked")
	}
}

func createTestSuite() []interface{} {
	objectsCount := 5
	testObjects := make([]interface{}, 0, objectsCount)

	for i := 0; i < objectsCount; i++ {
		testObjects = append(testObjects, strconv.Itoa(i))
	}

	return testObjects
}
