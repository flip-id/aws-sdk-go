package ses

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("valid new", func(t *testing.T) {
		client, err := New(context.TODO(), &ServiceOption{})
		assert.NoError(t, err)
		assert.NotNil(t, client)
	})
}
