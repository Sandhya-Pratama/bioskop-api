-- +goose Up
CREATE TABLE "Transaksi" (
  "transaksi_id" BIGSERIAL PRIMARY KEY,
  "user_id" INT NOT NULL REFERENCES "User" ("user_id"),
  "jadwal_id" INT NOT NULL REFERENCES "Jadwal_Tayang" ("jadwal_id"),
  "kursi_id" INT NOT NULL REFERENCES "Kursi" ("kursi_id"),
  "status_transaksi" VARCHAR(20) NOT NULL,
  "tanggal_transaksi" TIMESTAMP NOT NULL,
  "refund" BOOLEAN DEFAULT false
);

-- +goose Down
DROP TABLE IF EXISTS "Transaksi";
