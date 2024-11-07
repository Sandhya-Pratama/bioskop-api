CREATE TABLE "User" (
  "user_id" SERIAL PRIMARY KEY,
  "username" VARCHAR(50) UNIQUE NOT NULL,
  "password" VARCHAR(100) NOT NULL,
  "role" VARCHAR(20) NOT NULL
);

CREATE TABLE "Bioskop" (
  "bioskop_id" SERIAL PRIMARY KEY,
  "nama_bioskop" VARCHAR(100) NOT NULL,
  "lokasi" VARCHAR(255) NOT NULL
);

CREATE TABLE "Film" (
  "film_id" SERIAL PRIMARY KEY,
  "judul" VARCHAR(100) NOT NULL,
  "durasi" INT NOT NULL,
  "genre" VARCHAR(50) NOT NULL
);

CREATE TABLE "Jadwal_Tayang" (
  "jadwal_id" SERIAL PRIMARY KEY,
  "film_id" INT NOT NULL,
  "bioskop_id" INT NOT NULL,
  "tanggal_tayang" DATE NOT NULL,
  "waktu_tayang" TIME NOT NULL
);

CREATE TABLE "Kursi" (
  "kursi_id" SERIAL PRIMARY KEY,
  "bioskop_id" INT NOT NULL,
  "jadwal_id" INT NOT NULL,
  "status_kursi" VARCHAR(20) NOT NULL
);

CREATE TABLE "Transaksi" (
  "transaksi_id" SERIAL PRIMARY KEY,
  "user_id" INT NOT NULL,
  "jadwal_id" INT NOT NULL,
  "kursi_id" INT NOT NULL,
  "status_transaksi" VARCHAR(20) NOT NULL,
  "tanggal_transaksi" TIMESTAMP NOT NULL,
  "refund" BOOLEAN DEFAULT false
);

ALTER TABLE "Jadwal_Tayang" ADD FOREIGN KEY ("film_id") REFERENCES "Film" ("film_id");

ALTER TABLE "Jadwal_Tayang" ADD FOREIGN KEY ("bioskop_id") REFERENCES "Bioskop" ("bioskop_id");

ALTER TABLE "Kursi" ADD FOREIGN KEY ("bioskop_id") REFERENCES "Bioskop" ("bioskop_id");

ALTER TABLE "Kursi" ADD FOREIGN KEY ("jadwal_id") REFERENCES "Jadwal_Tayang" ("jadwal_id");

ALTER TABLE "Transaksi" ADD FOREIGN KEY ("user_id") REFERENCES "User" ("user_id");

ALTER TABLE "Transaksi" ADD FOREIGN KEY ("jadwal_id") REFERENCES "Jadwal_Tayang" ("jadwal_id");

ALTER TABLE "Transaksi" ADD FOREIGN KEY ("kursi_id") REFERENCES "Kursi" ("kursi_id");

ALTER TABLE "Transaksi" ADD FOREIGN KEY ("kursi_id") REFERENCES "Transaksi" ("transaksi_id");

ALTER TABLE "Bioskop" ADD FOREIGN KEY ("bioskop_id") REFERENCES "Bioskop" ("nama_bioskop");

ALTER TABLE "User" ADD FOREIGN KEY ("role") REFERENCES "User" ("username");

ALTER TABLE "Film" ADD FOREIGN KEY ("film_id") REFERENCES "Film" ("judul");
