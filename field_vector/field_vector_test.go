package fieldVector

import (
	"testing"
	testData "volleyball_go/example_data"
)

// func ExampleNewVolley() {
// 	df := testData.ExampleVolleyDF()
// 	test := NewVolley(df, 3, "*")
// 	fmt.Println(test.Aktion)
// 	fmt.Println(test.VolleyField)
// 	// Output:
// [{1 a9 M -} {4 a6 H #}]
//[{1 *2 -1 0.5} {0.5 *2D -0.75 0.75} {0.5 *2A -1.25 0.75} {0.5 *2C -0.75 0.25} {0.5 *2B -1.25 0.25} {1 2 1 -0.5} {1 *3 0 0.5} {0.5 *3D 0.25 0.75} {0.5 *3A -0.25 0.75} {0.5 *3C 0.25 0.25} {0.5 *3B -0.25 0.25} {1 3 -0 -0.5} {1 *4 1 0.5} {0.5 *4D 1.25 0.75} {0.5 *4A 0.75 0.75} {0.5 *4C 1.25 0.25} {0.5 *4B 0.75 0.25} {1 4 -1 -0.5} {1 *9 -1 1.5} {0.5 *9D -0.75 1.75} {0.5 *9A -1.25 1.75} {0.5 *9C -0.75 1.25} {0.5 *9B -1.25 1.25} {1 9 1 -1.5} {1 *8 0 1.5} {0.5 *8D 0.25 1.75} {0.5 *8A -0.25 1.75} {0.5 *8C 0.25 1.25} {0.5 *8B -0.25 1.25} {1 8 -0 -1.5} {1 *7 1 1.5} {0.5 *7D 1.25 1.75} {0.5 *7A 0.75 1.75} {0.5 *7C 1.25 1.25} {0.5 *7B 0.75 1.25} {1 7 -1 -1.5} {1 *1 -1 2.5} {0.5 *1D -0.75 2.75} {0.5 *1A -1.25 2.75} {0.5 *1C -0.75 2.25} {0.5 *1B -1.25 2.25} {1 1 1 -2.5} {1 *6 0 2.5} {0.5 *6D 0.25 2.75} {0.5 *6A -0.25 2.75} {0.5 *6C 0.25 2.25} {0.5 *6B -0.25 2.25} {1 6 -0 -2.5} {1 *5 1 2.5} {0.5 *5D 1.25 2.75} {0.5 *5A 0.75 2.75} {0.5 *5C 1.25 2.25} {0.5 *5B 0.75 2.25} {1 5 -1 -2.5}]
// }

func TestNewVolley(t *testing.T) {
	df := testData.ExampleVolleyDF()
	test := NewVolley(df, 3, "*")
	if test.FeldSize != 3 {
		t.Errorf("Feldgroesse stimmt nicht. Gewollte groesse: %d, wirkliche groesse: %v", 3, test.FeldSize)
	}
}
