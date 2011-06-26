package chunkstore

import (
	"log"
	"io"
	"os"

	. "chunkymonkey/types"
	"nbt"
)

// Returned to chunks to pull their data from.
type chunkReader struct {
	chunkTag *nbt.NamedTag
}

// Load a chunk from its NBT representation
func newChunkReader(reader io.Reader) (r *chunkReader, err os.Error) {
	chunkTag, err := nbt.Read(reader)
	if err != nil {
		return
	}

	r = &chunkReader{
		chunkTag: chunkTag,
	}

	return
}

func (r *chunkReader) ChunkLoc() ChunkXz {
	return ChunkXz{
		X: ChunkCoord(r.chunkTag.Lookup("/Level/xPos").(*nbt.Int).Value),
		Z: ChunkCoord(r.chunkTag.Lookup("/Level/zPos").(*nbt.Int).Value),
	}
}

func (r *chunkReader) Blocks() []byte {
	return r.chunkTag.Lookup("/Level/Blocks").(*nbt.ByteArray).Value
}

func (r *chunkReader) BlockData() []byte {
	return r.chunkTag.Lookup("/Level/Data").(*nbt.ByteArray).Value
}

func (r *chunkReader) BlockLight() []byte {
	return r.chunkTag.Lookup("/Level/BlockLight").(*nbt.ByteArray).Value
}

func (r *chunkReader) SkyLight() []byte {
	return r.chunkTag.Lookup("/Level/SkyLight").(*nbt.ByteArray).Value
}

func (r *chunkReader) HeightMap() []byte {
	return r.chunkTag.Lookup("/Level/HeightMap").(*nbt.ByteArray).Value
}

func (r *chunkReader) Entities() []*nbt.Compound {
	list := r.chunkTag.Lookup("/Level/Entities").(*nbt.List).Value
	entities := make([]*nbt.Compound, len(list))

	for idx, data := range list {
		entity, ok := data.(*nbt.Compound)
		if ok {
			entities[idx] = entity
		} else {
			log.Printf("Non-Compound entity in /Level/Entities/%d", idx)
		}
	}
	return entities
}

func (r *chunkReader) GetRootTag() *nbt.NamedTag {
	return r.chunkTag
}