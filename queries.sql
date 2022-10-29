-- SQLite
create table if not exists categories(id string, name string, description string);

create table if not exists courses(id string, name string, description string, categoryId string);
