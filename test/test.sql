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
          'admin',
          'merema-admin',
          '0123456789',
          'adu@gmail.com',
          'admin'
          );

select    *
from      accounts