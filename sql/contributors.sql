CREATE TABLE IF NOT EXISTS contributors (
    id         BIGINT       NOT NULL AUTO_INCREMENT,
    name       VARCHAR(255) NOT NULL,
    email      VARCHAR(255) NOT NULL,
    password   TEXT         NOT NULL,
    about      TEXT         DEFAULT NULL,
    join_since TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    photo      TEXT         DEFAULT NULL,
    
    PRIMARY KEY (id)
);

ALTER TABLE contributors ADD INDEX email (email);
ALTER TABLE contributors ADD INDEX name (name);