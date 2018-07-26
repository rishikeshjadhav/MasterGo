// package token  // white box testing
package token_test // black box testing

import (
	"reflect"
	"testing"

	"github.com/rishikeshjadhav/MasterGo/src/16_testing/token"
)

func TestScan(t *testing.T) {
	input := "This is tokenised test"
	expected := []string{"This", "is", "tokenised", "test"}

	actual := token.Scan(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %#v\nActual %#v\n", expected, actual)
	}
}
