CREATE TABLE links (
   link_id        BIGINT      AUTO_INCREMENT PRIMARY KEY ,
   link           TEXT        NOT NULL,
   shortened_link VARCHAR(12) NOT NULL
);