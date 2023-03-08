CREATE TABLE IF NOT EXISTS public.meters
(
    id                UUID PRIMARY KEY                          NOT NULL DEFAULT uuid_generator(),
    brand             VARCHAR                                   NOT NULL,
    "serial"          VARCHAR                                   NOT NULL,
    in_use            BOOLEAN                                   NOT NULL DEFAULT FALSE,
    last_installation TIMESTAMP                                 NULL,
    lines             INT CHECK ( lines >= 0 AND lines <= 10)   NOT NULL,
    created_at        TIMESTAMP                                 NOT NULL DEFAULT NOW()
);

INSERT INTO public.meters (id,brand, serial, in_use, last_installation, lines)
VALUES ('2b2911b6-3c59-4a28-9379-c75e82278960','ETELCA', 'A1001', TRUE, '2023-02-01 00:00:00.000000', 1),
       ('6e033a13-3c1b-435d-a891-3fa0ace4d047','ETELCA', 'A2001', TRUE, '2022-03-10 00:00:00.000000', 1),
       ('cdae884e-ffb6-4c50-93f0-f4c0bdfdea0d','ETELCA', 'A3001', FALSE, '2023-01-10 00:00:00.000000', 1),
       ('f77bc602-19d6-4f06-902e-83f0f0c64b3e','TRANSCOR', '132ABE', TRUE, '2020-10-25 00:00:00.000000', 1),
       ('d4063676-cf0d-4ef7-aefb-15cff2ef0ff5','TRANSCOR', '142FRT', FALSE, NULL, 1);