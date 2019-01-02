package bot

import (
	"math/rand"
	"testing"
	"time"

	"github.com/chesnovsky/fliboobstier/storage"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	range_multiplier         = 1000
	random_id_number         = 25
	bucket_balance_max_delta = 0.07
)

var empty_strings = []string{}

func prepareRandomSeed() {
	rand.Seed(time.Now().UnixNano())
}

func genRandomUuids(uuids_num int) []string {
	string_array := []string{}
	for i := 1; i <= uuids_num; i++ {
		string_array = append(string_array, uuid.New().String())
	}
	return string_array
}

func TestSelectMediaTypeNone(t *testing.T) {
	prepareRandomSeed()
	null_media := storage.RegexActionElements{
		Images:    empty_strings,
		Stickers:  empty_strings,
		Gifs:      empty_strings,
		Documents: empty_strings,
	}
	for i := 0; i <= range_multiplier; i++ {
		assert.Equal(t, selectMediaType(null_media), MediaTypes.None)
	}
}

func TestSelectMediaTypeImagesOnly(t *testing.T) {
	prepareRandomSeed()
	for i := 0; i <= range_multiplier; i++ {
		test_media := storage.RegexActionElements{
			Images:    genRandomUuids(rand.Intn(random_id_number) + 1),
			Stickers:  empty_strings,
			Gifs:      empty_strings,
			Documents: empty_strings,
		}
		assert.Equal(t, MediaTypes.Image, selectMediaType(test_media))
	}
}

func TestSelectMediaTypeStickersOnly(t *testing.T) {
	prepareRandomSeed()
	for i := 0; i <= range_multiplier; i++ {
		test_media := storage.RegexActionElements{
			Images:    empty_strings,
			Stickers:  genRandomUuids(rand.Intn(random_id_number) + 1),
			Gifs:      empty_strings,
			Documents: empty_strings,
		}
		assert.Equal(t, MediaTypes.Sticker, selectMediaType(test_media))
	}
}

func TestSelectMediaTypeGifsOnly(t *testing.T) {
	prepareRandomSeed()
	for i := 0; i <= range_multiplier; i++ {
		test_media := storage.RegexActionElements{
			Images:    empty_strings,
			Stickers:  empty_strings,
			Gifs:      genRandomUuids(rand.Intn(random_id_number) + 1),
			Documents: empty_strings,
		}
		assert.Equal(t, MediaTypes.Gif, selectMediaType(test_media))
	}
}

func TestSelectMediaTypeBalanced(t *testing.T) {
	bucketImage := 0
	bucketGif := 0
	bucketSticker := 0
	prepareRandomSeed()
	for i := 0; i <= range_multiplier; i++ {
		random_strings := genRandomUuids(rand.Intn(random_id_number) + 1)
		test_media := storage.RegexActionElements{
			Images:    random_strings,
			Stickers:  random_strings,
			Gifs:      random_strings,
			Documents: random_strings,
		}
		selected_media := selectMediaType(test_media)
		switch selected_media {
		case MediaTypes.Image:
			bucketImage++
		case MediaTypes.Sticker:
			bucketSticker++
		case MediaTypes.Gif:
			bucketGif++
		}
	}

	bucket_part := 1 / 3.0
	assert.InDelta(t, bucket_part, float64(bucketGif)/range_multiplier, bucket_balance_max_delta)
	assert.InDelta(t, bucket_part, float64(bucketSticker)/range_multiplier, bucket_balance_max_delta)
	assert.InDelta(t, bucket_part, float64(bucketImage)/range_multiplier, bucket_balance_max_delta)
}

func TestNumInRange(t *testing.T) {
	assert.True(t, numInRange(0, 0, 0))
	assert.False(t, numInRange(0, 0, 1))
	assert.True(t, numInRange(1, 5, 3))
	assert.False(t, numInRange(1, 5, 6))
}
