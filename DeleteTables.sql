-- Desactivar temporalmente las restricciones de clave foránea
SET CONSTRAINTS ALL DEFERRED;

-- Borrar el contenido de las tablas en orden inverso a sus dependencias
DELETE FROM alerta;
DELETE FROM medicion;
DELETE FROM sensor;
DELETE FROM variable_ambiental;
DELETE FROM sala;
DELETE FROM escuela;

-- Reactivar las restricciones de clave foránea
SET CONSTRAINTS ALL IMMEDIATE;

-- Reiniciar las secuencias de ID
ALTER SEQUENCE alerta_id_seq RESTART WITH 1;
ALTER SEQUENCE medicion_id_seq RESTART WITH 1;
ALTER SEQUENCE sensor_id_seq RESTART WITH 1;
ALTER SEQUENCE variable_ambiental_id_seq RESTART WITH 1;
ALTER SEQUENCE sala_id_seq RESTART WITH 1;
ALTER SEQUENCE escuela_id_seq RESTART WITH 1; 