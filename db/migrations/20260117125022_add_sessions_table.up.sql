CREATE TABLE sessions (
  id UUID PRIMARY KEY NOT NULL,
  user_email VARCHAR(255) NOT NULL,
  refresh_token VARCHAR(512) NOT NULL,
  is_revoked BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW(),
  expires_at TIMESTAMP
);
