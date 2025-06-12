-- Inserta una escuela de ejemplo si no existe
INSERT INTO escuela (nombre, direccion, comuna)
VALUES ('Escuela Demo', 'Calle Falsa 123', 'Comuna Central')
ON CONFLICT (nombre) DO NOTHING;

-- Poblar las salas (usamos escuela_id = 1 directamente)
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
INSERT INTO umbrales (sala_id, variable_id, umbral_bajo, umbral_alto) VALUES
(1, 1, 20, 24), (1, 2, 40, 60), (1, 3, 400, 1000), (1, 4, 0, 35), (1, 5, 200, 800),
(2, 1, 20, 24), (2, 2, 40, 60), (2, 3, 400, 1000), (2, 4, 0, 35), (2, 5, 200, 800),
(3, 1, 20, 24), (3, 2, 40, 60), (3, 3, 400, 1000), (3, 4, 0, 35), (3, 5, 200, 800),
(4, 1, 20, 24), (4, 2, 40, 60), (4, 3, 400, 1000), (4, 4, 0, 35), (4, 5, 200, 800),
(5, 1, 20, 24), (5, 2, 40, 60), (5, 3, 400, 1000), (5, 4, 0, 35), (5, 5, 200, 800);

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

-- Insertar usuario normal
INSERT INTO usuario (rut_usuario, nombre, email, password, rol, escuela_id) VALUES 
('12345678-9', 'Usuario Demo', 'usuario@demo.cl', 'usuario123', 'user', 1)
ON CONFLICT (email) DO NOTHING;
