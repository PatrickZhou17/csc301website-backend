package util

import (
	"fmt"
	"testing"

	"shopping-cart/common/util/assert"
)

type TestTypeSpec struct {
	ID   int    `form:"id" json:"id" valid:"Required"`
	Name string `form:"name"`
}

func TestFieldToMap(t *testing.T) {
	var ts TestTypeSpec
	ts.ID = 1
	toMap := FieldToMap(&ts, "json")

	assert.Assert(toMap["id"] == 1, "toMap failed")
}

type EmbedStruct struct {
	TestTypeSpec
}

func TestStructContainsJsonTag(t *testing.T) {
	var obj EmbedStruct
	if containJsonTag(&obj) {
		fmt.Println("contains json tag")
	} else {
		fmt.Println("not contains json tag")
	}
}

func TestBind(t *testing.T) {
	var data TestTypeSpec
	var req TestTypeSpec
	req.Name = "zhangsan"
	req.ID = 1

	err := Bind(&data, &req)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(data)
}
