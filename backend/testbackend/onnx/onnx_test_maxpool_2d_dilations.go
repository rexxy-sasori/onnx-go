package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("MaxPool", "TestMaxpool2dDilations", NewTestMaxpool2dDilations)
}

// NewTestMaxpool2dDilations version: 4.
func NewTestMaxpool2dDilations() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "MaxPool",
		Title:  "TestMaxpool2dDilations",
		ModelB: []byte{0x8, 0x4, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xa3, 0x1, 0xa, 0x4c, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x7, 0x4d, 0x61, 0x78, 0x50, 0x6f, 0x6f, 0x6c, 0x2a, 0x12, 0xa, 0x9, 0x64, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x40, 0x2, 0x40, 0x2, 0xa0, 0x1, 0x7, 0x2a, 0x15, 0xa, 0xc, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x40, 0x2, 0x40, 0x2, 0xa0, 0x1, 0x7, 0x2a, 0x10, 0xa, 0x7, 0x73, 0x74, 0x72, 0x69, 0x64, 0x65, 0x73, 0x40, 0x1, 0x40, 0x1, 0xa0, 0x1, 0x7, 0x12, 0x19, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x32, 0x64, 0x5f, 0x64, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5a, 0x1b, 0xa, 0x1, 0x78, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x4, 0x62, 0x1b, 0xa, 0x1, 0x79, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x2, 0x42, 0x2, 0x10, 0xa},

		/*

		   &pb.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "MaxPool",
		     Attributes: ([]*pb.AttributeProto) (len=3 cap=4) {
		    (*pb.AttributeProto)(0xc00011c100)(name:"dilations" type:INTS ints:2 ints:2 ),
		    (*pb.AttributeProto)(0xc00011c500)(name:"kernel_shape" type:INTS ints:2 ints:2 ),
		    (*pb.AttributeProto)(0xc00011c600)(name:"strides" type:INTS ints:1 ints:1 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 1, 4, 4),
				tensor.WithBacking([]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 1, 2, 2),
				tensor.WithBacking([]float32{11, 12, 15, 16}),
			),
		},
	}
}