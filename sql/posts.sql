CREATE TABLE IF NOT EXISTS posts (
    id           BIGINT        NOT NULL AUTO_INCREMENT,
    title        VARCHAR(1024) NOT NULL,
    content      TEXT          NOT NULL,
    author       BIGINT        NOT NULL,
    published    BOOLEAN       DEFAULT FALSE,
    created      DATETIME      DEFAULT CURRENT_TIMESTAMP,
    updated      DATETIME      DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    
    PRIMARY KEY (id),
    FOREIGN KEY (author) REFERENCES contributors(id)
);