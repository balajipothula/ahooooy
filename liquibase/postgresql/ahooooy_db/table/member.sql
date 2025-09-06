-- liquibase formatted sql

-- changeset BalajiPothula:2025-09-06T07:11:01Z
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
