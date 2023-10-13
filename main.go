package main

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/kvii/json_null/pb"
	"google.golang.org/protobuf/proto"
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
}

func format(want, got string) {
	fmt.Printf("want: %s\n got: %s\n\n", want, got)
}

func case1() (want, got string) {
	v := &pb.Reply{
		A: proto.Int32(1),
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
