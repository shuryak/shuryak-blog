ALTER TABLE users
ADD role TEXT CONSTRAINT role_constraint CHECK (role = 'user' OR role = 'admin');
