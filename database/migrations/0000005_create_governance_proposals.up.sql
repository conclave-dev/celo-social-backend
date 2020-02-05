CREATE TABLE IF NOT EXISTS governance_proposals(
   id serial NOT NULL PRIMARY KEY,
   title VARCHAR (200) NOT NULL,
   description TEXT NOT NULL
);
