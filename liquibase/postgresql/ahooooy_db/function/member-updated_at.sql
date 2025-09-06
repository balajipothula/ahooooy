-- liquibase formatted sql

-- changeset BalajiPothula:2025-09-06T13:39:45Z splitStatements:false endDelimiter:$$
CREATE OR REPLACE FUNCTION member_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
--rollback DROP FUNCTION member_updated_at() CASCADE;
