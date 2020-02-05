ALTER TABLE governance_votes ADD CONSTRAINT fk_governance_vote_proposal FOREIGN KEY (proposal_id) REFERENCES governance_proposals (id);

ALTER TABLE governance_votes ADD CONSTRAINT fk_governance_vote_user FOREIGN KEY (user_id) REFERENCES users (id);
