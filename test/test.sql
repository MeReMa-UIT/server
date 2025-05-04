select    *
from      accounts;

select    *
from      patients;

delete    from accounts
where     acc_id > 1;

delete    from patients
where     patient_id >= 1;

update    accounts
set       role = 'admin'
where     acc_id >= 1;