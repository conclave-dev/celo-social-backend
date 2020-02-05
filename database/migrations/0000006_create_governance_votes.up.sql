CREATE TABLE IF NOT EXISTS governance_votes(
   id serial NOT NULL PRIMARY KEY,
   proposal_id serial NOT NULL,
   user_id VARCHAR (100) NOT NULL,
   voted BOOLEAN NOT NULL,
   extra_data JSON
);
