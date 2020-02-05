CREATE TABLE IF NOT EXISTS validators(
   id VARCHAR (100) NOT NULL PRIMARY KEY,
   name VARCHAR (200) NOT NULL,
   signer_index SMALLINT NOT NULL,
   affiliation VARCHAR (100) NOT NULL,
   score BIGINT NOT NULL,
   block_number BIGINT NOT NULL
);
