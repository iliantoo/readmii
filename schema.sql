-- Schema pour CS2 Lineups Database

-- Table des cartes
CREATE TABLE IF NOT EXISTS maps (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    display_name VARCHAR(100) NOT NULL,
    cover_image VARCHAR(255),
    total_lineups INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table des lineups
CREATE TABLE IF NOT EXISTS lineups (
    id VARCHAR(100) PRIMARY KEY,
    map_id VARCHAR(50) REFERENCES maps(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    grenade_type VARCHAR(20) NOT NULL,
    side VARCHAR(10) NOT NULL,
    difficulty VARCHAR(20) NOT NULL,
    throw_zone VARCHAR(100) NOT NULL,
    landing_zone VARCHAR(100) NOT NULL,
    action_type VARCHAR(50) NOT NULL,
    
    -- Action details (JSON ou colonnes separées)
    bind_required BOOLEAN DEFAULT FALSE,
    bind_command TEXT,
    movement VARCHAR(50),
    click_type VARCHAR(50),
    
    -- Media (chemins vers les fichiers)
    position_image VARCHAR(255),
    position_thumbnail VARCHAR(255),
    crosshair_image VARCHAR(255),
    crosshair_zoom_level VARCHAR(10),
    demo_gif VARCHAR(255),
    demo_thumbnail VARCHAR(255),
    
    -- Stats
    popularity DECIMAL(3,1) DEFAULT 0.0,
    views INTEGER DEFAULT 0,
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table des tags
CREATE TABLE IF NOT EXISTS tags (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL
);

-- Table de liaison lineup-tags
CREATE TABLE IF NOT EXISTS lineup_tags (
    lineup_id VARCHAR(100) REFERENCES lineups(id) ON DELETE CASCADE,
    tag_id INTEGER REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY (lineup_id, tag_id)
);

-- Index pour améliorer les performances
CREATE INDEX IF NOT EXISTS idx_lineups_map_id ON lineups(map_id);
CREATE INDEX IF NOT EXISTS idx_lineups_grenade_type ON lineups(grenade_type);
CREATE INDEX IF NOT EXISTS idx_lineups_side ON lineups(side);
CREATE INDEX IF NOT EXISTS idx_lineups_difficulty ON lineups(difficulty);
CREATE INDEX IF NOT EXISTS idx_lineups_popularity ON lineups(popularity DESC);
CREATE INDEX IF NOT EXISTS idx_lineups_views ON lineups(views DESC);
