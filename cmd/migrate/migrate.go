package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// Connexion à PostgreSQL
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		// Pour le développement local
		databaseURL = "postgres://localhost/cs2_lineups?sslmode=disable"
	}

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("Erreur connexion DB:", err)
	}
	defer db.Close()

	// Test de connexion
	if err = db.Ping(); err != nil {
		log.Fatal("Erreur ping DB:", err)
	}

	log.Println("✅ Connexion à la base de données réussie")

	// Charger le fichier JSON
	data, err := ioutil.ReadFile("/tmp/uploads/data.json")
	if err != nil {
		log.Fatal("Erreur lecture lineups.json:", err)
	}

	var library Library
	if err = json.Unmarshal(data, &library); err != nil {
		log.Fatal("Erreur parsing JSON:", err)
	}

	log.Printf("📦 %d cartes trouvées dans le JSON\n", len(library.Maps))

	// Migrer les données
	for _, m := range library.Maps {
		// Insérer la carte
		_, err = db.Exec(`
			INSERT INTO maps (id, name, display_name, cover_image, total_lineups)
			VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT (id) DO UPDATE SET
				name = $2,
				display_name = $3,
				cover_image = $4,
				total_lineups = $5
		`, m.ID, m.Name, m.DisplayName, m.CoverImage, m.TotalLineups)

		if err != nil {
			log.Printf("❌ Erreur insertion map %s: %v\n", m.ID, err)
			continue
		}

		log.Printf("✅ Map %s insérée\n", m.Name)

		// Insérer les lineups
		for _, lineup := range m.Lineups {
			_, err = db.Exec(`
				INSERT INTO lineups (
					id, map_id, title, description, grenade_type, side, difficulty,
					throw_zone, landing_zone, action_type,
					bind_required, bind_command, movement, click_type,
					position_image, position_thumbnail, crosshair_image, crosshair_zoom_level,
					demo_gif, demo_thumbnail,
					popularity, views
				) VALUES (
					$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
					$11, $12, $13, $14,
					$15, $16, $17, $18, $19, $20,
					$21, $22
				)
				ON CONFLICT (id) DO UPDATE SET
					title = $3,
					description = $4,
					grenade_type = $5,
					side = $6,
					difficulty = $7,
					throw_zone = $8,
					landing_zone = $9,
					action_type = $10,
					bind_required = $11,
					bind_command = $12,
					movement = $13,
					click_type = $14,
					position_image = $15,
					position_thumbnail = $16,
					crosshair_image = $17,
					crosshair_zoom_level = $18,
					demo_gif = $19,
					demo_thumbnail = $20,
					popularity = $21,
					views = $22,
					updated_at = CURRENT_TIMESTAMP
			`,
				lineup.ID, m.ID, lineup.Title, lineup.Description,
				lineup.GrenadeType, lineup.Side, lineup.Difficulty,
				lineup.ThrowZone, lineup.LandingZone, lineup.ActionType,
				lineup.ActionDetails.BindRequired, lineup.ActionDetails.BindCommand,
				lineup.ActionDetails.Movement, lineup.ActionDetails.ClickType,
				lineup.Media.PositionImage, lineup.Media.PositionThumbnail,
				lineup.Media.CrosshairImage, lineup.Media.CrosshairZoomLevel,
				lineup.Media.DemoGif, lineup.Media.DemoThumbnail,
				lineup.Popularity, lineup.Views,
			)

			if err != nil {
				log.Printf("❌ Erreur insertion lineup %s: %v\n", lineup.ID, err)
				continue
			}

			// Insérer les tags
			for _, tagName := range lineup.Tags {
				// Insérer le tag s'il n'existe pas
				var tagID int
				err = db.QueryRow(`
					INSERT INTO tags (name) VALUES ($1)
					ON CONFLICT (name) DO UPDATE SET name = $1
					RETURNING id
				`, tagName).Scan(&tagID)

				if err != nil {
					log.Printf("❌ Erreur insertion tag %s: %v\n", tagName, err)
					continue
				}

				// Lier le tag au lineup
				_, err = db.Exec(`
					INSERT INTO lineup_tags (lineup_id, tag_id)
					VALUES ($1, $2)
					ON CONFLICT DO NOTHING
				`, lineup.ID, tagID)

				if err != nil {
					log.Printf("❌ Erreur liaison tag: %v\n", err)
				}
			}

			log.Printf("  ✅ Lineup %s inséré\n", lineup.Title)
		}
	}

	log.Println("🎉 Migration terminée avec succès!")
}
