//go:build !skip_msgp_testing
// +build !skip_msgp_testing

package transactions

// Code generated by github.com/DePINNetwork/msgp DO NOT EDIT.

import (
	"testing"

	"github.com/DePINNetwork/msgp/msgp"

	"github.com/DePINNetwork/depin-sdk/protocol"
	"github.com/DePINNetwork/depin-sdk/test/partitiontest"
)

func TestMarshalUnmarshalApplicationCallTxnFields(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := ApplicationCallTxnFields{}
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

func TestRandomizedEncodingApplicationCallTxnFields(t *testing.T) {
	protocol.RunEncodingTest(t, &ApplicationCallTxnFields{})
}

func BenchmarkMarshalMsgApplicationCallTxnFields(b *testing.B) {
	v := ApplicationCallTxnFields{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgApplicationCallTxnFields(b *testing.B) {
	v := ApplicationCallTxnFields{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalApplicationCallTxnFields(b *testing.B) {
	v := ApplicationCallTxnFields{}
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

func TestMarshalUnmarshalApplyData(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := ApplyData{}
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

func TestRandomizedEncodingApplyData(t *testing.T) {
	protocol.RunEncodingTest(t, &ApplyData{})
}

func BenchmarkMarshalMsgApplyData(b *testing.B) {
	v := ApplyData{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgApplyData(b *testing.B) {
	v := ApplyData{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalApplyData(b *testing.B) {
	v := ApplyData{}
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

func TestMarshalUnmarshalAssetConfigTxnFields(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := AssetConfigTxnFields{}
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

func TestRandomizedEncodingAssetConfigTxnFields(t *testing.T) {
	protocol.RunEncodingTest(t, &AssetConfigTxnFields{})
}

func BenchmarkMarshalMsgAssetConfigTxnFields(b *testing.B) {
	v := AssetConfigTxnFields{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgAssetConfigTxnFields(b *testing.B) {
	v := AssetConfigTxnFields{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalAssetConfigTxnFields(b *testing.B) {
	v := AssetConfigTxnFields{}
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

func TestMarshalUnmarshalAssetFreezeTxnFields(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := AssetFreezeTxnFields{}
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

func TestRandomizedEncodingAssetFreezeTxnFields(t *testing.T) {
	protocol.RunEncodingTest(t, &AssetFreezeTxnFields{})
}

func BenchmarkMarshalMsgAssetFreezeTxnFields(b *testing.B) {
	v := AssetFreezeTxnFields{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgAssetFreezeTxnFields(b *testing.B) {
	v := AssetFreezeTxnFields{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalAssetFreezeTxnFields(b *testing.B) {
	v := AssetFreezeTxnFields{}
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

func TestMarshalUnmarshalAssetTransferTxnFields(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := AssetTransferTxnFields{}
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

func TestRandomizedEncodingAssetTransferTxnFields(t *testing.T) {
	protocol.RunEncodingTest(t, &AssetTransferTxnFields{})
}

func BenchmarkMarshalMsgAssetTransferTxnFields(b *testing.B) {
	v := AssetTransferTxnFields{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgAssetTransferTxnFields(b *testing.B) {
	v := AssetTransferTxnFields{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalAssetTransferTxnFields(b *testing.B) {
	v := AssetTransferTxnFields{}
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

func TestMarshalUnmarshalBoxRef(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := BoxRef{}
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

func TestRandomizedEncodingBoxRef(t *testing.T) {
	protocol.RunEncodingTest(t, &BoxRef{})
}

func BenchmarkMarshalMsgBoxRef(b *testing.B) {
	v := BoxRef{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgBoxRef(b *testing.B) {
	v := BoxRef{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalBoxRef(b *testing.B) {
	v := BoxRef{}
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

func TestMarshalUnmarshalEvalDelta(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := EvalDelta{}
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

func TestRandomizedEncodingEvalDelta(t *testing.T) {
	protocol.RunEncodingTest(t, &EvalDelta{})
}

func BenchmarkMarshalMsgEvalDelta(b *testing.B) {
	v := EvalDelta{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgEvalDelta(b *testing.B) {
	v := EvalDelta{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalEvalDelta(b *testing.B) {
	v := EvalDelta{}
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

func TestMarshalUnmarshalHeader(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := Header{}
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

func TestRandomizedEncodingHeader(t *testing.T) {
	protocol.RunEncodingTest(t, &Header{})
}

func BenchmarkMarshalMsgHeader(b *testing.B) {
	v := Header{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgHeader(b *testing.B) {
	v := Header{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalHeader(b *testing.B) {
	v := Header{}
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

func TestMarshalUnmarshalHeartbeatTxnFields(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := HeartbeatTxnFields{}
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

func TestRandomizedEncodingHeartbeatTxnFields(t *testing.T) {
	protocol.RunEncodingTest(t, &HeartbeatTxnFields{})
}

func BenchmarkMarshalMsgHeartbeatTxnFields(b *testing.B) {
	v := HeartbeatTxnFields{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgHeartbeatTxnFields(b *testing.B) {
	v := HeartbeatTxnFields{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalHeartbeatTxnFields(b *testing.B) {
	v := HeartbeatTxnFields{}
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

func TestMarshalUnmarshalKeyregTxnFields(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := KeyregTxnFields{}
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

func TestRandomizedEncodingKeyregTxnFields(t *testing.T) {
	protocol.RunEncodingTest(t, &KeyregTxnFields{})
}

func BenchmarkMarshalMsgKeyregTxnFields(b *testing.B) {
	v := KeyregTxnFields{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgKeyregTxnFields(b *testing.B) {
	v := KeyregTxnFields{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalKeyregTxnFields(b *testing.B) {
	v := KeyregTxnFields{}
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

func TestMarshalUnmarshalLogicSig(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := LogicSig{}
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

func TestRandomizedEncodingLogicSig(t *testing.T) {
	protocol.RunEncodingTest(t, &LogicSig{})
}

func BenchmarkMarshalMsgLogicSig(b *testing.B) {
	v := LogicSig{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgLogicSig(b *testing.B) {
	v := LogicSig{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalLogicSig(b *testing.B) {
	v := LogicSig{}
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

func TestMarshalUnmarshalPaymentTxnFields(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := PaymentTxnFields{}
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

func TestRandomizedEncodingPaymentTxnFields(t *testing.T) {
	protocol.RunEncodingTest(t, &PaymentTxnFields{})
}

func BenchmarkMarshalMsgPaymentTxnFields(b *testing.B) {
	v := PaymentTxnFields{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgPaymentTxnFields(b *testing.B) {
	v := PaymentTxnFields{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalPaymentTxnFields(b *testing.B) {
	v := PaymentTxnFields{}
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

func TestMarshalUnmarshalPayset(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := Payset{}
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

func TestRandomizedEncodingPayset(t *testing.T) {
	protocol.RunEncodingTest(t, &Payset{})
}

func BenchmarkMarshalMsgPayset(b *testing.B) {
	v := Payset{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgPayset(b *testing.B) {
	v := Payset{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalPayset(b *testing.B) {
	v := Payset{}
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

func TestMarshalUnmarshalSignedTxn(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := SignedTxn{}
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

func TestRandomizedEncodingSignedTxn(t *testing.T) {
	protocol.RunEncodingTest(t, &SignedTxn{})
}

func BenchmarkMarshalMsgSignedTxn(b *testing.B) {
	v := SignedTxn{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgSignedTxn(b *testing.B) {
	v := SignedTxn{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalSignedTxn(b *testing.B) {
	v := SignedTxn{}
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

func TestMarshalUnmarshalSignedTxnInBlock(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := SignedTxnInBlock{}
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

func TestRandomizedEncodingSignedTxnInBlock(t *testing.T) {
	protocol.RunEncodingTest(t, &SignedTxnInBlock{})
}

func BenchmarkMarshalMsgSignedTxnInBlock(b *testing.B) {
	v := SignedTxnInBlock{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgSignedTxnInBlock(b *testing.B) {
	v := SignedTxnInBlock{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalSignedTxnInBlock(b *testing.B) {
	v := SignedTxnInBlock{}
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

func TestMarshalUnmarshalSignedTxnWithAD(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := SignedTxnWithAD{}
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

func TestRandomizedEncodingSignedTxnWithAD(t *testing.T) {
	protocol.RunEncodingTest(t, &SignedTxnWithAD{})
}

func BenchmarkMarshalMsgSignedTxnWithAD(b *testing.B) {
	v := SignedTxnWithAD{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgSignedTxnWithAD(b *testing.B) {
	v := SignedTxnWithAD{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalSignedTxnWithAD(b *testing.B) {
	v := SignedTxnWithAD{}
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

func TestMarshalUnmarshalStateProofTxnFields(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := StateProofTxnFields{}
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

func TestRandomizedEncodingStateProofTxnFields(t *testing.T) {
	protocol.RunEncodingTest(t, &StateProofTxnFields{})
}

func BenchmarkMarshalMsgStateProofTxnFields(b *testing.B) {
	v := StateProofTxnFields{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgStateProofTxnFields(b *testing.B) {
	v := StateProofTxnFields{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalStateProofTxnFields(b *testing.B) {
	v := StateProofTxnFields{}
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

func TestMarshalUnmarshalTransaction(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := Transaction{}
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

func TestRandomizedEncodingTransaction(t *testing.T) {
	protocol.RunEncodingTest(t, &Transaction{})
}

func BenchmarkMarshalMsgTransaction(b *testing.B) {
	v := Transaction{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgTransaction(b *testing.B) {
	v := Transaction{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalTransaction(b *testing.B) {
	v := Transaction{}
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

func TestMarshalUnmarshalTxGroup(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := TxGroup{}
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

func TestRandomizedEncodingTxGroup(t *testing.T) {
	protocol.RunEncodingTest(t, &TxGroup{})
}

func BenchmarkMarshalMsgTxGroup(b *testing.B) {
	v := TxGroup{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgTxGroup(b *testing.B) {
	v := TxGroup{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalTxGroup(b *testing.B) {
	v := TxGroup{}
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
