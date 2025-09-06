-- liquibase formatted sql

-- changeset BalajiPothula:2025-09-06T13:20:37Z
CREATE OR REPLACE FUNCTION func_member_updated_at()
RETURNS TRIGGER AS $FUNC$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$FUNC$ LANGUAGE plpgsql;
--rollback DROP FUNCTION func_member_updated_at() CASCADE;
