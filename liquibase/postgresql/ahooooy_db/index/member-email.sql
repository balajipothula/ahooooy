-- liquibase formatted sql

-- changeset BalajiPothula:2025-09-05T17:23:48Z
CREATE INDEX idx_member_email ON member(email);
--rollback DROP INDEX idx_member_email;
