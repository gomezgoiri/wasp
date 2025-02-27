package trie

import (
	"encoding/hex"
)

// nodeStore immutable node store
type nodeStore struct {
	trieStore  KVReader
	valueStore KVReader
}

const defaultCacheSize = 10_000

const (
	partitionTrieNodes = byte(iota)
	partitionValues
)

// MustInitRoot initializes a new empty trie
func MustInitRoot(store KVWriter) Hash {
	rootNodeData := newNodeData()
	n := newBufferedNode(rootNodeData, nil)

	trieStore := makeWriterPartition(store, partitionTrieNodes)
	valueStore := makeWriterPartition(store, partitionValues)
	n.commitNode(trieStore, valueStore)

	return n.nodeData.commitment
}

func openNodeStore(store KVReader, cacheSize ...int) *nodeStore {
	size := defaultCacheSize
	if len(cacheSize) > 0 {
		size = cacheSize[0]
	}

	store = makeCachedKVReader(store, size)
	return &nodeStore{
		trieStore:  makeReaderPartition(store, partitionTrieNodes),
		valueStore: makeReaderPartition(store, partitionValues),
	}
}

func (ns *nodeStore) FetchNodeData(nodeCommitment Hash) (*nodeData, bool) {
	dbKey := nodeCommitment.Bytes()
	nodeBin := ns.trieStore.Get(dbKey)
	if len(nodeBin) == 0 {
		return nil, false
	}
	ret, err := nodeDataFromBytes(nodeBin)
	assertf(err == nil, "NodeStore::FetchNodeData err: '%v' nodeBin: '%s', commitment: %s",
		err, hex.EncodeToString(nodeBin), nodeCommitment)
	ret.commitment = nodeCommitment
	return ret, true
}

func (ns *nodeStore) MustFetchNodeData(nodeCommitment Hash) *nodeData {
	ret, ok := ns.FetchNodeData(nodeCommitment)
	assertf(ok, "NodeStore::MustFetchNodeData: cannot find node data: commitment: '%s'", nodeCommitment.String())
	return ret
}

func (ns *nodeStore) FetchChild(n *nodeData, childIdx byte, trieKey []byte) (*nodeData, []byte) {
	c := n.children[childIdx]
	if c == nil {
		return nil, nil
	}
	childTriePath := concat(trieKey, n.pathExtension, []byte{childIdx})

	ret, ok := ns.FetchNodeData(*c)
	assertf(ok, "immutable::FetchChild: failed to fetch node. trieKey: '%s', childIndex: %d",
		hex.EncodeToString(trieKey), childIdx)
	return ret, childTriePath
}
