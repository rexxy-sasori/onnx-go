package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	//	testbackend.Register("Reshape", "TestReshapeOneDim", NewTestReshapeOneDim)
}

// NewTestReshapeOneDim version: 3.
func NewTestReshapeOneDim() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Reshape",
		Title:  "TestReshapeOneDim",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x81, 0x1, 0xa, 0x20, 0xa, 0x4, 0x64, 0x61, 0x74, 0x61, 0xa, 0x5, 0x73, 0x68, 0x61, 0x70, 0x65, 0x12, 0x8, 0x72, 0x65, 0x73, 0x68, 0x61, 0x70, 0x65, 0x64, 0x22, 0x7, 0x52, 0x65, 0x73, 0x68, 0x61, 0x70, 0x65, 0x12, 0x14, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x68, 0x61, 0x70, 0x65, 0x5f, 0x6f, 0x6e, 0x65, 0x5f, 0x64, 0x69, 0x6d, 0x5a, 0x1a, 0xa, 0x4, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0x5a, 0x13, 0xa, 0x5, 0x73, 0x68, 0x61, 0x70, 0x65, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x1, 0x62, 0x16, 0xa, 0x8, 0x72, 0x65, 0x73, 0x68, 0x61, 0x70, 0x65, 0x64, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x18, 0x42, 0x2, 0x10, 0x9},

		/*

		   &pb.NodeProto{
		     Input:     []string{"data", "shape"},
		     Output:    []string{"reshaped"},
		     Name:      "",
		     OpType:    "Reshape",
		     Attributes: ([]*pb.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 3, 4),
				tensor.WithBacking([]float32{0.5488135, 0.71518934, 0.60276335, 0.5448832, 0.4236548, 0.6458941, 0.4375872, 0.891773, 0.96366274, 0.3834415, 0.79172504, 0.5288949, 0.56804454, 0.92559665, 0.071036056, 0.0871293, 0.020218397, 0.83261985, 0.77815676, 0.87001216, 0.9786183, 0.7991586, 0.46147937, 0.7805292}),
			),

			tensor.New(
				tensor.WithShape(1),
				tensor.WithBacking([]int{24}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(24),
				tensor.WithBacking([]float32{0.5488135, 0.71518934, 0.60276335, 0.5448832, 0.4236548, 0.6458941, 0.4375872, 0.891773, 0.96366274, 0.3834415, 0.79172504, 0.5288949, 0.56804454, 0.92559665, 0.071036056, 0.0871293, 0.020218397, 0.83261985, 0.77815676, 0.87001216, 0.9786183, 0.7991586, 0.46147937, 0.7805292}),
			),
		},
	}
}
