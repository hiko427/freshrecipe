-- name: CreateRecipe :one
INSERT INTO recipes (
  cuisine,
  num_people,
  recipe_time,
  unwanted_ingredients,
  wanted_ingredients,
  recipe
) VALUES (
  $1, $2 ,$3 ,$4 ,$5 ,$6
) RETURNING *;

-- name: GetAllRecipe :one
SELECT * FROM recipes;

-- name: GetRecipe :one
SELECT * FROM recipes
WHERE id = $1 LIMIT 1;

-- name: ListRecipe :many
SELECT * FROM recipes
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteRecipe :exec
DELETE FROM recipes
WHERE id = $1;