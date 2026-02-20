# 🎮 CS2 Lineups - Site Professionnel de Lineups Counter-Strike 2

<div align="center">

![CS2 Lineups](https://img.shields.io/badge/CS2-Lineups-orange?style=for-the-badge)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-blue?style=for-the-badge&logo=postgresql)
![Go](https://img.shields.io/badge/Go-1.21-00ADD8?style=for-the-badge&logo=go)
![Status](https://img.shields.io/badge/Status-Production%20Ready-success?style=for-the-badge)

**Site web moderne pour apprendre et maîtriser les lineups de grenades sur Counter-Strike 2**

[Documentation complète](GUIDE_COMPLET.md) • [Installation](#installation) • [Déploiement](#déploiement)

</div>

---

## 📋 Table des matières

- [À propos](#-à-propos)
- [Fonctionnalités](#-fonctionnalités)
- [Installation](#-installation)
- [Utilisation](#-utilisation)
- [Déploiement](#-déploiement)
- [Documentation](#-documentation)
- [Stack technique](#-stack-technique)
- [Roadmap](#-roadmap)

---

## 🔥 À propos

**CS2 Lineups** est un site web professionnel qui centralise les meilleurs lineups de grenades pour Counter-Strike 2. 

### Problème résolu

Les joueurs de CS2 ont besoin de lineups fiables pour :
- 💨 Smokes stratégiques
- ⚡ Flashs parfaitement timés
- 🔥 Molotovs pour bloquer les rush
- 💣 HE grenades pour infliger des dégâts

Ce site offre une interface moderne avec filtres avancés pour trouver le lineup parfait en quelques clics.

### Caractéristiques principales

✅ **7 maps du pool compétitif** - Mirage, Dust II, Inferno, Nuke, Ancient, Anubis, Vertigo  
✅ **54 lineups détaillés** - Avec position, visée, et démo GIF  
✅ **Filtres temps réel** - Par grenade, camp (T/CT), et difficulté  
✅ **Design moderne** - Interface dark mode responsive  
✅ **Base PostgreSQL** - Performance et scalabilité  

---

## ⚡ Fonctionnalités

### Interface utilisateur

- 🗺️ **Grid de maps interactive** - Navigation visuelle par map
- 🔍 **Filtres avancés** - Grenade, Side, Difficulté en temps réel
- 📱 **Responsive design** - Mobile, tablet, desktop
- 🌙 **Dark mode** - Thème sombre optimisé pour les gamers
- ⚡ **Performance** - Chargement ultra-rapide

### Détails des lineups

Chaque lineup inclut:
- 📍 Position exact du joueur (screenshot)
- 🎯 Placement du viseur (crosshair zoomé)
- 🎬 Démonstration animée (GIF)
- ⚙️ Bind console (copiable en 1 clic)
- 📊 Métadonnées (difficulté, popularité, vues)
- 🏷️ Tags de catégorisation

### Backend

- 🗄️ **PostgreSQL** - Base de données optimisée
- 🔄 **Fallback JSON** - Développement sans DB
- 📈 **Queries optimisées** - Index et JOIN performants
- 🛡️ **Gestion d'erreurs** - Robustesse production

---

## 🚀 Installation

### Prérequis

- Go 1.21+
- PostgreSQL 15+
- Git

### Installation rapide

```bash
# 1. Cloner le repository
git clone https://github.com/votre-username/cs2-lineups.git
cd yboost_scalingo

# 2. Installer les dépendances Go
go mod download

# 3. Configurer PostgreSQL
sudo -u postgres psql -c "CREATE DATABASE cs2_lineups;"
sudo -u postgres psql -d cs2_lineups -f schema.sql

# 4. Créer un utilisateur
sudo -u postgres psql -d cs2_lineups -c "
  CREATE USER cs2user WITH PASSWORD 'cs2password';
  GRANT ALL PRIVILEGES ON DATABASE cs2_lineups TO cs2user;
  GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO cs2user;
  GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO cs2user;
"

# 5. Importer les données de départ
DATABASE_URL="postgresql://cs2user:cs2password@localhost:5432/cs2_lineups?sslmode=disable" \
  go run cmd/seed/seed.go

# 6. Lancer le serveur
DATABASE_URL="postgresql://cs2user:cs2password@localhost:5432/cs2_lineups?sslmode=disable" \
  go run main.go models.go
```

Le site est maintenant accessible sur **http://localhost:8080** 🎉

---

## 💻 Utilisation

### Démarrage rapide

```bash
# Avec PostgreSQL (recommandé)
DATABASE_URL="postgresql://cs2user:cs2password@localhost:5432/cs2_lineups?sslmode=disable" \
  go run main.go models.go

# En mode développement (JSON uniquement)
go run main.go models.go
```

### Variables d'environnement

``bash
# Database URL (obligatoire pour PostgreSQL)
DATABASE_URL="postgresql://user:password@host:port/database?sslmode=disable"

# Port du serveur (optionnel, défaut: 8080)
PORT=8080
```

### Ajouter des lineups

**Méthode 1: Éditer seed_data.json**

```bash
# 1. Éditer seed_data.json
# 2. Réimporter
echo "Y" | DATABASE_URL="..." go run cmd/seed/seed.go
```

**Méthode 2: SQL direct**

```sql
INSERT INTO lineups (id, map_id, title, grenade_type, side, ...) VALUES (...);
```

Voir [GUIDE_COMPLET.md](GUIDE_COMPLET.md) pour plus de détails.

---

## 🚢 Déploiement

### Déploiement sur Scalingo

```bash
# 1. Créer l'application
scalingo create votre-app-cs2

# 2. Ajouter PostgreSQL
scalingo -a votre-app-cs2 addons-add postgresql postgresql-starter-512

# 3. Déployer
git push scalingo main

# 4. Importer les données
scalingo -a votre-app-cs2 run bash
DATABASE_URL=$DATABASE_URL go run cmd/seed/seed.go
exit
```

Votre site est maintenant en ligne sur **https://votre-app-cs2.osc-fr1.scalingo.io** 🚀

---

## 📚 Documentation

### Guides disponibles

| Document | Description |
|----------|-------------|
| [RESUME.md](RESUME.md) | Vue d'ensemble rapide du projet (5 min) |
| [GUIDE_COMPLET.md](GUIDE_COMPLET.md) | Guide exhaustif (installation, utilisation, déploiement) |
| [SOURCES_DONNEES.md](SOURCES_DONNEES.md) | Comment obtenir du contenu légalement |
| [CHECKLIST.md](CHECKLIST.md) | TODO et planning 4 semaines |
| [DATABASE_SETUP.md](DATABASE_SETUP.md) | Configuration PostgreSQL avancée |

### Routes disponibles

| Route | Description |
|-------|-------------|
| `GET /` | Page d'accueil (liste des maps) |
| `GET /map/{id}` | Détail d'une map avec filtres |
| `GET /lineup/{id}` | Détail d'un lineup |
| `GET /static/*` | Fichiers statiques (CSS, JS) |
| `GET /assets/*` | Médias (images, GIFs) |

---

## 🛠️ Stack technique

### Backend

- **Langage:** Go 1.21+
- **Router:** Gorilla Mux
- **Database:** PostgreSQL 15
- **Driver:** lib/pq
- **Templates:** Go html/template

### Frontend

- **HTML5** - Structure sémantique
- **CSS3** - Variables CSS, Grid, Flexbox
- **JavaScript Vanilla** - Filtres temps réel
- **Design:** Dark mode, gradients, animations

### Infrastructure

- **Hosting:** Scalingo (production)
- **Database:** PostgreSQL (Scalingo addon)

---

## 🎯 Roadmap

### ✅ Version 1.0 (Actuelle)

- [x] Base de données complète (7 maps, 54 lineups)
- [x] Interface moderne avec filtres
- [x] 3 pages principales (home, map, lineup)
- [x] Documentation complète
- [x] Scripts d'import/seed

### 🚧 Version 1.1 (Prochaines 2-3 semaines)

- [ ] Vraies images/GIFs (remplacer placeholders)
- [ ] Déploiement production Scalingo
- [ ] SEO optimization (meta tags, sitemap)
- [ ] Google Analytics

### 🔮 Version 2.0 (4-6 semaines)

- [ ] Interface d'administration
- [ ] API REST publique
- [ ] Système de votes et commentaires
- [ ] Multi-langue (FR/EN/ES)

---

## 📊 Statistiques actuelles

```
✅ 7 maps du pool compétitif
✅ 54 lineups détaillés
✅ 27 tags différents
✅ 4 types de grenades (smoke, flash, molotov, HE)
✅ Support T et CT side
✅ 3 niveaux de difficulté
```

---

## 🔧 Structure du projet

```
yboost_scalingo/
├── main.go                 # Serveur web principal
├── models.go               # Structures de données
├── schema.sql              # Schéma PostgreSQL
├── seed_data.json          # Données complètes (54 lineups)
├── go.mod                  # Dépendances Go
├── Procfile                # Configuration Scalingo
├── cmd/
│   └── seed/
│       └── seed.go         # Script d'import
├── templates/
│   ├── home.html           # Page d'accueil
│   ├── map-detail.html     # Page map avec filtres
│   └── lineup-detail.html  # Page lineup détaillé
├── static/
│   ├── css/
│   │   └── style.css       # Styles personnalisés
│   └── js/
│       ├── app.js          # JavaScript général
│       └── reader.js       # (ancien, à supprimer)
└── public/
    └── assets/
        ├── lineups/        # Images de lineups
        ├── maps/           # Images de maps
        └── ui/             # Assets UI
```

---

## 📄 Licence

Ce projet est sous licence **MIT**. 

**Disclaimer:** 
CS2 Lineups n'est pas affilié à Valve Corporation. Counter-Strike et le logo CS sont des marques déposées de Valve Corporation. Tout le contenu de gameplay appartient à Valve Corporation.

---

## 🙏 Remerciements

- **Valve Corporation** - Pour Counter-Strike 2
- **Communauté CS2** - Pour les lineups et stratégies
- **Open Source** - Go, PostgreSQL, et toutes les librairies utilisées

---

<div align="center">

**Fait avec ❤️ pour la communauté CS2**

⭐ Star ce projet si vous le trouvez utile!

</div>
