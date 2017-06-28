-- Internal
-----------

CREATE TABLE company (
    -- created_time         TIMESTAMP       NOT NULL DEFAULT now(),
    company_id           SERIAL          NOT NULL PRIMARY KEY,
    company_marker       STRING(20)      NOT NULL UNIQUE,
    company_name         STRING(50)      NOT NULL
);

CREATE TABLE account (
    -- created_time         TIMESTAMP       NOT NULL DEFAULT now(),
    account_id           SERIAL          NOT NULL PRIMARY KEY,
    company_id           INTEGER         NOT NULL REFERENCES company (company_id),
    external_account_id  STRING(100)     NOT NULL,
    account_data         BYTES           NOT NULL,    
    full_name            STRING(100)     NOT NULL,
    phone_number         STRING(30)      NOT NULL,
    email_address        string(30)      NOT NULL,

    UNIQUE (company_id, external_account_id)
);

CREATE TABLE account_lookup (
    -- created_time         TIMESTAMP       NOT NULL DEFAULT now(),
    field_type           BIT(16)         NOT NULL,
    field_value          STRING(300)     NOT NULL,
    account_id           INTEGER         NOT NULL REFERENCES account (account_id),
    -- field_stem          STRING(200)     NOT NULL,
    
    PRIMARY KEY (field_type, field_value, account_id)
    -- PRIMARY KEY (field_type, field_stem, account_id)
);


CREATE TABLE account_secret (
    -- created_time         TIMESTAMP       NOT NULL DEFAULT now(),
    account_id           INTEGER         NOT NULL REFERENCES account (account_id),
    secret               STRING(24)      NOT NULL PRIMARY KEY,
    active               BOOL            NOT NULL
);

CREATE TABLE account_enrollment (
    enrollment_id        SERIAL          NOT NULL PRIMARY KEY,
    account_id           INTEGER         NOT NULL REFERENCES account (account_id)
);

CREATE TABLE account_funnel (
    account_id           INTEGER        NOT NULL REFERENCES account (account_id),
    funnel_step          INTEGER        NOT NULL,
    
    PRIMARY KEY (account_id, funnel_step)
);