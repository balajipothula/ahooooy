-- liquibase formatted sql

-- changeset BalajiPothula:2025-09-06T09:21:15Z
CREATE INDEX idx_member_email ON member(email);
--rollback DROP INDEX idx_member_email;
