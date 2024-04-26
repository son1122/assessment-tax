DROP TABLE IF EXISTS public.master_tax_level;
create table public.master_tax_level
(
    id        integer generated always as identity
        constraint master_tax_level_pk
            primary key,
    floor     integer not null,
    ceil      integer not null,
    create_at timestamp with time zone default now(),
    tax_value integer not null
);

alter table master_tax_level
    owner to postgres;

INSERT INTO public.master_tax_level (id, floor, ceil, create_at, tax_value)
VALUES (default, 0, 150000, '2024-04-18 15:48:08.675230 +00:00', 0);
INSERT INTO public.master_tax_level (id, floor, ceil, create_at, tax_value)
VALUES (default, 150000, 500000, '2024-04-18 15:48:08.675230 +00:00', 10);
INSERT INTO public.master_tax_level (id, floor, ceil, create_at, tax_value)
VALUES (default, 500000, 1000000, '2024-04-18 15:57:48.090174 +00:00', 15);
INSERT INTO public.master_tax_level (id, floor, ceil, create_at, tax_value)
VALUES (default, 1000000, 2000000, '2024-04-18 15:57:48.090174 +00:00', 20);
INSERT INTO public.master_tax_level (id, floor, ceil, create_at, tax_value)
VALUES (default, 2000000, -1, '2024-04-18 15:57:48.090174 +00:00', 35);
