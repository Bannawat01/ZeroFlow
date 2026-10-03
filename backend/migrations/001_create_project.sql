CREATE TABLE projects (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    project_name text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),

    CONSTRAINT projects_name_not_blank
        CHECK (btrim(project_name) <> '')
);
