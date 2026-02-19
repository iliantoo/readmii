# 🗄️ Configuration de la Base de Données PostgreSQL

## 📋 Vue d'ensemble

Votre projet CS2 Lineups supporte maintenant PostgreSQL avec fallback automatique sur JSON.

**Mode de fonctionnement :**
- ✅ Si `DATABASE_URL` est définie → PostgreSQL 
- ✅ Si `DATABASE_URL` absente → Fichier JSON

---

## 🚀 Configuration sur Scalingo

### Étape 1 : Ajouter PostgreSQL

```bash
# Ajouter l'addon PostgreSQL à votre app
scalingo -a <votre-app> addons-add postgresql postgresql-starter-512

# Vérifier que la DB est créée
scalingo -a <votre-app> addons
```

Scalingo configure automatiquement la variable `DATABASE_URL`.

### Étape 2 : Créer le schéma

```bash
# Se connecter à la console Scalingo
scalingo -a <votre-app> pgsql-console

# Copier-coller le contenu de schema.sql
# Ou l'exécuter directement :
scalingo -a <votre-app> run bash
cat schema.sql | psql $DATABASE_URL
```

### Étape 3 : Migrer les données

```bash
# Depuis votre machine locale
scalingo -a <votre-app> run go run cmd/migrate/migrate.go

# Ou déployer et exécuter :
git add cmd/migrate/
git commit -m "Add migration script"
git push scalingo main
scalingo -a <votre-app> run go run cmd/migrate/migrate.go
```

---

## 💻 Configuration en Local

### Étape 1 : Installer PostgreSQL

**Ubuntu/Debian :**
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```

**macOS :**
```bash
brew install postgresql@15
brew services start postgresql@15
```

### Étape 2 : Créer la base de données

```bash
# Se connecter à PostgreSQL
sudo -u postgres psql

# Créer la base de données
CREATE DATABASE cs2_lineups;

# Créer un utilisateur (optionnel)
CREATE USER cs2user WITH PASSWORD 'votre_password';
GRANT ALL PRIVILEGES ON DATABASE cs2_lineups TO cs2user;

# Quitter
\q
```

### Étape 3 : Appliquer le schéma

```bash
# Depuis le dossier du projet
psql -U postgres -d cs2_lineups -f schema.sql

# Ou avec votre utilisateur
psql -U cs2user -d cs2_lineups -f schema.sql
```

### Étape 4 : Migrer les données JSON

```bash
# Définir DATABASE_URL
export DATABASE_URL="postgres://localhost/cs2_lineups?sslmode=disable"

# Ou avec utilisateur/password
export DATABASE_URL="postgres://cs2user:votre_password@localhost/cs2_lineups?sslmode=disable"

# Exécuter la migration
cd cmd/migrate
go run migrate.go models.go
```

### Étape 5 : Lancer le serveur avec PostgreSQL

```bash
# Avec DATABASE_URL définie, le serveur utilisera PostgreSQL
go run main.go

# Vous verrez :
# "🚀 Serveur CS2 Lineups démarré sur http://localhost:8080 (PostgreSQL)"
```

---

## 🔄 Passer de JSON à PostgreSQL

Le code détecte automatiquement la présence de `DATABASE_URL` :

**Sans DATABASE_URL (JSON) :**
```bash
go run main.go
# → "🚀 Serveur CS2 Lineups démarré (JSON)"
```

**Avec DATABASE_URL (PostgreSQL) :**
```bash
export DATABASE_URL="postgres://localhost/cs2_lineups?sslmode=disable"
go run main.go
# → "🚀 Serveur CS2 Lineups démarré (PostgreSQL)"
```

---

## 📊 Commandes PostgreSQL Utiles

### Voir les données

```sql
-- Se connecter
psql -U postgres -d cs2_lineups

-- Lister les tables
\dt

-- Voir toutes les cartes
SELECT id, name, total_lineups FROM maps;

-- Voir les lineups d'une carte
SELECT id, title, grenade_type, side FROM lineups WHERE map_id = 'mirage';

-- Compter les lineups par type de grenade
SELECT grenade_type, COUNT(*) 
FROM lineups 
GROUP BY grenade_type;

-- Voir les lineups les plus populaires
SELECT title, popularity, views 
FROM lineups 
ORDER BY popularity DESC, views DESC 
LIMIT 10;
```

### Gérer les données

```sql
-- Ajouter un lineup
INSERT INTO lineups (
    id, map_id, title, description, grenade_type, side, difficulty,
    throw_zone, landing_zone, action_type,
    bind_required, movement, click_type,
    popularity, views
) VALUES (
    'mir_smoke_005', 'mirage', 'Nouveau Smoke', 'Description...',
    'smoke', 'T', 'easy',
    'T Spawn', 'B Site', 'jumpthrow',
    true, 'none', 'left_click',
    4.5, 0
);

-- Mettre à jour les vues
UPDATE lineups SET views = views + 1 WHERE id = 'mir_smoke_001';

-- Supprimer un lineup
DELETE FROM lineups WHERE id = 'mir_smoke_005';

-- Vider toutes les tables (attention !)
TRUNCATE maps, lineups, tags, lineup_tags CASCADE;
```

---

## 🔍 Vérification de la Migration

Après la migration, vérifiez que tout est OK :

```bash
# Compter les cartes
psql -U postgres -d cs2_lineups -c "SELECT COUNT(*) FROM maps;"

# Compter les lineups
psql -U postgres -d cs2_lineups -c "SELECT COUNT(*) FROM lineups;"

# Voir un lineup complet
psql -U postgres -d cs2_lineups -c "SELECT * FROM lineups WHERE id = 'mir_smoke_001';"
```

---

## 🛠️ Troubleshooting

### Erreur : "database does not exist"
```bash
createdb cs2_lineups
```

### Erreur : "role does not exist"
```bash
createuser cs2user
```

### Erreur : "connection refused"
```bash
# Vérifier que PostgreSQL tourne
sudo systemctl status postgresql
# Ou sur mac :
brew services list
```

### Réinitialiser complètement la base

```bash
# Se connecter
psql -U postgres

# Supprimer et recréer
DROP DATABASE cs2_lineups;
CREATE DATABASE cs2_lineups;
\q

# Réappliquer le schéma
psql -U postgres -d cs2_lineups -f schema.sql

# Remigrer les données
go run migrate.go
```

---

## 📈 Performance

### Index créés automatiquement

Le fichier `schema.sql` crée les index suivants :
- `idx_lineups_map_id` - Recherche par carte
- `idx_lineups_grenade_type` - Filtre par type de grenade
- `idx_lineups_side` - Filtre par side (T/CT)
- `idx_lineups_difficulty` - Filtre par difficulté
- `idx_lineups_popularity` - Tri par popularité
- `idx_lineups_views` - Tri par vues

### Optimisations futures

Si vous avez beaucoup de lineups (1000+), considérez :
- Pagination des résultats
- Cache Redis pour les requêtes fréquentes
- Connection pooling

---

## 🎯 Prochaines Étapes

1. ✅ **Déployer sur Scalingo** avec PostgreSQL
2. ✅ **Migrer vos données** JSON vers PostgreSQL
3. ✅ **Tester** que tout fonctionne
4. 🔜 **Ajouter des lineups** directement en base
5. 🔜 **Créer une interface admin** pour gérer les lineups

---

**Votre projet CS2 Lineups est maintenant prêt pour PostgreSQL ! 🚀**
