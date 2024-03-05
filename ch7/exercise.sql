-- create database
create database go_kominfo;

-- create table
create table users(
  id serial primary key not null,
  username varchar(100) not null, 
  first_name varchar(255) not null, 
  last_name varchar(255),
  dob date,
  created_at timestamp default now() not null, 
  updated_at timestamp default now() not null, 
  deleted_at timestamp
);

-- insert data
insert 
into users("username", "first_name", "last_name", "dob") 
values('golang006', 'golang', 'student', '2024-01-01');
insert 
into users("username", "first_name", "last_name", "dob") 
values('golang007', 'golang7', 'student', '2024-01-01');

-- insert multiple data
insert 
into users("username", "first_name", "last_name", "dob") 
values
('golang009', 'golang9', 'student', '2024-01-01'),
('golang008', 'golang8', 'student', '2024-01-01');

-- generate random data
insert into users("username", "first_name", "last_name", "dob")
select 
	concat('username', generate_series(0,100)) as username, 
	md5(random()::text) as first_name, 
	md5(random()::text) as last_name,
	'2024-01-01' as dob; 

-- "" => column name
-- '' => data 


-- select data
select * from users;
-- hanya return id, username, dob
select "id", username, dob from users;

-- filter data
-- mendapatkan data yang memiliki first_name golang
select "id", username, dob from users
where first_name = 'golang';
-- mendapatkan data yang memiliki kata golang di username
select "id", username, dob from users
where username like '%golang%';

-- kalian buat 1 database
-- buat table users
-- insert 100 data

-- update data
update users set first_name = 'first_name'
where id = 5;

select * from users where id = 5;

-- delete data
-- hard delete
delete from users
where id = 5;

-- soft delete
update users set deleted_at = now()
where id = 88;

-- users u => alias 
select * from users u where id = 1;


create table user_media_socials(
  id serial primary key not null,
  user_id int not null,
  title varchar(100) not null,
  url text not null,
  created_at timestamp default now() not null, 
  updated_at timestamp default now() not null, 
  deleted_at timestamp,
  CONSTRAINT fk_user_media_social
      FOREIGN KEY(user_id) 
        REFERENCES users(id)
); 

insert into user_media_socials(user_id, title, url)
values (1, 'facebook', 'https://facebook.com');

select 
	u.id, 
	u.username, 
	ums.title, 
	ums.url  
from users u
join user_media_socials ums 
on u.id = ums.user_id;
