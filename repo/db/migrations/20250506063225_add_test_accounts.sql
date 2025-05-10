-- migrate:up
INSERT    INTO accounts (citizen_id, password_hash, phone, email, role)
VALUES    (
          '123123123123',
          '$2a$10$bqcjrpmJ8qID2bkUscD15uFKzE0tiYFpbM3oex3GxCeqDy9IuhZ2K',
          '0111222333',
          '23520199@gm.uit.edu.vn',
          'admin'
          );

INSERT    INTO accounts (citizen_id, password_hash, phone, email, role)
VALUES    (
          '123412341234',
          '$2a$10$bqcjrpmJ8qID2bkUscD15uFKzE0tiYFpbM3oex3GxCeqDy9IuhZ2K',
          '0123456789',
          '23521734@gm.uit.edu.vn',
          'admin'
          );

-- migrate:down
DELETE    FROM accounts
WHERE     citizen_id = '123123123123';

DELETE    FROM accounts
WHERE     citizen_id = '123412341234';