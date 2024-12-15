package perm

import (
	"testing"

	"github.com/0x28F4/aoc2024/utils"
)

func TestEqual(t *testing.T) {
	perms := Equal(3, []string{"a", "b"})
	utils.MustLen(perms, 2*2*2)
}
