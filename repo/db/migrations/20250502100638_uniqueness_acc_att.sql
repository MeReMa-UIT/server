-- migrate:up
ALTER     TABLE accounts ADD CONSTRAINT unique_citizen_id UNIQUE (citizen_id);

ALTER     TABLE accounts ADD CONSTRAINT unique_email UNIQUE (email);

ALTER     TABLE accounts ADD CONSTRAINT unique_phone UNIQUE (phone);

-- migrate:down
ALTER     TABLE accounts
DROP      CONSTRAINT IF EXISTS unique_citizen_id;

ALTER     TABLE accounts
DROP      CONSTRAINT IF EXISTS unique_email;

ALTER     TABLE accounts
DROP      CONSTRAINT IF EXISTS unique_phone;