CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SEQUENCE my_uuid_sequence;

CREATE OR REPLACE FUNCTION uuid_generator()
    RETURNS uuid
AS
$$
BEGIN
RETURN uuid_generate_v4();
END;
$$ LANGUAGE plpgsql;