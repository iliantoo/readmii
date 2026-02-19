package main

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Global library
var library Library

// Template cache
var templates *template.Template

// Database connection
var db *sql.DB
var useDatabase bool

// Init database connection
func initDatabase() error {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Println("⚠️  DATABASE_URL non définie, utilisation du fichier JSON")
		useDatabase = false
		return nil
	}

	var err error
	db, err = sql.Open("postgres", databaseURL)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		log.Printf("⚠️  Erreur connexion DB: %v. Utilisation du fichier JSON\n", err)
		useDatabase = false
		return nil
	}

	useDatabase = true
	log.Println("✅ Connexion PostgreSQL réussie")
	return nil
}

// Load maps from database
func loadMapsFromDB() ([]Map, error) {
	rows, err := db.Query(`
		SELECT id, name, display_name, cover_image, total_lineups
		FROM maps
		ORDER BY name
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var maps []Map
	for rows.Next() {
		var m Map
		err := rows.Scan(&m.ID, &m.Name, &m.DisplayName, &m.CoverImage, &m.TotalLineups)
		if err != nil {
			return nil, err
		}
		maps = append(maps, m)
	}

	return maps, nil
}

// Get map with lineups from database
func getMapFromDB(mapID string) (*Map, error) {
	var m Map
	err := db.QueryRow(`
		SELECT id, name, display_name, cover_image, total_lineups
		FROM maps WHERE id = $1
	`, mapID).Scan(&m.ID, &m.Name, &m.DisplayName, &m.CoverImage, &m.TotalLineups)

	if err != nil {
		return nil, err
	}

	// Charger les lineups
	rows, err := db.Query(`
		SELECT id, title, description, grenade_type, side, difficulty,
			throw_zone, landing_zone, action_type,
			bind_required, bind_command, movement, click_type,
			position_image, position_thumbnail, crosshair_image, crosshair_zoom_level,
			demo_gif, demo_thumbnail,
			popularity, views
		FROM lineups
		WHERE map_id = $1
		ORDER BY popularity DESC, views DESC
	`, mapID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lineups []Lineup
	for rows.Next() {
		var l Lineup
		err := rows.Scan(
			&l.ID, &l.Title, &l.Description, &l.GrenadeType, &l.Side, &l.Difficulty,
			&l.ThrowZone, &l.LandingZone, &l.ActionType,
			&l.ActionDetails.BindRequired, &l.ActionDetails.BindCommand,
			&l.ActionDetails.Movement, &l.ActionDetails.ClickType,
			&l.Media.PositionImage, &l.Media.PositionThumbnail,
			&l.Media.CrosshairImage, &l.Media.CrosshairZoomLevel,
			&l.Media.DemoGif, &l.Media.DemoThumbnail,
			&l.Popularity, &l.Views,
		)
		if err != nil {
			return nil, err
		}

		// Charger les tags
		tagRows, err := db.Query(`
			SELECT t.name
			FROM tags t
			JOIN lineup_tags lt ON t.id = lt.tag_id
			WHERE lt.lineup_id = $1
		`, l.ID)

		if err == nil {
			var tags []string
			for tagRows.Next() {
				var tag string
				if err := tagRows.Scan(&tag); err == nil {
					tags = append(tags, tag)
				}
			}
			tagRows.Close()
			l.Tags = tags
		}

		lineups = append(lineups, l)
	}

	m.Lineups = lineups
	return &m, nil
}

// Load library from JSON file
func loadLibrary() error {
	data, err := ioutil.ReadFile("lineups.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &library)
	if err != nil {
		return err
	}

	log.Printf("✅ Bibliothèque chargée : %d cartes", len(library.Maps))
	return nil
}

// Load templates
func loadTemplates() {
	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	}

	templates = template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*.html"))
	log.Println("✅ Templates chargés")
}

// Helper functions
func getMapByID(id string) *Map {
	if useDatabase {
		m, err := getMapFromDB(id)
		if err != nil {
			log.Printf("Erreur récupération map depuis DB: %v\n", err)
			return nil
		}
		return m
	}

	// Fallback sur JSON
	for i := range library.Maps {
		if library.Maps[i].ID == id {
			return &library.Maps[i]
		}
	}
	return nil
}

func getAllMaps() []Map {
	if useDatabase {
		maps, err := loadMapsFromDB()
		if err != nil {
			log.Printf("Erreur récupération maps depuis DB: %v\n", err)
			return library.Maps
		}
		return maps
	}

	return library.Maps
}

// Handlers
func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Maps      []Map
		PageTitle string
	}{
		Maps:      getAllMaps(),
		PageTitle: "CS2 Lineups - Home",
	}

	err := templates.ExecuteTemplate(w, "home.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Erreur template home:", err)
	}
}

func mapViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mapID := vars["id"]

	mapData := getMapByID(mapID)
	if mapData == nil {
		http.NotFound(w, r)
		return
	}

	data := struct {
		Map       *Map
		PageTitle string
	}{
		Map:       mapData,
		PageTitle: mapData.Name + " - Lineups",
	}

	err := templates.ExecuteTemplate(w, "map-view.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Erreur template map-view:", err)
	}
}

// Main
func main() {
	// Init database connection
	if err := initDatabase(); err != nil {
		log.Printf("⚠️  Erreur initialisation DB: %v\n", err)
	}

	// Load data from JSON as fallback
	if !useDatabase {
		if err := loadLibrary(); err != nil {
			log.Fatal("❌ Erreur de chargement de lineups.json:", err)
		}
	}

	// Load templates
	loadTemplates()

	// Router
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/map/{id}", mapViewHandler).Methods("GET")

	// Static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("public/assets"))))

	// Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if useDatabase {
		log.Printf("🚀 Serveur CS2 Lineups démarré sur http://localhost:%s (PostgreSQL)\n", port)
	} else {
		log.Printf("🚀 Serveur CS2 Lineups démarré sur http://localhost:%s (JSON)\n", port)
	}

	log.Fatal(http.ListenAndServe(":"+port, r))
}
