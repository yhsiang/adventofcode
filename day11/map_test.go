package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap_Search(t *testing.T) {
	t.Run("leftmost L", func(t *testing.T) {
		lines := []string{
			".............",
			".L.L.#.#.#.#.",
			".............",
		}

		map1 := NewMap(lines)

		assert.Equal(t, ".", map1.Search(1, 1, "w"))
		assert.Equal(t, ".", map1.Search(1, 1, "nw"))
		assert.Equal(t, ".", map1.Search(1, 1, "n"))
		assert.Equal(t, ".", map1.Search(1, 1, "ne"))
		assert.Equal(t, "L", map1.Search(1, 1, "e"))
		assert.Equal(t, ".", map1.Search(1, 1, "se"))
		assert.Equal(t, ".", map1.Search(1, 1, "s"))
		assert.Equal(t, ".", map1.Search(1, 1, "sw"))
	})

	t.Run("all occupied", func(t *testing.T) {
		lines := []string{
			".......#.",
			"...#.....",
			".#.......",
			".........",
			"..#L....#",
			"....#....",
			".........",
			"#........",
			"...#.....",
		}

		map1 := NewMap(lines)

		assert.Equal(t, "#", map1.Search(4, 3, "w"))
		assert.Equal(t, "#", map1.Search(4, 3, "nw"))
		assert.Equal(t, "#", map1.Search(4, 3, "n"))
		assert.Equal(t, "#", map1.Search(4, 3, "ne"))
		assert.Equal(t, "#", map1.Search(4, 3, "e"))
		assert.Equal(t, "#", map1.Search(4, 3, "se"))
		assert.Equal(t, "#", map1.Search(4, 3, "s"))
		assert.Equal(t, "#", map1.Search(4, 3, "sw"))
	})

	t.Run("no occupied", func(t *testing.T) {
		lines := []string{
			".##.##.",
			"#.#.#.#",
			"##...##",
			"...L...",
			"##...##",
			"#.#.#.#",
			".##.##.",
		}

		map1 := NewMap(lines)

		assert.Equal(t, ".", map1.Search(3, 3, "w"))
		assert.Equal(t, ".", map1.Search(3, 3, "nw"))
		assert.Equal(t, ".", map1.Search(3, 3, "n"))
		assert.Equal(t, ".", map1.Search(3, 3, "ne"))
		assert.Equal(t, ".", map1.Search(3, 3, "e"))
		assert.Equal(t, ".", map1.Search(3, 3, "se"))
		assert.Equal(t, ".", map1.Search(3, 3, "s"))
		assert.Equal(t, ".", map1.Search(3, 3, "sw"))
	})

}
