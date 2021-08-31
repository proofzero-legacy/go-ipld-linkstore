// Adds IPLD format architecture support to a prime architecture link system.
package linkstore

import (
	"io/ioutil"

	// IPFS
	cid "github.com/ipfs/go-cid"

	// IPLD Format Architecture
	blocks "github.com/ipfs/go-block-format"
	carv1 "github.com/ipld/go-car"

	// IPLD Prime Architecture
	"github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/storage"
)

// A "StorageLinkSystem" is a "LinkSystem" with format architecture storage duct
// taped to it so, eg, it can be written to carfiles.
type StorageLinkSystem struct {
    ipld.LinkSystem
	ReadStore  carv1.ReadStore
	WriteStore carv1.Store
}

// Create a new storage link system with no attached storage.
func NewStorageLinkSystemWithNoStorage() *StorageLinkSystem {
    return &StorageLinkSystem{
        cidlink.DefaultLinkSystem(),
        nil,
        nil,
    }
}

// Creates a new StorageLinkSystem with supplied storage.
func NewStorageLinkSystemWithStorage(store storage.Memory) *StorageLinkSystem {
	return NewStorageLinkSystemWithNoStorage().ConfigureStorage(store)
}

// Creates a new StorageLinkSystem with new storage.
func NewStorageLinkSystemWithNewStorage() *StorageLinkSystem {
	var store = storage.Memory{}
	return NewStorageLinkSystemWithStorage(store)
}

// Configure link system storage for prime architecture nodes and set up the
// duct tape storage for format architecture nodes.
func (slinks *StorageLinkSystem) ConfigureStorage(store storage.Memory) *StorageLinkSystem {
	// Standard link system storage handler setup:
	slinks.LinkSystem.StorageWriteOpener = (&store).OpenWrite
	slinks.LinkSystem.StorageReadOpener = (&store).OpenRead

	// IMPORTANT: These lines make the new prime architecture linksystem store
	// (an IPLD memory store) compatible with the old format architecture
	// block-based store that the v1 carfile API expects.

	// Wrap a "prime architecture" (see below) BlockReadOpener interface in our
	// readStore type so we can extend it with a Get method to make it compatible
	// with the "format architecture" (see below) go-car v1 block-based interface.
	slinks.ReadStore = readStore(slinks.LinkSystem.StorageReadOpener)

	// Wrap a "prime architecture" (see below) BlockWriteOpener interface in our
	// writeStore type so we can extend it with a Put method to make it compatible
	// with the "format architecture" (see below) go-car v1 block-based interface.
	slinks.WriteStore = writeStore(slinks.LinkSystem.StorageWriteOpener)

	// Return self for chaining.
	return slinks
}

////////////////////////////////////////////////////////////////////////////////
//                                 DUCT TAPE
////////////////////////////////////////////////////////////////////////////////
//
// This is based on the following discussion and example for the hash-array
// mapped trie advanced data layout. See:
//
// - https://github.com/ipld/go-ipld-adl-hamt/issues/24
// - https://github.com/ipld/go-ipld-adl-hamt/issues/24#issuecomment-891315942
// - https://clbin.com/iD5rF
//
// These readStore and writeStore types allow us to extend ipld.BlockReadOpener,
// and ipld.BlockWriteOpener, respectively, to "duct tape the interface" between
// the old block format node type and the new IPLD prime nodes.
//
// The "store" architecture is old and "opener" is new. The go-car module
// (specifically for v1 carfiles) uses the older "store" architecture.

type readStore ipld.BlockReadOpener
type writeStore ipld.BlockWriteOpener

// See also:
//
// - https://github.com/ipld/go-ipld-prime  <- new IPLD prime nodes
// - https://github.com/ipfs/go-ipld-format <- old IPLD format nodes
//
// Format nodes have a `.Block` member that gives access to their representation
// as a byte array. Prime nodes do not have this member and instead have their
// byte array representation computed by LinkSystems, so that multicodecs,
// multihashes, etc., can be applied in a consistent way. This is more flexible
// and, in the end, easier to understand, but requires a conceptual change that
// makes prime nodes and format nodes architecturally incompatible. The old
// format architecture, currently and unfortunately, has much better
// documentation and library support than newer, better prime. `go-car` is one
// of these format architecture libraries and we want to use it with prime
// architecture nodes.
//
// Therefore, tl;dr: we wrap the newer prime BlockReadOpener and
// BlockWriteOpener so we can add Get and Put methods to make their prime
// architecture interfaces compatible with the carv1 block-based (ie, format
// architecture) storage APIs.
//
// The "format architecture" block-based storage API requires a Get method that
// returns, for a passed CID, a new block created from its representation as a
// byte array. We attach this method to the (prime) LinkSystem's BlockReadOpener
// so that it can be used in the (format) go-car v1 selective car writer (as the
// selective car writer needs to read/Get blocks for CIDs from memory).

func (rs readStore) Get(c cid.Cid) (blocks.Block, error) {
	// Create a Link from the passed CID (this is what our LinkSystem uses too).
	link := cidlink.Link{Cid: c}

	// Get the io.Reader for the byte array representation of the node.
	r, err := rs(ipld.LinkContext{}, link)
	if err != nil {
		return nil, err
	}

	// Get all the data from the io.Reader.
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	// Make a block using that data and the passed CID. Saves the compute of
	// recalculating the CID for the data we retrieved -- it should be the same.
	return blocks.NewBlockWithCid(data, c)
}

// The "format architecture" block-based storage API requires a Put method that
// takes blocks and writes the data they contain to storage. We attach this
// method to the (prime) LinkSystem's BlockWriteOpener so that it can be used in
// the (format) go-car v1 selective car reader (the selective car reader needs
// to write/Put blocks from the file it is reading into memory).

func (ws writeStore) Put(b blocks.Block) error {
	// Make a Link from the link in the passed block (from our LinkSystem).
	link := cidlink.Link{Cid: b.Cid()}

	// Get the byte array the block contains.
	data := b.RawData()

	// Get the io.Writer and BlockWriteCommitter for writing the blocks to the
	// (in our case, memory) store. From the go-ipld-prime comments for
	// `linking/types.go`:
	//
	// > BlockWriteOpener implementations are expected to start writing their content immediately,
	// > and later, the returned BlockWriteCommitter should also be able to expect that
	// > the Link which it is given is a reasonable hash of the content.
	// > (To give an example of how this might be efficiently implemented:
	// > One might imagine that if implementing a disk storage mechanism,
	// > the io.Writer returned from a BlockWriteOpener will be writing a new tempfile,
	// > and when the BlockWriteCommiter is called, it will flush the writes
	// > and then use a rename operation to place the tempfile in a permanent path based the Link.)
	writer, committer, err := ws(ipld.LinkContext{})
	if err != nil {
		return err
	}

	// Write/commit/return
	writer.Write(data)
	return committer(link)
}

//////////////////////////////////END OF TAPE//////////////////////////////////
