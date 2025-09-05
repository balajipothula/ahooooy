-- liquibase formatted sql

-- changeset BalajiPothula:2025-09-05T15:23:01Z
CREATE TABLE member (
  virtual_number VARCHAR(20) PRIMARY KEY,
  email          VARCHAR(255) UNIQUE NOT NULL,
  verified       BOOLEAN DEFAULT FALSE,
  first_name     VARCHAR(100),
  family_name    VARCHAR(100),
  dob            DATE,
  gender         VARCHAR(10),
  created_at     TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at     TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
--rollback DROP TABLE member;

-- changeset BalajiPothula:2025-09-05T15:23:02Z
CREATE INDEX idx_member_email ON member(email);
--rollback DROP INDEX idx_member_email;

-- changeset BalajiPothula:2025-09-05T15:23:03Z
CREATE OR REPLACE FUNCTION func_member_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;
--rollback DROP FUNCTION func_member_updated_at() CASCADE;

-- changeset BalajiPothula:2025-09-05T15:23:04Z
CREATE TRIGGER trigger_func_member_updated_at
BEFORE UPDATE ON member
FOR EACH ROW
EXECUTE FUNCTION func_updated_at();
--rollback DROP TRIGGER trigger_func_member_updated_at ON member;
