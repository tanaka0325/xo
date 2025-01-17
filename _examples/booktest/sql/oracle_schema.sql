CREATE TABLE authors (
  author_id INTEGER GENERATED ALWAYS AS IDENTITY NOT NULL CONSTRAINT authors_pkey PRIMARY KEY,
  name NVARCHAR2(255) DEFAULT '' NOT NULL
);

CREATE INDEX authors_name_idx ON authors(name);

CREATE TABLE books (
  book_id INTEGER GENERATED ALWAYS AS IDENTITY NOT NULL CONSTRAINT books_pkey PRIMARY KEY,
  author_id INTEGER NOT NULL CONSTRAINT books_author_id_fkey REFERENCES authors(author_id),
  isbn NVARCHAR2(255) DEFAULT '' NOT NULL CONSTRAINT books_isbn_key UNIQUE,
  title NVARCHAR2(255) DEFAULT '' NOT NULL,
  year INTEGER DEFAULT 2000 NOT NULL,
  available TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
  tags NCLOB DEFAULT '' NOT NULL
);

CREATE INDEX books_title_idx ON books(title, year);

CREATE FUNCTION say_hello(name IN NVARCHAR2) RETURN INTEGER AS
BEGIN
  RETURN 'hello ' || name\;
END\;;
