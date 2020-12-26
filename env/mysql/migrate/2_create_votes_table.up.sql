create table votes (
    id integer auto_increment primary key,
    game_master_id integer,
    from_member_id integer,
    to_member_id integer
)
