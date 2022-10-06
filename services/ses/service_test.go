package ses

import (
	"context"
	"testing"

	"github.com/flip-id/aws-sdk-go/client"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("valid new", func(t *testing.T) {
		client, err := New(context.TODO(), client.APSoutheast1)
		assert.NoError(t, err)
		assert.NotNil(t, client)
	})
}
