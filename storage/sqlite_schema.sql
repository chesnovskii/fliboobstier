CREATE TABLE regex_action_images (
    action_id TEXT,
    element_id TEXT
);
CREATE UNIQUE INDEX regex_action_images_idx
ON regex_action_images(action_id, element_id);

CREATE TABLE regex_action_gifs (
    action_id TEXT,
    element_id TEXT
);
CREATE UNIQUE INDEX regex_action_gifs_idx
ON regex_action_gifs(action_id, element_id);

CREATE TABLE regex_action_documents (
    action_id TEXT,
    element_id TEXT
);
CREATE UNIQUE INDEX regex_action_documents_idx
ON regex_action_documents(action_id, element_id);

CREATE TABLE regex_action_stickers (
    action_id TEXT,
    element_id TEXT
);
CREATE UNIQUE INDEX regex_action_stickers_idx
ON regex_action_stickers(action_id, element_id);

INSERT INTO regex_action_images VALUES ("gopstop", "AgADAgADnqoxG22cUEn8gzuDaxEbmwvotw4ABP9cj0qxecfJtEsFAAEC");
INSERT INTO regex_action_gifs VALUES ("gopstop", "CgADBAADDaAAAvwYZAcsF65UbbXDbgI");
INSERT INTO regex_action_stickers VALUES ("gopstop", "CAADAgADXwMAAgw7AAEKTh8jAAH9Q-gAAQI");
INSERT INTO regex_action_stickers VALUES ("normal", "CAADAgADnQUAAlOx9wMjvcls38LyPwI");


CREATE TABLE admins (
    admin_login TEXT
);
CREATE UNIQUE INDEX admins_idx
ON admins(admin_login);

INSERT INTO admins VALUES ("newrushbolt"), ("cleargray"), ("chesnovsky");
