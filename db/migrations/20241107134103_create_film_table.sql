-- +goose Up
CREATE TABLE "Film" (
  "film_id" BIGSERIAL PRIMARY KEY,
  "judul" VARCHAR(100) NOT NULL,
  "durasi" INT NOT NULL,
  "genre" VARCHAR(50) NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS "Film";
