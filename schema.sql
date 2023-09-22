DROP TABLE IF EXISTS comments;

DROP TABLE IF EXISTS posts;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE posts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v1(),
    created_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    content TEXT,
    link TEXT UNIQUE NOT NULL
);

CREATE TABLE comments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v1(),
    news_id UUID REFERENCES posts(id) ON DELETE CASCADE,
    parent_id UUID DEFAULT NULL,
    created_at TIMESTAMP NOT NULL,
    content TEXT NOT NULL
);

