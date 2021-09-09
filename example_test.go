// Test linkstore assuming a CID LinkSystem.
package linkstore

import (
	"context"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/ipfs/go-cid"

	// IPLD Imports:
	carv1 "github.com/ipld/go-car"
	carv2 "github.com/ipld/go-car/v2"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent"
	"github.com/ipld/go-ipld-prime/linking"
	"github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	"github.com/ipld/go-ipld-prime/storage"
	"github.com/ipld/go-ipld-prime/traversal/selector"
	sbuilder "github.com/ipld/go-ipld-prime/traversal/selector/builder"
	"github.com/multiformats/go-multicodec"

	// Note: this dagcbor import populates the multicodec table as a side effect.
	// The CID LinkSystem requires this to function. See:
	// https://github.com/ipld/go-ipld-prime/blob/fc47eb2f400c1ab39da4df91c4e03f04e34e26cb/linkingExamples_test.go#L58
	_ "github.com/ipld/go-ipld-prime/codec/dagcbor"

	. "github.com/warpfork/go-wish"
)

func TestMain(m *testing.M) {
	failedTestCount := m.Run()
	os.Exit(failedTestCount)
}

func TestNewStorage(t *testing.T) {
	t.Run("Writing v1 carfile and wrapping to v2.", func(t *testing.T) {
		var testBlockReadOpener ipld.BlockReadOpener
		var testBlockWriteOpener ipld.BlockWriteOpener
		var testLinkSystem linking.LinkSystem

		// Filename for carfile wirting tests. Will end up begin written as, eg,
		// carfile.v1.car, carfile.v2.car, ie, with version and ".car" appended.
		// We don't use os.CreateTemp here, for consistency, because the carfile
		// v2 API handles file creation.
		fileprefix := "carfile"

		// Note: os.TempDir() doesn't guarantee existence or writability of the
		// returned location. This might cause tests to fail. TODO: Check.
		v1carfilename := path.Join(os.TempDir(), fileprefix+".v1"+".car")
		v2carfilename := path.Join(os.TempDir(), fileprefix+".v2"+".car")

		// Create our store link system, the system under test, and make some
		// assertions about its initialized state.
		storageLinkSystem := NewStorageLinkSystemWithNewStorage(cidlink.DefaultLinkSystem())
		Wish(t, storageLinkSystem.ReadStore, ShouldBeSameTypeAs, readStore(testBlockReadOpener))
		Wish(t, storageLinkSystem.WriteStore, ShouldBeSameTypeAs, writeStore(testBlockWriteOpener))
		Wish(t, storageLinkSystem.LinkSystem, ShouldBeSameTypeAs, testLinkSystem)

		// Build a trivial "hello world" node.
		n := fluent.MustBuildMap(basicnode.Prototype.Map, 1, func(na fluent.MapAssembler) {
			na.AssembleEntry("hello").AssignString("world")
		})

		// Reasonable defaults for our CID LinkSystem (kubelt compatible).
		lp := cidlink.LinkPrototype{cid.Prefix{
			Version:  1,
			Codec:    uint64(multicodec.DagCbor),  //0x71,
			MhType:   uint64(multicodec.Sha2_256), //0x12,
			MhLength: 32,                          // sha2-256 hash has a 32-byte sum.
		}}

		// Store the test node into the link system's embedded memory store.
		root := storageLinkSystem.MustStore(linking.LinkContext{}, lp, n)

		// Create a selective car (that selects everything) from the LinkSystem.
		sc := carv1.NewSelectiveCar(context.Background(),
			storageLinkSystem.ReadStore, // <- special sauce
			[]carv1.Dag{{
				// CID of the root node of the DAG to traverse.
				Root: root.(cidlink.Link).Cid,

				// Traversal convenience selector that gives us "everything".
				Selector: everything(),
			}})

		// Create the file on disk that will store the car representation.
		carfile_v1, err := os.Create(v1carfilename)
		if err != nil {
			fmt.Printf("err opening file: %v\n", err)
			panic(err)
		}

		// Write the selective carfile from the memory store to disk.
		sc.Write(carfile_v1)

		// Cheat to create a carv2 file...
		// A carv2 file is a carv1 file with a prepended header and appended index.
		// We "wrap" the carv1 file we just made to turn it into a carv2 file.
		// This is done for convenience. Currently carv1 is used directly by the
		// `ipfs dag import <file>` command, but carv2 is the future. There are more
		// verbose ways of creating carv2 files we're not going to delve into yet.
		if err := carv2.WrapV1File(v1carfilename, v2carfilename); err != nil {
			fmt.Printf("err wrapping carfile: %v\n", err)
			panic(err)
		}

		// Make a new memory store so we're not loading from cache. Overwrite
		// the link system's storage handlers to point at the new store.
		var fresh_store = storage.Memory{}
		storageLinkSystem.ConfigureStorage(fresh_store)

		read_carfile_v1, err := os.Open(v1carfilename)
		if err != nil {
			fmt.Printf("err opening file: %v\n", err)
			panic(err)
		}

		// Now restore the blocks in the link system from the file.
		_, err = carv1.LoadCar(storageLinkSystem.WriteStore, read_carfile_v1)
		if err != nil {
			fmt.Printf("carfile LoadCar error: %v\n", err)
			panic(err)
		}

		// Load the fluent reflection from the link system.
		node_from_carfile, err := storageLinkSystem.Load(
			ipld.LinkContext{},
			root,
			basicnode.Prototype.Any,
		)
		if err != nil {
			fmt.Printf("linksystem load error: %v\n", err)
			panic(err)
		}

		// De-reference the fluent reflection.
		world_node, err := node_from_carfile.LookupByString("hello")
		if err != nil {
			panic(err)
		}

		world, err := world_node.AsString()
		Wish(t, err, ShouldEqual, nil)
		Wish(t, world, ShouldEqual, "world")

		// Remove the v1 carfile.
		err = os.Remove(v1carfilename)
		Wish(t, err, ShouldEqual, nil)

		// Load carv2
		cr, err := carv2.OpenReader(v2carfilename)
		Wish(t, err, ShouldEqual, nil)
		defer func() {
			if err := cr.Close(); err != nil {
				panic(err)
			}
		}()

		// Do we get back the same root we saved (after type coercsion)?
		roots, err := cr.Roots()
		Wish(t, err, ShouldEqual, nil)
		Wish(t, roots[0], ShouldBeSameTypeAs, root.(cidlink.Link).Cid)
		Wish(t, roots[0], ShouldEqual, root.(cidlink.Link).Cid)

		// Remove the v2 carfile.
		err = os.Remove(v2carfilename)
		Wish(t, err, ShouldEqual, nil)
	})
}

// This block makes a selector that recurses over the fluent reflection of
// our go struct (ipld map) and gives us "everything". It's a rip job from
// https://github.com/ipld/go-ipld-adl-hamt/issues/24, see also:
// https://github.com/ipld/go-ipld-prime/issues/171, and:
// https://github.com/ipld/go-ipld-prime/pull/199/files
func everything() datamodel.Node {
	ssb := sbuilder.NewSelectorSpecBuilder(basicnode.Prototype.Any)
	return ssb.ExploreFields(func(efsb sbuilder.ExploreFieldsSpecBuilder) {
		efsb.Insert("Links", ssb.ExploreIndex(1, ssb.ExploreRecursive(selector.RecursionLimitNone(), ssb.ExploreAll(ssb.ExploreRecursiveEdge()))))
	}).Node()
}
