package persist

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestMemoryStore(t *testing.T) {
	memoryStore := NewMemoryStore(1 * time.Minute)
	ctx := context.Background()

	expectVal := "123"
	require.Nil(t, memoryStore.Set(ctx, "test", expectVal, 1*time.Second))

	value := ""
	assert.Nil(t, memoryStore.Get(ctx, "test", &value))
	assert.Equal(t, expectVal, value)

	time.Sleep(1 * time.Second)
	assert.Equal(t, ErrCacheMiss, memoryStore.Get(ctx, "test", &value))
}
