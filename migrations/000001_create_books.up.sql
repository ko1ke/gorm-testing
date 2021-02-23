CREATE SEQUENCE books_seq;
CREATE TABLE IF NOT EXISTS books(
   id INT DEFAULT nextval('books_seq')  PRIMARY KEY,
   title VARCHAR(255) NOT NULL,
   author VARCHAR(255) NOT NULL,
   created_at timestamp NOT NULL default CURRENT_TIMESTAMP,
   updated_at timestamp NOT NULL default CURRENT_TIMESTAMP,
   deleted_at timestamp
);
ALTER SEQUENCE books_seq OWNED BY books.id;
