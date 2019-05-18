package aws_test

import (
	"go-scripts/switch-aws-roles/pkg/aws"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRoles(t *testing.T) {
	_, err := aws.GetRoles()
	assert.Equal(t, nil, err)
}
