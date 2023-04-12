CREATE TABLE "recipes" (
  "id" bigserial PRIMARY KEY,
  "cuisine" varchar NOT NULL,
  "num_people" bigint NOT NULL,
  "ingredients" text NOT NULL,
  "excluded_ingredients" text NOT NULL,
  "mealtype" varchar NOT NULL,
  "recipe" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "recipes" ("cuisine");

CREATE INDEX ON "recipes" ("mealtype");

COMMENT ON COLUMN "recipes"."num_people" IS 'must be positive';
