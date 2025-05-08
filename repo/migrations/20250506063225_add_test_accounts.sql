-- migrate:up
insert    into accounts (citizen_id, password_hash, phone, email, role)
values    (
          '123123123123',
          '$2a$10$bqcjrpmJ8qID2bkUscD15uFKzE0tiYFpbM3oex3GxCeqDy9IuhZ2K',
          '0111222333',
          '23520199@gm.uit.edu.vn',
          'admin'
          );

-- migrate:down
delete    from accounts
where     citizen_id = '123123123123';
