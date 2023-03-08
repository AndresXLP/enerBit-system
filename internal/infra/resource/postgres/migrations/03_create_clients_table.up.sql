CREATE TABLE IF NOT EXISTS public.clients
(
    id
    SERIAL
    PRIMARY
    KEY
    UNIQUE
    NOT
    NULL,
    meter_id
    UUID
    NOT
    NULL,
    address
    VARCHAR
    NOT
    NULL,
    installation_date
    TIMESTAMP
    NOT
    NULL,
    retirement_date
    TIMESTAMP
    NULL,
    is_active
    BOOLEAN
    NOT
    NULL
);

ALTER TABLE clients
    ADD CONSTRAINT clients_meter_id_fkey
        FOREIGN KEY (meter_id)
            REFERENCES meters (id)
            ON DELETE CASCADE;

INSERT INTO public.clients (meter_id, address, installation_date, retirement_date, is_active)
VALUES ('2b2911b6-3c59-4a28-9379-c75e82278960', 'CALLE 53 # 9D - 75', '2023-02-01 00:00:00.000000', NULL, TRUE),
       ('6e033a13-3c1b-435d-a891-3fa0ace4d047', 'CARRERA 9D # 35A - 76', '2022-03-10 00:00:00.000000', NULL, TRUE),
       ('cdae884e-ffb6-4c50-93f0-f4c0bdfdea0d', 'CALLE 15 # 20 - 40', '2023-01-10 00:00:00.000000', '2023-03-10 00:00:00.000000', FALSE),
       ('f77bc602-19d6-4f06-902e-83f0f0c64b3e', 'CALLE 17 # 17B - 19', '2020-10-25 00:00:00.000000', NULL, FALSE);