DROP TABLE IF EXISTS cars;
CREATE TABLE cars (
    id                     SERIAL PRIMARY KEY,
    company                VARCHAR(100),
    car_name               VARCHAR(100),
    engine                 VARCHAR(100),
    capacity_cc            INTEGER,
    power_hp               INTEGER,
    max_speed_kmh          INTEGER,
    acceleration_0_100_sec FLOAT,
    price                  FLOAT,
    fuel_type              VARCHAR(50),
    seats                  INTEGER,
    torque_nm              INTEGER
);

CREATE TABLE IF NOT EXISTS users (
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(100) UNIQUE NOT NULL,
    email      VARCHAR(150) UNIQUE NOT NULL,
    password   VARCHAR(255) NOT NULL,
    role       VARCHAR(20) DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


