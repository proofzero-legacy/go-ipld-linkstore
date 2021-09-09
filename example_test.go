package linkstore

import (
	"os"
    "fmt"
	"testing"
	"github.com/ipfs/go-cid"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/fluent"
	"github.com/ipld/go-ipld-prime/linking"
	"github.com/ipld/go-ipld-prime/linking/cid"
	"github.com/ipld/go-ipld-prime/node/basicnode"
    "github.com/multiformats/go-multicodec"
	. "github.com/warpfork/go-wish"

	// Note: this dagcbor import populates the multicodec table as a side effect.
	// The CID LinkSystem requires this to function. See:
	// https://github.com/ipld/go-ipld-prime/blob/fc47eb2f400c1ab39da4df91c4e03f04e34e26cb/linkingExamples_test.go#L58
	_ "github.com/ipld/go-ipld-prime/codec/dagcbor"
)

func TestMain(m *testing.M) {
	failedTestCount := m.Run()
	os.Exit(failedTestCount)
}

func TestNewStorage(t *testing.T) {
	t.Run("Writing v1 carfile and wrapping to v2.", func(t *testing.T) {
        fmt.Println("Writing v1 carfile and wrapping to v2!")
		var testBlockReadOpener ipld.BlockReadOpener
		var testBlockWriteOpener ipld.BlockWriteOpener
		var testLinkSystem linking.LinkSystem

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
        fmt.Printf("CID: %v\n", root)
		t.Logf("Generated CID: %v\n", root)
	})
}

