package runtime

import (
	"context"
	"errors"

	"cosmossdk.io/core/store"
	dbm "github.com/cosmos/cosmos-db"

	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type kvStoreService struct {
	key *storetypes.KVStoreKey
}

func (k kvStoreService) OpenKVStore(ctx context.Context) store.KVStore {
	return newKVStore(sdk.UnwrapSDKContext(ctx).KVStore(k.key))
}

type memStoreService struct {
	key *storetypes.MemoryStoreKey
}

func (m memStoreService) OpenMemoryStore(ctx context.Context) store.KVStore {
	return newKVStore(sdk.UnwrapSDKContext(ctx).KVStore(m.key))
}

type transientStoreService struct {
	key *storetypes.TransientStoreKey
}

func (t transientStoreService) OpenTransientStore(ctx context.Context) store.KVStore {
	return newKVStore(sdk.UnwrapSDKContext(ctx).KVStore(t.key))
}

// CoreKVStore is a wrapper of Core/Store kvstore interface
// Remove after https://github.com/cosmos/cosmos-sdk/issues/14714 is closed
type coreKVStore struct {
	kvStore storetypes.KVStore
}

// NewKVStore returns a wrapper of Core/Store kvstore interface
// Remove once store migrates to core/store kvstore interface
func newKVStore(store storetypes.KVStore) store.KVStore {
	return coreKVStore{kvStore: store}
}

// Get returns nil iff key doesn't exist. Errors on nil key.
func (store coreKVStore) Get(key []byte) ([]byte, error) {
	return store.kvStore.Get(key), nil
}

// Has checks if a key exists. Errors on nil key.
func (store coreKVStore) Has(key []byte) (bool, error) {
	return store.kvStore.Has(key), nil
}

// Set sets the key. Errors on nil key or value.
func (store coreKVStore) Set(key, value []byte) error {
	store.kvStore.Set(key, value)
	return nil
}

// Delete deletes the key. Errors on nil key.
func (store coreKVStore) Delete(key []byte) error {
	store.kvStore.Delete(key)
	return nil
}

// Iterator iterates over a domain of keys in ascending order. End is exclusive.
// Start must be less than end, or the Iterator is invalid.
// Iterator must be closed by caller.
// To iterate over entire domain, use store.Iterator(nil, nil)
// CONTRACT: No writes may happen within a domain while an iterator exists over it.
// Exceptionally allowed for cachekv.Store, safe to write in the modules.
func (store coreKVStore) Iterator(start, end []byte) (store.Iterator, error) {
	return &coreIterator{store.kvStore.Iterator(start, end)}, nil
}

// ReverseIterator iterates over a domain of keys in descending order. End is exclusive.
// Start must be less than end, or the Iterator is invalid.
// Iterator must be closed by caller.
// CONTRACT: No writes may happen within a domain while an iterator exists over it.
// Exceptionally allowed for cachekv.Store, safe to write in the modules.
func (store coreKVStore) ReverseIterator(start, end []byte) (store.Iterator, error) {
	return &coreIterator{store.kvStore.ReverseIterator(start, end)}, nil
}

var _ store.Iterator = coreIterator{}

type coreIterator struct {
	iterator dbm.Iterator
}

// Domain implements Iterator.
func (itr *coreIterator) Domain() ([]byte, []byte) {
	return itr.iterator.Domain()
}

// Valid implements Iterator.
func (itr *coreIterator) Valid() bool {
	return itr.iterator.Valid()
}

// Key implements Iterator.
func (itr *coreIterator) Key() ([]byte, error) {
	// Key returns a copy of the current key.
	// See https://github.com/syndtr/goleveldb/blob/52c212e6c196a1404ea59592d3f1c227c9f034b2/leveldb/iterator/iter.go#L88
	if !itr.Valid() {
		return []byte{}, errors.New("iterator is invalid")
	}

	return itr.iterator.Key(), nil
}

// Value implements Iterator.
func (itr *coreIterator) Value() ([]byte, error) {
	// Value returns a copy of the current value.
	// See https://github.com/syndtr/goleveldb/blob/52c212e6c196a1404ea59592d3f1c227c9f034b2/leveldb/iterator/iter.go#L88
	if !itr.Valid() {
		return []byte{}, errors.New("iterator is invalid")
	}

	return itr.iterator.Value(), nil
}

// Next implements Iterator.
func (itr *coreIterator) Next() error {
	if !itr.Valid() {
		return errors.New("iterator is invalid")
	}
	itr.iterator.Next()

	return nil
}

// Error implements Iterator.
func (itr *coreIterator) Error() error {
	return itr.iterator.Error()
}

// Close implements Iterator.
func (itr *coreIterator) Close() error {
	return itr.iterator.Close()
}
