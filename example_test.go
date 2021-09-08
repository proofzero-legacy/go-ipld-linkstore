package linkstore

import (
	"os"
    "fmt"
	"testing"
	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/linking"
	"github.com/ipld/go-ipld-prime/linking/cid"
	. "github.com/warpfork/go-wish"
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

		//root := storageLinkSystem.MustStore(linking.LinkContext{}, cidlink.LinkPrototype{}, ipld.BasicNode{})
		//t.Logf("Generated CID: %v\n", root)
	})
}

