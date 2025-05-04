-- migrate:up
alter     table accounts ADD constraint unique_citizen_id unique (citizen_id);

alter     table accounts ADD constraint unique_email unique (email);

alter     table accounts ADD constraint unique_phone unique (phone);

-- migrate:down
alter     table accounts
drop      constraint IF exists unique_citizen_id;

drop      constraint IF exists unique_email;

drop      constraint IF exists unique_phone;