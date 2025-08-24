-- insert record into the Song table.
INSERT INTO webapp_db.public."Song"(
  "artist",
  "title",
  "difficulty",
  "level",
  "released"
)
VALUES(
  'The Yousicians',
  'A New Kennel',
  9.9,
  9,
  '2022-02-09'
);

-- insert record into the Rating table.
INSERT INTO webapp_db.public."Rating"(
  "id",
  "rate"
)
VALUES(
  5,
  2
);

-- bulk insertion into the Song table.
-- insert records into the Song table.
INSERT INTO webapp_db.public."Song"("artist", "title", "difficulty", "level", "released")
VALUES
  ('The Yousicians', 'Lycanthropic Metamorphosis',  '14.6',  '13', '2016-10-26'),
  ('The Yousicians', 'A New Kennel',                '9.1',   '9',  '2010-02-03'),
  ('Mr Fastfinger',  'Awaki-Waki',                  '15',    '13', '2012-05-11'),
  ('The Yousicians', 'You have Got The Power',      '13.22', '13', '2014-12-20'),
  ('The Yousicians', 'Wishing In The Night',        '10.98', '9',  '2016-01-01'),
  ('The Yousicians', 'Opa Opa Ta Bouzoukia',        '14.66', '13', '2013-04-27'),
  ('The Yousicians', 'Greasy Fingers - boss level', '2',     '3',  '2016-03-01'),
  ('The Yousicians', 'Alabama Sunrise',             '5',     '6',  '2016-04-01'),
  ('The Yousicians', 'Alabama Sunset',              '4',     '9',  '2016-04-01'),
  ('The Yousicians', 'Can not Buy Me Skills',       '9',     '9',  '2016-05-01'),
  ('The Yousicians', 'Vivaldi Allegro Mashup',      '13',    '13', '2016-06-01'),
  ('The Yousicians', 'Babysitting',                 '7',     '6',  '2016-07-01');

-- insert record into the Rating table.
INSERT INTO webapp_db.public."Rating"("id", "rate")
VALUES
  ( 1, 2),
  ( 2, 1),
  ( 3, 4),
  ( 4, 2),
  ( 5, 4),
  ( 6, 3),
  ( 7, 1),
  ( 8, 4),
  ( 9, 5),
  (10, 2),
  (11, 4),
  (12, 3),
  ( 1, 4),
  ( 2, 2),
  ( 3, 1),
  ( 4, 5),
  ( 5, 1),
  ( 6, 3),
  ( 7, 2),
  ( 8, 5),
  ( 9, 2),
  (10, 1),
  (11, 4),
  (12, 3);
