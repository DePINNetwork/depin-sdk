package simulation

import (
	"bytes"
	"fmt"

	"cosmossdk.io/collections"
	collcodec "cosmossdk.io/collections/codec"

	"github.com/depinnetwork/depin-sdk/types/kv"
)

// NewStoreDecoderFuncFromCollectionsSchema returns a function that decodes two kv pairs when the module fully uses collections
func NewStoreDecoderFuncFromCollectionsSchema(schema collections.Schema) func(kvA, kvB kv.Pair) string {
	colls := schema.ListCollections()
	prefixes := make([][]byte, len(colls))
	valueCodecs := make([]collcodec.UntypedValueCodec, len(colls))
	for i, coll := range colls {
		prefixes[i] = coll.GetPrefix()
		valueCodecs[i] = coll.ValueCodec()
	}

	return func(kvA, kvB kv.Pair) string {
		for i, prefix := range prefixes {
			if bytes.HasPrefix(kvA.Key, prefix) {
				if !bytes.HasPrefix(kvB.Key, prefix) {
					panic(fmt.Sprintf("prefix mismatch, keyA has prefix %x (%s), but keyB does not %x (%s)", prefix, prefix, kvB.Key, kvB.Key))
				}
				vc := valueCodecs[i]
				// unmarshal kvA.Value to the corresponding type
				vA, err := vc.Decode(kvA.Value)
				if err != nil {
					panic(err)
				}
				// unmarshal kvB.Value to the corresponding type
				vB, err := vc.Decode(kvB.Value)
				if err != nil {
					panic(err)
				}
				vAString, vBString := "<nil>", "<nil>"
				if vA != nil {
					if vAString, err = vc.Stringify(vA); err != nil {
						panic(err)
					}
				}
				if vB != nil {
					if vBString, err = vc.Stringify(vB); err != nil {
						panic(err)
					}
				}
				return vAString + "\n" + vBString
			}
		}
		panic(fmt.Errorf("unexpected key %X (%s)", kvA.Key, kvA.Key))
	}
}
