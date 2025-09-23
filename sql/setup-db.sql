CREATE TABLE Users (
    id int primary key auto_increment,
    name varchar(256) not null, 
    phone varchar(256) not null unique,
    createdAt timestamp default current_timestamp
) 

insert into Users(name , phone) values ("Amir" , "09921810208") , ("Mehrdad" , "09921537594");