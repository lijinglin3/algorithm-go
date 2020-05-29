package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRestaurant(t *testing.T) {
	assert.ElementsMatch(t, []string{"Shogun"}, findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}))
	assert.ElementsMatch(t, []string{"Shogun", "KFC", "Burger King"}, findRestaurant([]string{"Tapioca Express", "Burger King", "Shogun", "KFC"}, []string{"KFC", "Shogun", "Burger King"}))
}
