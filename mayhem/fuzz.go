package fuzz

import "github.com/hashicorp/terraform-json"

func mayhemit(bytes []byte) int {

    var test tfjson.Expression
    test.UnmarshalJSON(bytes)
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}