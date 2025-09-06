-- liquibase formatted sql

-- changeset BalajiPothula:2025-09-06T13:20:37Z splitStatements:false endDelimiter:$$
CREATE OR REPLACE FUNCTION func_member_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
--rollback DROP FUNCTION func_member_updated_at() CASCADE;
