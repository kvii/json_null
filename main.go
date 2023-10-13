package main

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/kvii/json_null/pb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	// 或者说 2 这种方式就不应该用吗？
	// https://github.com/golang/protobuf/issues/1313#issuecomment-1007692289

	// 1
	format(case1())
	// want: {"a":1}
	//  got: {"a":1}

	// 2
	format(case2())
	// want: {"a":null}
	//  got: {}

	// 3
	format(case3())
	// want: {"a":null}
	//  got: {"a":0}
}

func format(want, got string) {
	fmt.Printf("want: %s\n got: %s\n\n", want, got)
}

func case1() (want, got string) {
	v := &pb.Reply{
		A: wrapperspb.Int32(1),
	}

	bs, _ := json.MarshalOptions.Marshal(v)
	return `{"a":1}`, string(bs)
}

func case2() (want, got string) {
	v := &pb.Reply{
		A: nil,
	}

	bs, _ := json.MarshalOptions.Marshal(v)
	return `{"a":null}`, string(bs)
}

func case3() (want, got string) {
	v := &pb.Reply{
		A: wrapperspb.Int32(0),
	}

	bs, _ := json.MarshalOptions.Marshal(v)
	return `{"a":0}`, string(bs)
}
