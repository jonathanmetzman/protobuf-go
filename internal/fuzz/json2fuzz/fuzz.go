// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package jsonfuzz includes fuzzers for protojson.Marshal and protojson.Unmarshal.
package jsonfuzz

import (
	// "google.golang.org/protobuf/internal/encoding/json"
	// "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/internal/encoding/json"
	"fmt"
	// fuzzpb "google.golang.org/protobuf/internal/testprotos/fuzz"
)

	// m1 := &fuzzpb.Fuzz{}
	// if err := (protojson.UnmarshalOptions{
	// 	AllowPartial: true,
	// }).Unmarshal(data, m1); err != nil {
	// 	return 0
	// }
	// data1, err := protojson.MarshalOptions{
	// 	AllowPartial: true,
	// }.Marshal(m1)
	// if err != nil {
	// 	panic(err)
	// }
	// m2 := &fuzzpb.Fuzz{}
	// if err := (protojson.UnmarshalOptions{
	// 	AllowPartial: true,
	// }).Unmarshal(data1, m2); err != nil {
	// 	return 0
	// }
	// if !proto.Equal(m1, m2) {
	// 	panic("not equal")
	// }
	// return 1


// Fuzz is a fuzzer for proto.Marshal and proto.Unmarshal.
func Fuzz(data []byte) (score int) {
	decoder := json.NewDecoder(data)
	_, err := decoder.Read()
	if err != nil {
		fmt.Println("err")
		return 1
	}
	fmt.Println("good")
	if _, ok := val.Float(64); !ok {
			b.Fatal("not a float")
		}
	return 0
}
