# Data Models
 
## Overview
 
The application uses PostgreSQL 16 as its relational database. The schema is defined in `resources/schema.sql` and contains three tables: `cars`, `users`, and `user_favorites`. These three tables cover everything the application needs to store: the vehicle catalogue, the user accounts, and the relationship between users and their saved vehicles.

## The cars table

Stores the comprehensive vehicle technical telemetry dataset.
- `id`: `SERIAL PRIMARY KEY`
- `company`: `VARCHAR(100)` (Car manufacturer brand)
- `car_name`: `VARCHAR(100)` (Specific model label)
- `engine`: `VARCHAR(100)`
- `capacity_cc`: `INTEGER`
- `power_hp`: `INTEGER`
- `max_speed_kmh`: `INTEGER`
- `acceleration_0_100_sec`: `FLOAT`
- `price`: `FLOAT`
- `fuel_type`: `VARCHAR(50)`
- `seats`: `INTEGER`
- `torque_nm`: `INTEGER`
In the Go domain layer, the `Car` struct in `internal/core/domain/car.go` maps to this table. Some fields like `CapacityCC` and `TorqueNM` use pointer types (`*int`) to correctly handle `NULL` values that appear in the raw dataset when certain specs are not available for a vehicle. The database column `company` maps to the `Brand` field and `car_name` maps to `Model` in the Go struct, handled through `json` struct tags.

## The users table

Stores authentication credentials and roles.
- `id`: `SERIAL PRIMARY KEY`
- `username`: `VARCHAR(100) UNIQUE NOT NULL`
- `email`: `VARCHAR(150) UNIQUE NOT NULL`
- `password`: `VARCHAR(255) NOT NULL` (Secured cryptographically hashed string)
- `role`: `VARCHAR(20) DEFAULT 'user'` ('user' or 'admin')
- `created_at`: `TIMESTAMP DEFAULT CURRENT_TIMESTAMP`

The `users` table stores the registered user accounts. It holds the `username`, which must be unique, the `email` used to log in, which is also unique, the `password` field which stores the Bcrypt hash of the original password, a `role` field that defaults to `user` and can be `admin`, and a `created_at` timestamp that records when the account was created.
 
Passwords are never stored in plain text. When a user registers, `auth_service.go` runs the password through Bcrypt before inserting it into the database. Bcrypt generates a unique random salt for each hash, which means two users with identical passwords will have completely different values stored in the database. On login, `bcrypt.CompareHashAndPassword` verifies the submitted password against the stored hash without ever reversing the hash.
 
In the Go domain layer, the `User` struct uses the tag `json:"-"` on the `Password` field to ensure the hash is never included in any JSON API response, regardless of the context.

## The user_favorites table

The `user_favorites` table is a junction table that resolves the many-to-many relationship between users and cars. A user can save as many cars as they want, and the same car can be saved by multiple users. Each row simply stores a `user_id` referencing the `users` table and a `car_id` referencing the `cars` table. The primary key is a composite of both columns, which prevents a user from saving the same car twice. Both foreign keys use `ON DELETE CASCADE`, so if a user or a car is deleted, the associated favorites are automatically removed.

- `user_id`: `INTEGER REFERENCES users(id) ON DELETE CASCADE`
- `car_id`: `INTEGER REFERENCES cars(id) ON DELETE CASCADE`
- *Constraint*: Composite Primary Key `PRIMARY KEY (user_id, car_id)`.
