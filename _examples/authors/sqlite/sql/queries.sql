/* name: getAuthor :one */
SELECT * FROM authors
WHERE id = ? LIMIT 1;

/* name: listAuthors :many */
SELECT * FROM authors
ORDER BY name;

/* name: createAuthor :execresult */
INSERT INTO authors (
  name, bio
) VALUES (
  ?, ? 
);

/* name: deleteAuthor :exec */
DELETE FROM authors
WHERE id = ?;