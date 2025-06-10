-- Inserta una escuela de ejemplo si no existe
INSERT INTO escuela (nombre, direccion, comuna)
VALUES ('Escuela Demo', 'Calle Falsa 123', 'Comuna Central')
ON CONFLICT DO NOTHING;

-- Obtén el id de la escuela recién creada o existente
-- (esto es para PostgreSQL 9.5+)
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

-- Supón que el id es 1, ajusta si es diferente
-- Poblar las salas
INSERT INTO sala (nombre, escuela_id) VALUES
('A101', 1),
('A102', 1),
('A103', 1),
('B101', 1),
('B102', 1);

INSERT INTO variable_ambiental (nombre, unidad_medida, umbral_bajo, umbral_alto) VALUES
('Temperatura', '°C', 18, 28),
('Humedad', '%', 30, 70),
('CO2', 'ppm', 400, 1000),
('Ruido', 'dB', 30, 85),
('Iluminación', 'lux', 200, 800);

INSERT INTO sensor (modelo, estado, fecha_instalacion, sala_id, variable_id) VALUES
('ESP32', 'activo', CURRENT_DATE, 1, 1), -- Sala 1, Temperatura
('ESP32', 'activo', CURRENT_DATE, 1, 2), -- Sala 1, Humedad
('ESP32', 'activo', CURRENT_DATE, 1, 3), -- Sala 1, CO2
('ESP32', 'activo', CURRENT_DATE, 1, 4), -- Sala 1, Ruido
('ESP32', 'activo', CURRENT_DATE, 1, 5), -- Sala 1, Iluminación

('ESP32', 'activo', CURRENT_DATE, 2, 1),
('ESP32', 'activo', CURRENT_DATE, 2, 2),
('ESP32', 'activo', CURRENT_DATE, 2, 3),
('ESP32', 'activo', CURRENT_DATE, 2, 4),
('ESP32', 'activo', CURRENT_DATE, 2, 5),

('ESP32', 'activo', CURRENT_DATE, 3, 1),
('ESP32', 'activo', CURRENT_DATE, 3, 2),
('ESP32', 'activo', CURRENT_DATE, 3, 3),
('ESP32', 'activo', CURRENT_DATE, 3, 4),
('ESP32', 'activo', CURRENT_DATE, 3, 5),

('ESP32', 'activo', CURRENT_DATE, 4, 1),
('ESP32', 'activo', CURRENT_DATE, 4, 2),
('ESP32', 'activo', CURRENT_DATE, 4, 3),
('ESP32', 'activo', CURRENT_DATE, 4, 4),
('ESP32', 'activo', CURRENT_DATE, 4, 5),

('ESP32', 'activo', CURRENT_DATE, 5, 1),
('ESP32', 'activo', CURRENT_DATE, 5, 2),
('ESP32', 'activo', CURRENT_DATE, 5, 3),
('ESP32', 'activo', CURRENT_DATE, 5, 4),
('ESP32', 'activo', CURRENT_DATE, 5, 5);