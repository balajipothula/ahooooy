-- liquibase formatted sql

-- changeset BalajiPothula:2025-09-05T17:28:26Z
CREATE TRIGGER trigger_func_member_updated_at
BEFORE UPDATE ON member
FOR EACH ROW
EXECUTE FUNCTION func_member_updated_at();
--rollback DROP TRIGGER trigger_func_member_updated_at ON member;
