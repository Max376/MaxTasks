-- MaxTasks initial relational schema.
-- The migration runner records this file as version 1.

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT users_email_not_blank CHECK (btrim(email) <> '')
);

CREATE UNIQUE INDEX users_email_lower_key
    ON users (lower(email));

CREATE TABLE task_lists (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT task_lists_user_fk
        FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT task_lists_name_not_blank CHECK (btrim(name) <> '')
);

CREATE UNIQUE INDEX task_lists_user_name_lower_key
    ON task_lists (user_id, lower(name));

CREATE INDEX task_lists_user_updated_idx
    ON task_lists (user_id, updated_at DESC, id);

CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    list_id UUID NOT NULL,
    title TEXT NOT NULL,
    notes TEXT,
    due_at TIMESTAMPTZ,
    completed_at TIMESTAMPTZ,
    position INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT tasks_list_fk
        FOREIGN KEY (list_id) REFERENCES task_lists (id) ON DELETE CASCADE,
    CONSTRAINT tasks_title_not_blank CHECK (btrim(title) <> ''),
    CONSTRAINT tasks_position_non_negative CHECK (position >= 0),
    CONSTRAINT tasks_completed_at_consistent CHECK (completed_at IS NULL OR completed_at >= created_at)
);

CREATE INDEX tasks_list_position_idx
    ON tasks (list_id, position, id);

CREATE INDEX tasks_list_due_at_idx
    ON tasks (list_id, due_at, id)
    WHERE due_at IS NOT NULL;

CREATE INDEX tasks_open_idx
    ON tasks (list_id, position, id)
    WHERE completed_at IS NULL;
