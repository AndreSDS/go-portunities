CREATE TABLE
    opennings (
        id INTEGER PRIMARY KEY,
        role text NOT NULL,
        company text NOT NULL,
        location text NOT NULL,
        remote boolean NOT NULL,
        link text NOT NULL,
        salary INTEGER NOT NULL
    );