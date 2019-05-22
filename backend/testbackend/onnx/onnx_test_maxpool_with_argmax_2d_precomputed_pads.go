package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("MaxPool", "TestMaxpoolWithArgmax2dPrecomputedPads", NewTestMaxpoolWithArgmax2dPrecomputedPads)
}

// NewTestMaxpoolWithArgmax2dPrecomputedPads version: 3.
func NewTestMaxpoolWithArgmax2dPrecomputedPads() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "MaxPool",
		Title:  "TestMaxpoolWithArgmax2dPrecomputedPads",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xc3, 0x1, 0xa, 0x3c, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x12, 0x1, 0x7a, 0x22, 0x7, 0x4d, 0x61, 0x78, 0x50, 0x6f, 0x6f, 0x6c, 0x2a, 0x15, 0xa, 0xc, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x40, 0x5, 0x40, 0x5, 0xa0, 0x1, 0x7, 0x2a, 0x11, 0xa, 0x4, 0x70, 0x61, 0x64, 0x73, 0x40, 0x2, 0x40, 0x2, 0x40, 0x2, 0x40, 0x2, 0xa0, 0x1, 0x7, 0x12, 0x2c, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x61, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x61, 0x72, 0x67, 0x6d, 0x61, 0x78, 0x5f, 0x32, 0x64, 0x5f, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x64, 0x73, 0x5a, 0x1b, 0xa, 0x1, 0x78, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x5, 0xa, 0x2, 0x8, 0x5, 0x62, 0x1b, 0xa, 0x1, 0x79, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x5, 0xa, 0x2, 0x8, 0x5, 0x62, 0x1b, 0xa, 0x1, 0x7a, 0x12, 0x16, 0xa, 0x14, 0x8, 0x7, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x5, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0x9},

		/*

		   &pb.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y", "z"},
		     Name:      "",
		     OpType:    "MaxPool",
		     Attributes: ([]*pb.AttributeProto) (len=2 cap=2) {
		    (*pb.AttributeProto)(0xc000126100)(name:"kernel_shape" type:INTS ints:5 ints:5 ),
		    (*pb.AttributeProto)(0xc000126300)(name:"pads" type:INTS ints:2 ints:2 ints:2 ints:2 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 1, 5, 5),
				tensor.WithBacking([]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 1, 5, 5),
				tensor.WithBacking([]float32{13, 14, 15, 15, 15, 18, 19, 20, 20, 20, 23, 24, 25, 25, 25, 23, 24, 25, 25, 25, 23, 24, 25, 25, 25}),
			),

			tensor.New(
				tensor.WithShape(1, 1, 5, 5),
				tensor.WithBacking([]int64{12, 13, 14, 14, 14, 17, 18, 19, 19, 19, 22, 23, 24, 24, 24, 22, 23, 24, 24, 24, 22, 23, 24, 24, 24}),
			),
		},
	}
}
