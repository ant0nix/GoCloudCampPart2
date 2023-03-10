CREATE TABLE playlist
(
    id serial unique,
    duration int not null
);

CREATE PROCEDURE insert_playlist(duration_in INT)
LANGUAGE plpgsql
AS $$
BEGIN
    INSERT INTO playlist (duration) VALUES (duration_in);
    COMMIT;
END;
$$;
