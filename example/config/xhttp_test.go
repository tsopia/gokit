package example

import (
	"fmt"
	"github.com/tsopia/gokit/xhttp"
	"testing"
)

func TestNewHttpClient(t *testing.T) {

	var result MyStruct
	err := xhttp.Get("http://example.com1", &result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Field1:", result.Field1)
	fmt.Println("Field2:", result.Field2)

}

type MyStruct struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}
