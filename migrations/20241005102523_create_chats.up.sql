CREATE TABLE chats (
    id SERIAL PRIMARY KEY,
    user1_id bigserial NOT NULL,
    user2_id bigserial NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user1 FOREIGN KEY (user1_id) REFERENCES users(id),
    CONSTRAINT fk_user2 FOREIGN KEY (user2_id) REFERENCES users(id),
    CONSTRAINT unique_chat UNIQUE (user1_id, user2_id)
);