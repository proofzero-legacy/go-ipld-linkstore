package linkstore

import (
	"os"
    "fmt"
	"testing"
	"context"

	"github.com/ipfs/go-cid"
	carv1 "github.com/ipld/go-car"
	carv2 "github.com/ipld/go-car/v2"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/fluent"
	"github.com/ipld/go-ipld-prime/linking"
	"github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/node/basicnode"
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
        fileprefix := "carfile"

		storageLinkSystem := NewStorageLinkSystemWithNewStorage(cidlink.DefaultLinkSystem())
		Wish(t, storageLinkSystem.ReadStore, ShouldBeSameTypeAs, readStore(testBlockReadOpener))
		Wish(t, storageLinkSystem.WriteStore, ShouldBeSameTypeAs, writeStore(testBlockWriteOpener))
		Wish(t, storageLinkSystem.LinkSystem, ShouldBeSameTypeAs, testLinkSystem)

        n := fluent.MustBuildMap(basicnode.Prototype.Map, 1, func(na fluent.MapAssembler) {
            na.AssembleEntry("hello").AssignString("world")
        })

        lp := cidlink.LinkPrototype{cid.Prefix{
            Version:  1,    // Usually '1'.
            Codec:    uint64(multicodec.DagCbor), //0x71, // 0x71 means "dag-cbor"
            MhType:   uint64(multicodec.Sha2_256), //0x12, // 0x12 means "sha2-256"
            MhLength: 32,   // sha2-256 hash has a 32-byte sum.
        }}

		root := storageLinkSystem.MustStore(linking.LinkContext{}, lp, n)

        sc := carv1.NewSelectiveCar(context.Background(),
            storageLinkSystem.ReadStore,
            []carv1.Dag{{
                // CID of the root node of the DAG to traverse.
                Root: root.(cidlink.Link).Cid,

                // DAG traversal convenience selector that gives us "everything".
                Selector: everythingSelector(),
            }})

        // Create the file on disk that will store the car representation.
        carfile_v1, err := os.Create(fileprefix + ".v1.car")
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
        if err := carv2.WrapV1File(fileprefix+".v1.car", fileprefix+".v2.car"); err != nil {
            fmt.Printf("err wrapping carfile: %v\n", err)
            panic(err)
        }

        // TODO: Test reading back carfiles into StorageLinkSystem.

        // TODO: Delete test files.
    })
}

// This block makes a selector that recurses over the fluent reflection of
// our go struct (ipld map) and gives us "everything". It's a rip job from
// https://github.com/ipld/go-ipld-adl-hamt/issues/24, see also:
// https://github.com/ipld/go-ipld-prime/issues/171, and:
// https://github.com/ipld/go-ipld-prime/pull/199/files
func everythingSelector() datamodel.Node {
    ssb := sbuilder.NewSelectorSpecBuilder(basicnode.Prototype.Any)
    return ssb.ExploreFields(func(efsb sbuilder.ExploreFieldsSpecBuilder) {
        efsb.Insert("Links", ssb.ExploreIndex(1, ssb.ExploreRecursive(selector.RecursionLimitNone(), ssb.ExploreAll(ssb.ExploreRecursiveEdge()))))
    }).Node()
}

