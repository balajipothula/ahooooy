-- liquibase formatted sql

-- changeset BalajiPothula:2025-09-06T13:44:56Z
CREATE TRIGGER member_updated_at
BEFORE UPDATE ON member
FOR EACH ROW
EXECUTE FUNCTION member_updated_at();
--rollback DROP TRIGGER member_updated_at ON member;
