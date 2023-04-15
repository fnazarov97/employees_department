create database employees;

create table role (
    "id" uuid primary key not null,
    "r_title" varchar not null,
    "description" varchar not null,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);

create table department (
   "id" uuid primary key not null,
   "d_number" int,
   "name" varchar,
   "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   "updated_at" TIMESTAMP
);

create table employee (
    "id" uuid primary key not null,
    "f_name" varchar(55) not null,
    "l_name" varchar(55) not null,
    "address" varchar not null,
    "phone" varchar,
    "role_id" uuid references role(id),
    "department_id" uuid references department(id),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP
);



