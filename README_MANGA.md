# MangaReader - Application Go

Application web complète de lecture de mangas en Go (Golang).

## 🚀 Fonctionnalités

- ✅ Page d'accueil avec Top Mangas et Recommandations
- ✅ Page de détail avec synopsis et liste des chapitres
- ✅ Lecteur (Reader) avec défilement vertical
- ✅ Dark Mode / Light Mode
- ✅ Responsive (Mobile, Tablette, Desktop)
- ✅ Raccourcis clavier
- ✅ Sauvegarde de progression de lecture
- ✅ Architecture MVC
- ✅ Aucune base de données (fichier JSON)

## 📁 Structure du projet

```
mangareader/
├── main.go                 # Serveur Go et routes
├── library.json           # Base de données JSON
├── go.mod                 # Dépendances Go
├── templates/
│   ├── home.html         # Page d'accueil
│   ├── detail.html       # Page de détail manga
│   └── reader.html       # Lecteur de chapitres
└── static/
    ├── css/
    │   └── style.css     # Styles CSS avec dark mode
    └── js/
        ├── app.js        # JavaScript global
        └── reader.js     # JavaScript du lecteur
```

## 🛠️ Installation

### Prérequis
- Go 1.21 ou supérieur

### Étapes

1. **Installer les dépendances Go:**
```bash
go mod download
```

2. **Lancer l'application:**
```bash
go run main.go
```

3. **Accéder à l'application:**
Ouvrez votre navigateur sur `http://localhost:8080`

## 📚 Structure du fichier library.json

Le fichier `library.json` contient tous les mangas avec la structure suivante:

```json
{
  "mangas": [
    {
      "id": "manga-id",
      "title": "Titre du Manga",
      "author": "Nom de l'auteur",
      "cover_url": "URL de la couverture",
      "description": "Synopsis du manga",
      "genres": ["Genre1", "Genre2"],
      "status": "En cours" ou "Terminé",
      "year": 2020,
      "rating": 4.5,
      "chapters": [
        {
          "number": 1,
          "title": "Titre du chapitre",
          "pages": [
            "url_image_page1.jpg",
            "url_image_page2.jpg"
          ]
        }
      ]
    }
  ]
}
```

## ⌨️ Raccourcis clavier

### Global
- `T` - Changer le thème (Dark/Light)
- `Esc` - Retour à la page précédente

### Dans le lecteur
- `↑` / `W` - Page précédente
- `↓` / `S` - Page suivante
- `←` - Chapitre précédent
- `→` - Chapitre suivant
- `Home` - Première page
- `End` - Dernière page

## 🎨 Personnalisation

### Ajouter des mangas

Modifiez le fichier `library.json` et ajoutez vos mangas en suivant la structure décrite ci-dessus.

### Modifier les couleurs

Les couleurs sont définies dans `static/css/style.css` via les variables CSS:

```css
:root {
    --accent-primary: #6366f1;
    --accent-secondary: #8b5cf6;
    /* ... */
}
```

## 🔧 Développement

### Compiler pour production

```bash
go build -o mangareader main.go
```

### Lancer le binaire

```bash
./mangareader
```

## 📱 Compatibilité

- ✅ Chrome, Firefox, Safari, Edge (dernières versions)
- ✅ iOS Safari
- ✅ Android Chrome
- ✅ Support tactile pour mobile

## 🚀 Déploiement

### Heroku

```bash
# Créer un Procfile
echo "web: ./mangareader" > Procfile

# Déployer
git push heroku main
```

### Docker

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o mangareader main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/mangareader .
COPY library.json .
COPY templates/ templates/
COPY static/ static/
EXPOSE 8080
CMD ["./mangareader"]
```

## 📄 Licence

MIT

## 🤝 Contribution

Les contributions sont les bienvenues ! N'hésitez pas à ouvrir une issue ou un pull request.

---

Fait avec ❤️ en Go
