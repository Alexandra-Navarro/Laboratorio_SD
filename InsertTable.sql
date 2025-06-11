-- Inserta una escuela de ejemplo si no existe
INSERT INTO escuela (nombre, direccion, comuna)
VALUES ('Escuela Demo', 'Calle Falsa 123', 'Comuna Central')
ON CONFLICT DO NOTHING;

-- Obtén el id de la escuela recién creada o existente
WITH ins AS (
  INSERT INTO escuela (nombre, direccion, comuna)
  VALUES ('Escuela Demo', 'Calle Falsa 123', 'Comuna Central')
  ON CONFLICT (nombre) DO NOTHING
  RETURNING id
)
SELECT id FROM ins
UNION
SELECT id FROM escuela WHERE nombre = 'Escuela Demo'
LIMIT 1;

-- Poblar las salas
INSERT INTO sala (nombre, escuela_id) VALUES
('A101', 1),
('A102', 1),
('A103', 1),
('B101', 1),
('B102', 1);

-- Insertar variables ambientales con sus umbrales
INSERT INTO variable_ambiental (nombre, unidad_medida, umbral_bajo, umbral_alto) VALUES
('Temperatura', '°C', 20, 24),
('Humedad', '%', 40, 60),
('CO2', 'ppm', 400, 1000),
('Ruido', 'dB[A]', 0, 35),
('Iluminación', 'lux', 200, 800);

-- Insertar sensores para cada sala
-- Sala A101
INSERT INTO sensor (modelo, estado, fecha_instalacion, sala_id, variable_id) VALUES
('DHT22', 'activo', CURRENT_DATE, 1, 1), -- Temperatura
('DHT22', 'activo', CURRENT_DATE, 1, 2), -- Humedad
('MH-Z19', 'activo', CURRENT_DATE, 1, 3), -- CO2
('KY-038', 'activo', CURRENT_DATE, 1, 4), -- Ruido
('LDR', 'activo', CURRENT_DATE, 1, 5); -- Iluminación

-- Sala A102
INSERT INTO sensor (modelo, estado, fecha_instalacion, sala_id, variable_id) VALUES
('DHT22', 'activo', CURRENT_DATE, 2, 1), -- Temperatura
('DHT22', 'activo', CURRENT_DATE, 2, 2), -- Humedad
('MH-Z19', 'activo', CURRENT_DATE, 2, 3), -- CO2
('KY-038', 'activo', CURRENT_DATE, 2, 4), -- Ruido
('LDR', 'activo', CURRENT_DATE, 2, 5); -- Iluminación

-- Sala A103
INSERT INTO sensor (modelo, estado, fecha_instalacion, sala_id, variable_id) VALUES
('DHT22', 'activo', CURRENT_DATE, 3, 1), -- Temperatura
('DHT22', 'activo', CURRENT_DATE, 3, 2), -- Humedad
('MH-Z19', 'activo', CURRENT_DATE, 3, 3), -- CO2
('KY-038', 'activo', CURRENT_DATE, 3, 4), -- Ruido
('LDR', 'activo', CURRENT_DATE, 3, 5); -- Iluminación

-- Sala B101
INSERT INTO sensor (modelo, estado, fecha_instalacion, sala_id, variable_id) VALUES
('DHT22', 'activo', CURRENT_DATE, 4, 1), -- Temperatura
('DHT22', 'activo', CURRENT_DATE, 4, 2), -- Humedad
('MH-Z19', 'activo', CURRENT_DATE, 4, 3), -- CO2
('KY-038', 'activo', CURRENT_DATE, 4, 4), -- Ruido
('LDR', 'activo', CURRENT_DATE, 4, 5); -- Iluminación

-- Sala B102
INSERT INTO sensor (modelo, estado, fecha_instalacion, sala_id, variable_id) VALUES
('DHT22', 'activo', CURRENT_DATE, 5, 1), -- Temperatura
('DHT22', 'activo', CURRENT_DATE, 5, 2), -- Humedad
('MH-Z19', 'activo', CURRENT_DATE, 5, 3), -- CO2
('KY-038', 'activo', CURRENT_DATE, 5, 4), -- Ruido
('LDR', 'activo', CURRENT_DATE, 5, 5); -- Iluminación