-- select records from the Song table.
SELECT
  "songId",
  "artist",
  "title",
  "difficulty",
  "level",
  "released"
FROM
  webapp_db.public."Song";

-- select particular song record from the Song table.
SELECT
  "songId",
  "artist",
  "title",
  "difficulty",
  "level",
  "released"
FROM
  webapp_db.public."Song"
WHERE
  "songId" = 1;

-- select songs in a pagination way from the Song table.
SELECT
  *
FROM webapp_db.public."Song"
ORDER BY
  "songId" ASC
LIMIT  2
OFFSET 0;

-- select songs based on artist or title - case insensitive search.
SELECT
  *
FROM
  webapp_db.public."Song"
WHERE
     artist ~* 'the Y'
  OR title  ~* 'pic M';

-- select records from the Rating table.
SELECT
  "id",
  "rate"
FROM
  webapp_db.public."Rating";

-- select particular song rating records from the Rating table.
SELECT
  "rate"
FROM
  webapp_db.public."Rating"
WHERE
  "id" = 4;

-- select particular song average, minimum and maximum rating from the Rating table.
SELECT
  AVG("rate")::NUMERIC(10,2) AS "avgRating",
  MIN("rate")                AS "minRating",
  MAX("rate")                AS "maxRating"
FROM
  webapp_db.public."Rating"
WHERE
  "id" = 9;

-- select average difficulty of all songs from Song table.
SELECT
  AVG("difficulty")::numeric(10,2) AS "avgDifficulty"
FROM
  webapp_db.public."Song";

-- select average difficulty of all songs belongs to particular level from Song table.
SELECT
  AVG("difficulty")::numeric(10,2) AS "avgDifficulty"
FROM
  webapp_db.public."Song"
WHERE
  "level" = 9;
