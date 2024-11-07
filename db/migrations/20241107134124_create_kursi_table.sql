-- +goose Up
CREATE TABLE "Kursi" (
  "kursi_id" BIGSERIAL PRIMARY KEY,
  "bioskop_id" INT NOT NULL REFERENCES "Bioskop" ("bioskop_id"),
  "jadwal_id" INT NOT NULL REFERENCES "Jadwal_Tayang" ("jadwal_id"),
  "status_kursi" VARCHAR(20) NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS "Kursi";
