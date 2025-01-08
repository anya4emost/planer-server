CREATE TYPE sex AS ENUM ('мужчина', 'женщина', 'иное');

CREATE TYPE event_repit_type AS ENUM ('EveryHour', 'EveryDay', 'EveryWeek', 'EveryMonth', 'EveryYear');
CREATE TYPE event_remind_type AS ENUM ('FiveMinBefore', 'TenMinBefore', 'FifteenMinBefore', 'ThirtyMinBefore', 'HourBefore', 'DayBefore', 'WeekBefore', 'MonthBefore');
CREATE TYPE event_category AS ENUM ('Task', 'MemorableDate');
CREATE TYPE task_type AS ENUM ('Urgent', 'Important');
CREATE TYPE task_status AS ENUM ('Analysis', 'InProgress', 'Done', 'Canceled', 'OnHold');

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

create table tasks(
      id text not null primary key default nanoid(),
      name text not null,
      status task_status not null,
      description text,
      icon text,
      color text,
      type task_type not null,
      creator_id text not null references users(id),
      doer_id text not null references users(id),
      aim_id text not null references aims(id)
);

create table events(
      id text not null primary key default nanoid(),
      category event_category,
      date timestamptz,
      time time not null,
      repit event_repit_type,
      remind event_remind_type,
      custom_category_id text references custom_category(id),
      task_id text unique references tasks(id) on delete cascade
);

create table sessions(
     refresh_token text not null primary key,
     user_id text not null references users(id),
     created_at text,
     expires_at text,
     family text,
     revoked boolean
);
