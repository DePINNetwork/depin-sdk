package prefix

import (
	"bytes"
	"errors"
	"io"

	"github.com/depinnetwork/depin-sdk/store/cachekv"
	"github.com/depinnetwork/depin-sdk/store/tracekv"
	"github.com/depinnetwork/depin-sdk/store/types"
)

var _ types.KVStore = Store{}

// Store is similar with cometbft/cometbft-db/blob/v1.0.1/prefixdb.go
// both gives access only to the limited subset of the store
// for convenience or safety
type Store struct {
	parent types.KVStore
	prefix []byte
}

func NewStore(parent types.KVStore, prefix []byte) Store {
	return Store{
		parent: parent,
		prefix: prefix,
	}
}

func cloneAppend(bz, tail []byte) (res []byte) {
	res = make([]byte, len(bz)+len(tail))
	copy(res, bz)
	copy(res[len(bz):], tail)
	return
}

func (s Store) key(key []byte) (res []byte) {
	if key == nil {
		panic("nil key on Store")
	}
	res = cloneAppend(s.prefix, key)
	return
}

// GetStoreType implements Store
func (s Store) GetStoreType() types.StoreType {
	return s.parent.GetStoreType()
}

// CacheWrap implements CacheWrap
func (s Store) CacheWrap() types.CacheWrap {
	return cachekv.NewStore(s)
}

// CacheWrapWithTrace implements the KVStore interface.
func (s Store) CacheWrapWithTrace(w io.Writer, tc types.TraceContext) types.CacheWrap {
	return cachekv.NewStore(tracekv.NewStore(s, w, tc))
}

// Get implements KVStore
func (s Store) Get(key []byte) []byte {
	res := s.parent.Get(s.key(key))
	return res
}

// Has implements KVStore
func (s Store) Has(key []byte) bool {
	return s.parent.Has(s.key(key))
}

// Set implements KVStore
func (s Store) Set(key, value []byte) {
	types.AssertValidKey(key)
	types.AssertValidValue(value)
	s.parent.Set(s.key(key), value)
}

// Delete implements KVStore
func (s Store) Delete(key []byte) {
	s.parent.Delete(s.key(key))
}

// Iterator implements KVStore
// Check https://github.com/cometbft/cometbft-db/blob/v1.0.1/prefixdb.go#L109
func (s Store) Iterator(start, end []byte) types.Iterator {
	newstart := cloneAppend(s.prefix, start)

	var newend []byte
	if end == nil {
		newend = cpIncr(s.prefix)
	} else {
		newend = cloneAppend(s.prefix, end)
	}

	iter := s.parent.Iterator(newstart, newend)

	return newPrefixIterator(s.prefix, start, end, iter)
}

// ReverseIterator implements KVStore
// Check https://github.com/cometbft/cometbft-db/blob/v1.0.1/prefixdb.go#L132
func (s Store) ReverseIterator(start, end []byte) types.Iterator {
	newstart := cloneAppend(s.prefix, start)

	var newend []byte
	if end == nil {
		newend = cpIncr(s.prefix)
	} else {
		newend = cloneAppend(s.prefix, end)
	}

	iter := s.parent.ReverseIterator(newstart, newend)

	return newPrefixIterator(s.prefix, start, end, iter)
}

var _ types.Iterator = (*prefixIterator)(nil)

type prefixIterator struct {
	prefix []byte
	start  []byte
	end    []byte
	iter   types.Iterator
	valid  bool
}

func newPrefixIterator(prefix, start, end []byte, parent types.Iterator) *prefixIterator {
	return &prefixIterator{
		prefix: prefix,
		start:  start,
		end:    end,
		iter:   parent,
		valid:  parent.Valid() && bytes.HasPrefix(parent.Key(), prefix),
	}
}

// Domain implements Iterator
func (pi *prefixIterator) Domain() ([]byte, []byte) {
	return pi.start, pi.end
}

// Valid implements Iterator
func (pi *prefixIterator) Valid() bool {
	return pi.valid && pi.iter.Valid()
}

// Next implements Iterator
func (pi *prefixIterator) Next() {
	if !pi.valid {
		panic("prefixIterator invalid, cannot call Next()")
	}

	if pi.iter.Next(); !pi.iter.Valid() || !bytes.HasPrefix(pi.iter.Key(), pi.prefix) {
		// TODO: shouldn't pi be set to nil instead?
		pi.valid = false
	}
}

// Key implements Iterator
func (pi *prefixIterator) Key() (key []byte) {
	if !pi.valid {
		panic("prefixIterator invalid, cannot call Key()")
	}

	key = pi.iter.Key()
	key = stripPrefix(key, pi.prefix)

	return
}

// Value implements Iterator
func (pi *prefixIterator) Value() []byte {
	if !pi.valid {
		panic("prefixIterator invalid, cannot call Value()")
	}

	return pi.iter.Value()
}

// Close implements Iterator
func (pi *prefixIterator) Close() error {
	return pi.iter.Close()
}

// Error returns an error if the prefixIterator is invalid defined by the Valid
// method.
func (pi *prefixIterator) Error() error {
	if !pi.Valid() {
		return errors.New("invalid prefixIterator")
	}

	return nil
}

// copied from github.com/cometbft/cometbft-db/blob/v1.0.1/prefixdb.go
func stripPrefix(key, prefix []byte) []byte {
	if len(key) < len(prefix) || !bytes.Equal(key[:len(prefix)], prefix) {
		panic("should not happen")
	}

	return key[len(prefix):]
}

// wrapping types.PrefixEndBytes
func cpIncr(bz []byte) []byte {
	return types.PrefixEndBytes(bz)
}
