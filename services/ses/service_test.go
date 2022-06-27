package ses_test

import (
	"context"
	"flip/aws-sdk-go/client"
	"flip/aws-sdk-go/services/ses"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("valid new", func(t *testing.T) {
		client, err := ses.New(context.TODO(), client.APSoutheast1)
		assert.NoError(t, err)
		assert.NotNil(t, client)
	})
}
