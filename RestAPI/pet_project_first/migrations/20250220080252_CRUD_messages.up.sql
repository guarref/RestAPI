-- create_tasks_table.up.sql
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_tables WHERE schemaname = 'public' AND tablename = 'tasks') THEN
        CREATE TABLE tasks (
            id SERIAL PRIMARY KEY,
            task VARCHAR(255) NOT NULL,
            is_done BOOLEAN DEFAULT FALSE,
            created_at TIMESTAMP NOT NULL DEFAULT NOW(),
            updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
            deleted_at TIMESTAMP DEFAULT NULL
        );
    END IF;
END $$;