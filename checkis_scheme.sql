create table if not exists public.insurers (
    regnum varchar(10) not null unique,
    kpsnum varchar(10) not null default '',
    inn varchar(12) not null default '',
    kpp varchar(9) not null default '',
    fullname varchar(300) not null default '',
    shortname varchar(300) not null default '',
    postindex varchar(6) not null default '',
    postaddress varchar(300) not null default ''
);

create table if not exists public.staff (
    id serial primary key,
    fio varchar(300) not null default '',
    fio_dative varchar(300) not null default '',
    fio_genitive varchar(300) not null default '',
    position varchar(100) not null default '',
    position_dative varchar(100) not null default '',
    position_genitive varchar(100) not null default '',
    phone varchar(100) not null default '',
    signer integer not null default 0
);

create table if not exists public.docs_list (
    id serial primary key,
    flag varchar(300),
    doc_title text,
    unique (flag, doc_title)
);

create table if not exists public.checks (
    id serial primary key,
    -- страхователь
    insurer_regnum varchar(10) not null default '',
    insurer_kpsnum varchar(10) not null default '',
    insurer_inn varchar(12) not null default '',
    insurer_kpp varchar(9) not null default '',
    insurer_fullname varchar(300) not null default '',
    insurer_shortname varchar(300) not null default '',
    insurer_postindex varchar(6) not null default '',
    insurer_postaddress varchar(300) not null default '',
    -- периоды
    check_date_start date not null default CURRENT_DATE,
    check_date_end date not null default CURRENT_DATE,
    check_period_start date not null default CURRENT_DATE,
    check_period_end date not null default CURRENT_DATE,
    -- проверяющий
    inspector_id integer not null default 0,
    inspector_fio varchar(300) not null default '',
    inspector_position varchar(100) not null default '',
    inspector_phone varchar(100) not null default '',
    -- 1. требование
    docs_req_number integer not null default 0,
    docs_req_date date not null default CURRENT_DATE,
    docs_req_flag_sicklist integer not null default 0,
    docs_req_flag_birth integer not null default 0,
    docs_req_flag_1_early integer not null default 0,
    docs_req_flag_1_birth integer not null default 0,
    docs_req_flag_15_child integer not null default 0,
    docs_req_flag_accident integer not null default 0,
    docs_req_flag_vacation integer not null default 0,
    docs_req_flag_4days integer not null default 0,
    docs_req_flag_burial_soc integer not null default 0,
    docs_req_flag_burial_spec integer not null default 0,
    -- 2. докладная записка
    docs_memorandum_number integer not null default 0,
    docs_memorandum_date date not null default CURRENT_DATE,
    -- 3. акт
    docs_act_number integer not null default 0,
    docs_act_date date not null default CURRENT_DATE,
    -- 4. решение
    docs_decision_number integer not null default 0,
    docs_decision_date date not null default CURRENT_DATE,
    -- 5. требование
    docs_charge_number integer not null default 0,
    docs_charge_date date not null default CURRENT_DATE
);

create table if not exists act_list (
    id_check integer not null,
    pay_year integer not null,
    pay_month integer not null,
    overpay_sicklist double precision not null default 0,
    overpay_birth double precision not null default 0,
    overpay_1_early double precision not null default 0,
    underpay_sicklist double precision not null default 0,
    underpay_birth double precision not null default 0,
    underpay_1_early double precision not null default 0,
    ndfl_sicklist double precision not null default 0,
    postal_expenses double precision not null default 0,
    unique(id_check, pay_year, pay_month)
);

--alter table staff add column fio_dative varchar(300) not null default '';
--alter table staff add column position_dative varchar(100) not null default '';

--alter table staff add column fio_genitive varchar(300) not null default '';
--alter table staff add column position_genitive varchar(100) not null default '';
