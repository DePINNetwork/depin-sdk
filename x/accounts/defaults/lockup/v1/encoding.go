package v1

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/cosmos/gogoproto/proto"

	codectypes "github.com/depinnetwork/depin-sdk/codec/types"
)

func UnpackAnyRaw(m *codectypes.Any) (proto.Message, error) {
	split := strings.Split(m.TypeUrl, "/")
	name := split[len(split)-1]
	typ := proto.MessageType(name)
	if typ == nil {
		return nil, fmt.Errorf("no message type found for %s", name)
	}
	concreteMsg := reflect.New(typ.Elem()).Interface().(proto.Message)
	err := proto.Unmarshal(m.Value, concreteMsg)
	if err != nil {
		return nil, err
	}

	return concreteMsg, nil
}
