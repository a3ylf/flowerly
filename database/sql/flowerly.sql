-- Cria um novo superusuário
CREATE ROLE flowermaster WITH LOGIN SUPERUSER PASSWORD '123456';

CREATE TABLE plants (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    scientific_name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    category VARCHAR(100) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    stock_quantity INTEGER NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    origin_location VARCHAR(255) NOT NULL
);

INSERT INTO plants (name, scientific_name, description, category, price, stock_quantity, image_url, origin_location)
VALUES 
('Rosa Vermelha', 'Rosa rubiginosa', 'Rosas vermelhas de alta qualidade, perfeitas para buquês.', 'Flores', 5.99, 150, 'rosa_vermelha.png', 'Areia'),
('Orquídea Branca', 'Phalaenopsis amabilis', 'Orquídea Phalaenopsis branca, ideal para presentes.', 'Plantas', 25.00, 60, 'orquidea_branca.png', 'Bananeiras'),
('Bromélia', 'Bromelia laciniosa', 'Planta tropical resistente, comum em regiões áridas.', 'Plantas', 18.50, 85, 'bromelia.png', 'Campina Grande'),
('Hibisco', 'Hibiscus rosa-sinensis', 'Flores vibrantes usadas em paisagismo e chás.', 'Flores', 7.99, 200, 'hibisco.png', 'João Pessoa'),
('Cacto Mandacaru', 'Cereus jamacaru', 'Cacto típico do semiárido nordestino.', 'Suculentas', 12.00, 100, 'mandacaru.png', 'Patos'),
('Palmeira Imperial', 'Roystonea oleracea', 'Palmeira de grande porte, muito usada em jardins.', 'Árvores', 120.00, 15, 'palmeira_imperial.png', 'Areia'),
('Jiboia', 'Epipremnum aureum', 'Planta trepadeira muito popular em interiores.', 'Plantas', 10.99, 75, 'jiboia.png', 'Cajazeiras'),
('Samambaia', 'Nephrolepis exaltata', 'Planta ornamental muito utilizada em varandas.', 'Plantas', 14.50, 95, 'samambaia.png', 'Mari'),
('Ipê Amarelo', 'Handroanthus albus', 'Árvore ornamental com flores amarelas vibrantes.', 'Árvores', 85.00, 40, 'ipe_amarelo.png', 'Areia'),
('Café', 'Coffea arabica', 'Planta cultivada para a produção de café.', 'Árvores', 30.00, 120, 'cafe.png', 'Brejo da Paraíba'),
('Flor de Maio', 'Schlumbergera truncata', 'Planta suculenta com flores coloridas.', 'Suculentas', 15.00, 60, 'flor_de_maio.png', 'Mari'),
('Maranta', 'Maranta leuconeura', 'Planta de folhas ornamentais, também conhecida como planta rezadeira.', 'Plantas', 20.00, 50, 'maranta.png', 'Guarabira'),
('Costela de Adão', 'Monstera deliciosa', 'Planta de folhas grandes e recortadas, muito usada em decoração.', 'Plantas', 35.00, 40, 'costela_de_adao.png', 'Campina Grande'),
('Violeta', 'Saintpaulia ionantha', 'Planta de pequenas flores, popular em interiores.', 'Flores', 8.50, 180, 'violeta.png', 'Alagoa Grande'),
('Aroeira', 'Schinus terebinthifolius', 'Árvore com frutos usados como tempero.', 'Árvores', 65.00, 30, 'aroeira.png', 'Mari'),
('Antúrio', 'Anthurium andraeanum', 'Planta com inflorescência vermelha, muito usada em decoração.', 'Flores', 22.00, 70, 'anturio.png', 'Alagoa Nova'),
('Espada de São Jorge', 'Dracaena trifasciata', 'Planta resistente e fácil de cuidar, muito usada em proteção espiritual.', 'Plantas', 18.00, 80, 'espada_de_sao_jorge.png', 'Santa Rita'),
('Manacá da Serra', 'Tibouchina mutabilis', 'Árvore ornamental com flores roxas e brancas.', 'Árvores', 90.00, 20, 'manaca_da_serra.png', 'Esperança'),
('Suculenta Echeveria', 'Echeveria elegans', 'Planta ornamental de fácil cuidado.', 'Suculentas', 12.00, 150, 'echeveria.png', 'Bananeiras'),
('Jacarandá', 'Jacaranda mimosifolia', 'Árvore ornamental com flores roxas.', 'Árvores', 110.00, 25, 'jacaranda.png', 'Solânea');


-- Concede permissões ao novo usuário
GRANT ALL PRIVILEGES ON DATABASE plants TO flowermaster;
