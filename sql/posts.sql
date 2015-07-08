CREATE TABLE IF NOT EXISTS posts (
    id           BIGINT        NOT NULL AUTO_INCREMENT,
    title        VARCHAR(1024) NOT NULL,
    content      TEXT          NOT NULL,
    author_id    BIGINT        NOT NULL,
    published    BOOLEAN       DEFAULT FALSE,
    created      TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
    updated      TIMESTAMP     DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    
    PRIMARY KEY (id),
    FOREIGN KEY (author) REFERENCES contributors(id)
);