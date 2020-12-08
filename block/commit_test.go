package block

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zarbchain/zarb-go/crypto"
	simpleMerkle "github.com/zarbchain/zarb-go/libs/merkle"
)

func TestNilCommitHash(t *testing.T) {
	var c Commit
	assert.Equal(t, c.Hash(), crypto.UndefHash)
}

func TestCommitMerkle(t *testing.T) {
	b, _ := GenerateTestBlock(nil)

	commiters := b.LastCommit().Commiters()
	data := make([]crypto.Hash, len(commiters))
	for i, c := range commiters {
		b := c.Address.RawBytes()
		data[i] = crypto.HashH(b)
	}
	merkle := simpleMerkle.NewTreeFromHashes(data)
	assert.Equal(t, merkle.Root(), b.LastCommit().CommitersHash())
}

func TestCommitSanityCheck(t *testing.T) {
	b, _ := GenerateTestBlock(nil)
	c := b.LastCommit()
	assert.NoError(t, c.SanityCheck())
	c.data.Commiters[0].Signed = false
	assert.Error(t, c.SanityCheck())
	c.data.Commiters[3].Signed = true
	assert.NoError(t, c.SanityCheck())
}