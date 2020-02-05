ALTER TABLE validators ADD CONSTRAINT fk_validator_group FOREIGN KEY (affiliation) REFERENCES validator_groups (id);
