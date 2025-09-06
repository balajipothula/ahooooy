# psql is the PostgreSQL interactive terminal.
#
# --host     : database server host or socket directory
# --port     : database server port (default: 5432)
# --username : database user name
# --password : force password prompt
# --dbname   : database name to connect to
psql \
  --host=aws-1-ap-south-1.pooler.supabase.com \
  --port=6543 \
  --username=postgres.kscpngmnxlqwmkzjlylw \
  --password

# creates a new PostgreSQL database with default settings.
CREATE DATABASE ahooooy_db;

# creates a new PostgreSQL database with custom settings.
#
# ENCODING   : Character set encoding to use in the new database.
# LC_COLLATE : Sets LC_COLLATE in the database server's operating system environment.
# LC_CTYPE   : Sets LC_CTYPE in the database server's operating system environment.
# TEMPLATE   : The name of the template from which to create the new database.
CREATE DATABASE ahooooy_db
  WITH OWNER = 'postgres'
  ENCODING   = 'UTF8'
  LC_COLLATE = 'en_US.UTF-8'
  LC_CTYPE   = 'en_US.UTF-8'
  TEMPLATE   = template0;

# 
SELECT * FROM "public"."databasechangelog";

# 
SELECT * FROM "public"."databasechangeloglock";
