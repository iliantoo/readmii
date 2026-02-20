# Changelog - CS2 Lineups

Toutes les modifications notables de ce projet seront documentées dans ce fichier.

Le format est basé sur [Keep a Changelog](https://keepachangelog.com/fr/1.0.0/),
et ce projet adhère au [Semantic Versioning](https://semver.org/lang/fr/).

## [Non publié]

### À venir
- Interface d'administration
- Recherche globale avec auto-complétion
- Partage social (Twitter, Reddit, Discord)
- Système de favoris
- Dark/Light mode toggle

---

## [1.0.0] - 2025-01-15

### 🎉 Version initiale

#### Ajouté
- **Base de données PostgreSQL complète** avec 4 tables (maps, lineups, tags, lineup_tags)
- **7 maps du pool compétitif**: Mirage, Dust II, Inferno, Nuke, Ancient, Anubis, Vertigo
- **54 lineups détaillés** avec métadonnées complètes
- **27 tags** pour catégorisation (execute, retake, standard, etc.)
- **3 pages principales**:
  - Homepage: Grid des maps avec statistiques
  - Map detail: Liste des lineups avec filtres temps réel
  - Lineup detail: Informations complètes avec médias
- **Système de filtres avancés**:
  - Par type de grenade (smoke, flash, molotov, HE)
  - Par side (T, CT)
  - Par difficulté (easy, medium, hard)
- **Design moderne**:
  - Dark mode optimisé pour gamers
  - Gradients et animations CSS
  - Responsive (mobile, tablet, desktop)
  - Variables CSS pour personnalisation facile
- **Backend Go performant**:
  - Serveur HTTP avec gorilla/mux
  - Connexion PostgreSQL avec lib/pq
  - Fallback JSON pour développement
  - Templates Go html/template
- **Scripts de gestion**:
  - `cmd/seed/seed.go`: Import de données depuis JSON
  - Option TRUNCATE pour réinitialisation complète
  - Gestion automatique des tags et relations
- **Documentation complète**:
  - README.md: Vue d'ensemble et installation
  - RESUME.md: Synthèse rapide (5 min)
  - GUIDE_COMPLET.md: Guide exhaustif (installation, utilisation, déploiement)
  - SOURCES_DONNEES.md: Stratégies d'acquisition de contenu
  - CHECKLIST.md: TODO et planning 4 semaines
  - DATABASE_SETUP.md: Configuration PostgreSQL avancée
  - CONTRIBUTING.md: Guide pour contributeurs
  - LICENSE: Licence MIT avec disclaimer Valve
- **Configuration Scalingo**:
  - Procfile pour déploiement
  - Support PostgreSQL addon
  - Variables d'environnement

#### Optimisations
- **Index PostgreSQL** sur colonnes critiques (grenade_type, side, difficulty, popularity, views)
- **Queries optimisées** avec JOIN au lieu de requêtes multiples
- **Assets optimisés** (recommandations de taille et format)
- **Chargement rapide** sans frameworks lourds

#### Sécurité
- Protection contre les injections SQL (statements préparés)
- Validation côté serveur
- Gestion d'erreurs robuste
- Logs pour debugging

---

## Statistiques version 1.0.0

```
✅ 7 maps
✅ 54 lineups
✅ 27 tags
✅ 4 types de grenades
✅ 3 niveaux de difficulté
✅ 2 sides (T/CT)
✅ ~2000 lignes de code
✅ 6 documents de documentation
```

---

## Formats de changelog

### Types de changements
- **Ajouté** - Nouvelles fonctionnalités
- **Modifié** - Changements dans les fonctionnalités existantes
- **Déprécié** - Fonctionnalités bientôt supprimées
- **Supprimé** - Fonctionnalités retirées
- **Corrigé** - Corrections de bugs
- **Sécurité** - Corrections de vulnérabilités

### Liens de version
[Non publié]: https://github.com/votre-username/cs2-lineups/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/votre-username/cs2-lineups/releases/tag/v1.0.0
