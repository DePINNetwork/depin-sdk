//go:build !skip_msgp_testing
// +build !skip_msgp_testing

package merklearray

// Code generated by github.com/DePINNetwork/msgp DO NOT EDIT.

import (
	"testing"

	"github.com/DePINNetwork/msgp/msgp"

	"github.com/DePINNetwork/depin-sdk/protocol"
	"github.com/DePINNetwork/depin-sdk/test/partitiontest"
)

func TestMarshalUnmarshalLayer(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := Layer{}
	bts := v.MarshalMsg(nil)
	left, err := v.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func TestRandomizedEncodingLayer(t *testing.T) {
	protocol.RunEncodingTest(t, &Layer{})
}

func BenchmarkMarshalMsgLayer(b *testing.B) {
	v := Layer{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgLayer(b *testing.B) {
	v := Layer{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalLayer(b *testing.B) {
	v := Layer{}
	bts := v.MarshalMsg(nil)
	b.ReportAllocs()
	b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := v.UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestMarshalUnmarshalProof(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := Proof{}
	bts := v.MarshalMsg(nil)
	left, err := v.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func TestRandomizedEncodingProof(t *testing.T) {
	protocol.RunEncodingTest(t, &Proof{})
}

func BenchmarkMarshalMsgProof(b *testing.B) {
	v := Proof{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgProof(b *testing.B) {
	v := Proof{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalProof(b *testing.B) {
	v := Proof{}
	bts := v.MarshalMsg(nil)
	b.ReportAllocs()
	b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := v.UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestMarshalUnmarshalSingleLeafProof(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := SingleLeafProof{}
	bts := v.MarshalMsg(nil)
	left, err := v.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func TestRandomizedEncodingSingleLeafProof(t *testing.T) {
	protocol.RunEncodingTest(t, &SingleLeafProof{})
}

func BenchmarkMarshalMsgSingleLeafProof(b *testing.B) {
	v := SingleLeafProof{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgSingleLeafProof(b *testing.B) {
	v := SingleLeafProof{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalSingleLeafProof(b *testing.B) {
	v := SingleLeafProof{}
	bts := v.MarshalMsg(nil)
	b.ReportAllocs()
	b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := v.UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestMarshalUnmarshalTree(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := Tree{}
	bts := v.MarshalMsg(nil)
	left, err := v.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func TestRandomizedEncodingTree(t *testing.T) {
	protocol.RunEncodingTest(t, &Tree{})
}

func BenchmarkMarshalMsgTree(b *testing.B) {
	v := Tree{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgTree(b *testing.B) {
	v := Tree{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalTree(b *testing.B) {
	v := Tree{}
	bts := v.MarshalMsg(nil)
	b.ReportAllocs()
	b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := v.UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}
