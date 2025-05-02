-- migrate:up
alter     table accounts ADD constraint unique_citizen_id unique (citizen_id);

-- migrate:down
alter     table accounts
drop      constraint IF exists unique_citizen_id;