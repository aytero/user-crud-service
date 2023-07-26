package service

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDeleteFileErr(t *testing.T) {
	err := DeleteFile("/temp/1")
	require.Error(t, err)
}
