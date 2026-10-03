CREATE TABLE tasks (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    project_id bigint NOT NULL REFERENCES projects(id),
    title text NOT NULL,
    status text NOT NULL DEFAULT 'backlog',
    assignee_id bigint NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    CHECK (btrim(title) <> ''),
    CHECK (
        status IN (
            'backlog',
            'todo',
            'in_progress',
            'review',
            'done'
        )
    ),
    FOREIGN KEY (project_id, assignee_id) REFERENCES project_members (project_id, user_id)
);