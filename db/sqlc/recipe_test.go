package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomRecipe(t *testing.T) Recipe {
	arg := CreateRecipeParams{
		Cuisine:             "和食",
		NumPeople:           3,
		ExcludedIngredients: "玉ねぎ",
		Ingredients:         "鶏肉",
		Mealtype:            "morning",
		Recipe:              "レシピ",
	}
	recipe, err := testQueries.CreateRecipe(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, recipe)
	require.Equal(t, arg.Cuisine, recipe.Cuisine)
	require.Equal(t, arg.NumPeople, recipe.NumPeople)
	require.Equal(t, arg.Mealtype, recipe.Mealtype)
	require.Equal(t, arg.ExcludedIngredients, recipe.ExcludedIngredients)
	require.Equal(t, arg.Ingredients, recipe.Ingredients)
	require.Equal(t, arg.Recipe, recipe.Recipe)
	require.NotZero(t, recipe.ID)
	require.NotZero(t, recipe.CreatedAt)
	return recipe
}

func TestCreateRecipe(t *testing.T) {
	createRandomRecipe(t)
}
