insert    into accounts (
          citizen_id,
          username,
          password_hash,
          phone,
          email,
          role
          )
values    (
          '123456123456',
          'merema-admin',
          '0123456789',
          'adu@gmail.com',
          'admin'
          );

select    *
from      accounts;

delete    from accounts
where     acc_id >= 1;