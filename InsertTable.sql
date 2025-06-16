-- Inserta una escuela de ejemplo
INSERT INTO escuela (nombre, direccion, comuna)
VALUES ('Escuela Demo', 'Calle Falsa 123', 'Comuna Central')
ON CONFLICT (nombre) DO NOTHING;

-- Poblar las salas
INSERT INTO sala (nombre, escuela_id) VALUES
('A101', 1),
('A102', 1),
('A103', 1),
('B101', 1),
('B102', 1);

-- Insertar variables ambientales
INSERT INTO variable_ambiental (nombre, unidad_medida) VALUES
('Temperatura', '°C'),
('Humedad', '%'),
('CO2', 'ppm'),
('Ruido', 'dB[A]'),
('Iluminación', 'lux');

-- Insertar umbrales personalizados por sala y variable
INSERT INTO umbrales (sala_id, variable_id, umbral_bajo, umbral_alto, umbral_bajo_prev, umbral_alto_prev) VALUES
(1, 1, 20, 24, 27, 25), (1, 2, 30, 70, 39, 61), (1, 3, 400, 1500, 1500, 1001), (1, 4, 0, 55, 55, 35), (1, 5, 350, 1000, 500, 751),
(2, 1, 20, 24, 27, 25), (2, 2, 30, 70, 39, 61), (2, 3, 400, 1500, 1500, 1001), (2, 4, 0, 55, 55, 35), (2, 5, 350, 1000, 500, 751),
(3, 1, 20, 24, 27, 25), (3, 2, 30, 70, 39, 61), (3, 3, 400, 1500, 1500, 1001), (3, 4, 0, 55, 55, 35), (3, 5, 350, 1000, 500, 751),
(4, 1, 20, 24, 27, 25), (4, 2, 30, 70, 39, 61), (4, 3, 400, 1500, 1500, 1001), (4, 4, 0, 55, 55, 35), (4, 5, 350, 1000, 500, 751),
(5, 1, 20, 24, 27, 25), (5, 2, 30, 70, 39, 61), (5, 3, 400, 1500, 1500, 1001), (5, 4, 0, 55, 55, 35), (5, 5, 350, 1000, 500, 751);

-- Insertar sensores por sala
INSERT INTO sensor (modelo, estado, fecha_instalacion, sala_id, variable_id) VALUES
-- Sala A101
('DHT22', 'activo', CURRENT_DATE, 1, 1),
('DHT22', 'activo', CURRENT_DATE, 1, 2),
('MH-Z19', 'activo', CURRENT_DATE, 1, 3),
('KY-038', 'activo', CURRENT_DATE, 1, 4),
('LDR', 'activo', CURRENT_DATE, 1, 5),
-- Sala A102
('DHT22', 'activo', CURRENT_DATE, 2, 1),
('DHT22', 'activo', CURRENT_DATE, 2, 2),
('MH-Z19', 'activo', CURRENT_DATE, 2, 3),
('KY-038', 'activo', CURRENT_DATE, 2, 4),
('LDR', 'activo', CURRENT_DATE, 2, 5),
-- Sala A103
('DHT22', 'activo', CURRENT_DATE, 3, 1),
('DHT22', 'activo', CURRENT_DATE, 3, 2),
('MH-Z19', 'activo', CURRENT_DATE, 3, 3),
('KY-038', 'activo', CURRENT_DATE, 3, 4),
('LDR', 'activo', CURRENT_DATE, 3, 5),
-- Sala B101
('DHT22', 'activo', CURRENT_DATE, 4, 1),
('DHT22', 'activo', CURRENT_DATE, 4, 2),
('MH-Z19', 'activo', CURRENT_DATE, 4, 3),
('KY-038', 'activo', CURRENT_DATE, 4, 4),
('LDR', 'activo', CURRENT_DATE, 4, 5),
-- Sala B102
('DHT22', 'activo', CURRENT_DATE, 5, 1),
('DHT22', 'activo', CURRENT_DATE, 5, 2),
('MH-Z19', 'activo', CURRENT_DATE, 5, 3),
('KY-038', 'activo', CURRENT_DATE, 5, 4),
('LDR', 'activo', CURRENT_DATE, 5, 5);

-- Insertar usuario
INSERT INTO usuario (rut_usuario, nombre, email, password, rol, escuela_id) VALUES 
('12345678-9', 'Usuario Demo', 'usuario@demo.cl', 'usuario123', 'admin', 1)
ON CONFLICT (email) DO NOTHING;
