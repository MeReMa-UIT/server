-- migrate:up
insert    into accounts (citizen_id, password_hash, phone, email, role)
values    (
          '123123123123',
          '$2a$10$l9RU0OPUn6cluh8WKf3JAOiyPNrquxx2o3iJz3aDJ2SlKnXlbx/2G',
          '0111222333',
          '23520199@gm.uit.edu.vn',
          'admin'
          );

-- migrate:down
delete    from accounts
where     citizen_id = '123123123123';