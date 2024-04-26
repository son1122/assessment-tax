DROP TABLE IF EXISTS public.master_deduct;
create table public.master_deduct
(
    amount_deduct double precision,
    id            integer generated always as identity
        constraint donation_id
            primary key,
    is_active     boolean                  default false not null,
    create_at     timestamp with time zone default now() not null,
    type_deduct   varchar(50)                            not null,
    version       integer                                not null
);

alter table master_deduct
    owner to postgres;

INSERT INTO public.master_deduct (amount_deduct, id, is_active, create_at, type_deduct, version)
VALUES (100000, default, true, '2024-04-24 21:50:13.809298 +00:00', 'max-donation', 1);
INSERT INTO public.master_deduct (amount_deduct, id, is_active, create_at, type_deduct, version)
VALUES (100000, default, true, '2024-04-24 21:50:13.809298 +00:00', 'max-personal', 1);
INSERT INTO public.master_deduct (amount_deduct, id, is_active, create_at, type_deduct, version)
VALUES (100000, default, true, '2024-04-24 21:50:13.809298 +00:00', 'max-k-receipt', 1);
INSERT INTO public.master_deduct (amount_deduct, id, is_active, create_at, type_deduct, version)
VALUES (0, default, true, '2024-04-25 11:17:28.241702 +00:00', 'floor-k-receipt', 1);
INSERT INTO public.master_deduct (amount_deduct, id, is_active, create_at, type_deduct, version)
VALUES (10000, default, true, '2024-04-25 11:17:28.241702 +00:00', 'floor-personal', 1);
INSERT INTO public.master_deduct (amount_deduct, id, is_active, create_at, type_deduct, version)
VALUES (0, default, true, '2024-04-25 11:17:28.241702 +00:00', 'floor-donation', 1);
INSERT INTO public.master_deduct (amount_deduct, id, is_active, create_at, type_deduct, version)
VALUES (100000, default, true, '2024-04-22 15:20:18.834932 +00:00', 'donation', 1);
INSERT INTO public.master_deduct (amount_deduct, id, is_active, create_at, type_deduct, version)
VALUES (60000, default, true, '2024-04-25 17:39:01.092348 +00:00', 'personal', 1);
INSERT INTO public.master_deduct (amount_deduct, id, is_active, create_at, type_deduct, version)
VALUES (50000, default, true, '2024-04-25 17:52:24.864783 +00:00', 'k-receipt', 1);
