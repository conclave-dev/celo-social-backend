CREATE TABLE IF NOT EXISTS validator_groups(
   id VARCHAR (100) NOT NULL PRIMARY KEY,
   name VARCHAR (200) NOT NULL, 
   total_votes BIGINT NOT NULL,
   block_number BIGINT NOT NULL
);
