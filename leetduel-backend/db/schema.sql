-- To get the gen_random_uuid() function, you need to enable the pgcrypto extension.
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Users
CREATE TABLE users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  firebase_uid TEXT UNIQUE NOT NULL,
  username TEXT NOT NULL,
  elo INTEGER DEFAULT 400
);

-- Questions
CREATE TABLE questions (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  difficulty TEXT CHECK (difficulty IN ('easy', 'medium', 'hard')),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Test Cases (linked to questions)
CREATE TABLE test_cases (
  id SERIAL PRIMARY KEY,
  question_id INTEGER REFERENCES questions(id) ON DELETE CASCADE,
  input TEXT NOT NULL,
  expected_output TEXT NOT NULL
);