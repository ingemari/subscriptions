CREATE TABLE subscriptions (
                               id SERIAL PRIMARY KEY,
                               service_name TEXT NOT NULL,
                               price INTEGER NOT NULL CHECK (price >= 0),
                               user_id UUID NOT NULL,
                               start_date DATE NOT NULL,
                               created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                               updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
