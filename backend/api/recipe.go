package api

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/hiko427/myrecipe/db/sqlc"
	util "github.com/hiko427/myrecipe/utils"
	"github.com/lib/pq"
	"github.com/sashabaranov/go-openai"
)

type RecipeData struct {
	Cuisine             string `form:"cuisine" binding:"required"`
	NumPeople           int64  `form:"num_people" binding:"required"`
	Ingredients         string `form:"ingredients" binding:"required"`
	ExcludedIngredients string `form:"excluded_ingredients" binding:"required"`
	MealType            string `form:"meal_type" binding:"required"`
}

func (server *Server) createRecipe(ctx *gin.Context) {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	var recipeData RecipeData
	if err := ctx.ShouldBind(&recipeData); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	content := "あなたはプロのシェフです。以下の条件に基づいてレシビを生成してください。\n *ジャンル：" +
		recipeData.Cuisine + "\n *食事する人数：" + string(rune(recipeData.NumPeople)) + "人\n *入れたい食材（他の食材を加えても大丈夫、ここに書かれた食材を必ずしも全て使う必要はない）：" + recipeData.Ingredients + "\n *いらない食材:" + recipeData.ExcludedIngredients +
		"\n *レシピの時間帯:" + recipeData.MealType + "補足：生成されたレシピには料理名、食事する人数分の食材や調味料の分量、作り方、かかる時間などの情報が含まれる。"
	client := openai.NewClient(config.GPTKEY)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}
	arg := db.CreateRecipeParams{
		Cuisine:             recipeData.Cuisine,
		NumPeople:           int64(recipeData.NumPeople),
		Ingredients:         recipeData.Ingredients,
		ExcludedIngredients: recipeData.ExcludedIngredients,
		Mealtype:            recipeData.MealType,
		Recipe:              resp.Choices[0].Message.Content,
	}
	recipe, err := server.store.CreateRecipe(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.Redirect(http.StatusFound, fmt.Sprintf("/show/%d", recipe.ID))
}

type getRecipeRequest struct {
	ID int64 `uri:"recipeID" binding:"required,min=1"`
}

func (server *Server) showRecipe(ctx *gin.Context) {
	var req getRecipeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	recipe, err := server.store.GetRecipe(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.HTML(http.StatusOK, "show.tmpl", gin.H{
		"cuisine":             recipe.Cuisine,
		"numpeople":           recipe.NumPeople,
		"ingredients":         recipe.Ingredients,
		"excludedingredients": recipe.ExcludedIngredients,
		"mealtype":            recipe.Mealtype,
		"recipe":              recipe.Recipe,
	})
}
