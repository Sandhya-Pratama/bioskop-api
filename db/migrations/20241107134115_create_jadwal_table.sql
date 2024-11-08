-- +goose Up
CREATE TABLE "Jadwal_Tayang" (
  "jadwal_id" BIGSERIAL PRIMARY KEY,
  "film_id" INT NOT NULL REFERENCES "Film" ("film_id"),
  "bioskop_id" INT NOT NULL REFERENCES "Bioskop" ("bioskop_id"),
  "tanggal_tayang" DATE NOT NULL,
  "waktu_tayang" TIME NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS "Jadwal_Tayang";

