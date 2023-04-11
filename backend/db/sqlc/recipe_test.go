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
		RecipeTime:          "morning",
		UnwantedIngredients: "玉ねぎ",
		WantedIngredients:   "鶏肉",
		Recipe:              "レシピ",
	}
	recipe, err := testQueries.CreateRecipe(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, recipe)
	require.Equal(t, arg.Cuisine, recipe.Cuisine)
	require.Equal(t, arg.NumPeople, recipe.NumPeople)
	require.Equal(t, arg.RecipeTime, recipe.RecipeTime)
	require.Equal(t, arg.UnwantedIngredients, recipe.UnwantedIngredients)
	require.Equal(t, arg.WantedIngredients, recipe.WantedIngredients)
	require.Equal(t, arg.Recipe, recipe.Recipe)
	require.NotZero(t, recipe.ID)
	require.NotZero(t, recipe.CreatedAt)
	return recipe
}

func TestCreateRecipe(t *testing.T) {
	createRandomRecipe(t)
}
