package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Models
type Chapter struct {
	Number int      `json:"number"`
	Title  string   `json:"title"`
	Pages  []string `json:"pages"`
}

type Manga struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	CoverURL    string    `json:"cover_url"`
	Description string    `json:"description"`
	Genres      []string  `json:"genres"`
	Status      string    `json:"status"`
	Year        int       `json:"year"`
	Rating      float64   `json:"rating"`
	Chapters    []Chapter `json:"chapters"`
}

type Library struct {
	Mangas []Manga `json:"mangas"`
}

// Global library
var library Library

// Template cache
var templates *template.Template

// Initialize
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Load library from JSON file
func loadLibrary() error {
	data, err := ioutil.ReadFile("library.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &library)
	if err != nil {
		return err
	}

	log.Printf("✅ Bibliothèque chargée : %d mangas", len(library.Mangas))
	return nil
}

// Load templates
func loadTemplates() {
	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	}

	var err error
	templates = template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*.html"))
	if err != nil {
		log.Fatal("Erreur de chargement des templates:", err)
	}
	log.Println("✅ Templates chargés")
}

// Helper functions
func getMangaByID(id string) *Manga {
	for i := range library.Mangas {
		if library.Mangas[i].ID == id {
			return &library.Mangas[i]
		}
	}
	return nil
}

func getTopMangas(limit int) []Manga {
	if limit > len(library.Mangas) {
		limit = len(library.Mangas)
	}
	return library.Mangas[:limit]
}

func getRandomMangas(count int) []Manga {
	if count > len(library.Mangas) {
		count = len(library.Mangas)
	}

	indices := rand.Perm(len(library.Mangas))
	result := make([]Manga, count)
	for i := 0; i < count; i++ {
		result[i] = library.Mangas[indices[i]]
	}
	return result
}

// Handlers
func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		TopMangas   []Manga
		Recommended []Manga
		AllMangas   []Manga
		PageTitle   string
	}{
		TopMangas:   getTopMangas(5),
		Recommended: getRandomMangas(4),
		AllMangas:   library.Mangas,
		PageTitle:   "Readmii - Accueil",
	}

	err := templates.ExecuteTemplate(w, "home.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Erreur template home:", err)
	}
}

func mangaDetailHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mangaID := vars["id"]

	manga := getMangaByID(mangaID)
	if manga == nil {
		http.NotFound(w, r)
		return
	}

	data := struct {
		Manga     *Manga
		PageTitle string
	}{
		Manga:     manga,
		PageTitle: manga.Title + " - Détails",
	}

	err := templates.ExecuteTemplate(w, "detail.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Erreur template detail:", err)
	}
}

func readerHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	mangaID := vars["id"]
	chapterNum, err := strconv.Atoi(vars["chapter"])
	if err != nil {
		http.Error(w, "Numéro de chapitre invalide", http.StatusBadRequest)
		return
	}

	manga := getMangaByID(mangaID)
	if manga == nil {
		http.NotFound(w, r)
		return
	}

	var chapter *Chapter
	for i := range manga.Chapters {
		if manga.Chapters[i].Number == chapterNum {
			chapter = &manga.Chapters[i]
			break
		}
	}

	if chapter == nil {
		http.NotFound(w, r)
		return
	}

	// Find previous and next chapters
	var prevChapter, nextChapter *int
	for i := range manga.Chapters {
		if manga.Chapters[i].Number == chapterNum-1 {
			num := manga.Chapters[i].Number
			prevChapter = &num
		}
		if manga.Chapters[i].Number == chapterNum+1 {
			num := manga.Chapters[i].Number
			nextChapter = &num
		}
	}

	data := struct {
		Manga       *Manga
		Chapter     *Chapter
		PrevChapter *int
		NextChapter *int
		PageTitle   string
	}{
		Manga:       manga,
		Chapter:     chapter,
		PrevChapter: prevChapter,
		NextChapter: nextChapter,
		PageTitle:   manga.Title + " - Chapitre " + strconv.Itoa(chapterNum),
	}

	err = templates.ExecuteTemplate(w, "reader.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Erreur template reader:", err)
	}
}

// Main
func main() {
	// Load data
	if err := loadLibrary(); err != nil {
		log.Fatal("❌ Erreur de chargement de library.json:", err)
	}

	// Load templates
	loadTemplates()

	// Router
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/manga/{id}", mangaDetailHandler).Methods("GET")
	r.HandleFunc("/read/{id}/{chapter:[0-9]+}", readerHandler).Methods("GET")

	// Static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Server
	port := "8080"
	log.Printf("🚀 Serveur démarré sur http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
