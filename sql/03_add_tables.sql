CREATE TYPE sex AS ENUM ('мужчина', 'женщина', 'иное');

CREATE TYPE event_repit_type AS ENUM ('repitable', 'no_repitable');
CREATE TYPE event_remind_type AS ENUM ('one_hour', 'one_day', 'not_remind');
CREATE TYPE event_category AS ENUM ('category_1', 'category_2', 'category_3');
CREATE TYPE task_type AS ENUM ('type_1', 'type_2', 'type_3');
CREATE TYPE task_status AS ENUM ('done', 'in_progress', 'on_hold');

create table groups(
     id text not null primary key default nanoid(),
     name text not null
);

create table user_groups(
    id text not null primary key default nanoid(),
    user_id text not null references users(id),
    group_id text not null references groups(id)
);

create table aims(
      id text not null primary key default nanoid(),
      user_id text not null references users(id),
      name text not null
);

create table custom_category(
    id text not null primary key default nanoid(),
    name text not null
);

create table events(
      id text not null primary key default nanoid(),
      category event_category,
      date timestamptz,
      time time not null,
      repit event_repit_type,
      remind event_remind_type,
      custom_category_id text not null references custom_category(id)
);

create table tasks(
      id text not null primary key default nanoid(),
      status task_status,
      description text,
      icon text,
      color text,
      type task_type,
      creator text not null references users(id),
      doer text not null references users(id),
      aim text not null references aims(id),
      event_id text references events(id)
);
