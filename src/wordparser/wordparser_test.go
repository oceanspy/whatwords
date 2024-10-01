package wordparser

import (
	"reflect"
	"testing"
)

type TestWordList struct {
	WordList []string
	Word     string
	Expected string
}

func TestCalculateOccurence(t *testing.T) {

}

func TestRemoveSliceElement(t *testing.T) {
	sl := []string{"simple", "test1", "test2", "enfin", "voila"}
	i := 2
	expectedSl := []string{"simple", "test2", "enfin", "voila"}

	RemoveSliceElement(&sl, i)

	if reflect.DeepEqual(sl, expectedSl) {
		t.Errorf("Expected sl doesn't match -- got: %s, expected: %s", sl, expectedSl)
	}
}
