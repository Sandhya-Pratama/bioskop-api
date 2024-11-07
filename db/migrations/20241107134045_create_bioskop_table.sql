-- +goose Up
CREATE TABLE "Bioskop" (
  "bioskop_id" BIGSERIAL PRIMARY KEY,
  "nama_bioskop" VARCHAR(100) NOT NULL,
  "lokasi" VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS "Bioskop";
