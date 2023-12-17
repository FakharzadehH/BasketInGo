CREATE TABLE
    users(
        id int NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
        username TEXT NOT NULL,
        password TEXT,
        created_at DATE DEFAULT NOW(),
        updated_at DATE DEFAULT NOW()
    );

CREATE TABLE
    baskets(
        id int NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
        data VARCHAR(2048) NOT NULL,
        user_id int NOT NULL,
        state TEXT not null,
        created_at DATE DEFAULT NOW(),
        updated_at DATE DEFAULT NOW()
    );