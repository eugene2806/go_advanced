CREATE TABLE users
(
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(64) NOT NULL
);

CREATE TABLE bookmarks
(
    bookmark TEXT NOT NULL
)