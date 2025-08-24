-- create Song table.
CREATE TABLE IF NOT EXISTS webapp_db.public."Song"(
  "songId"     BIGSERIAL PRIMARY KEY,
  "artist"     VARCHAR   NOT NULL,
  "title"      VARCHAR   NOT NULL,
  "difficulty" FLOAT(2)  NOT NULL,
  "level"      SMALLINT  NOT NULL CHECK(0 < "level"  AND "level" < 100),
  "released"   VARCHAR   NOT NULL
);

-- create Rating table.
CREATE TABLE IF NOT EXISTS webapp_db.public."Rating"(
  "id" BIGINT   NOT NULL,
  "rate" SMALLINT NOT NULL CHECK(0 < "rate" AND "rate" < 6)
);
