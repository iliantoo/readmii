# 🔨 Guide de Compilation

## ✅ Compilation Correcte

Pour compiler le projet, utilisez **UNE** de ces commandes :

### Option 1 : Compiler tout le package (RECOMMANDÉ)
```bash
go build -o bin/readmii .
```

### Option 2 : Utiliser go run (pour debug)
```bash
go run .
```

### Option 3 : Compiler en spécifiant tous les fichiers
```bash
go build -o bin/readmii main.go models.go
```

---

## ❌ NE PAS FAIRE

**Ne compilez jamais uniquement main.go :**
```bash
go build main.go  # ❌ ERREUR : undefined Library, Map, Lineup, etc.
```

**Raison :** Les structures (Map, Lineup, Library, etc.) sont définies dans [models.go](models.go), pas dans [main.go](main.go).

---

## 🚀 Lancer le Serveur

### Local (JSON)
```bash
./bin/readmii
# ou
go run .
```

### Local (PostgreSQL)
```bash
export DATABASE_URL="postgres://localhost/cs2_lineups?sslmode=disable"
./bin/readmii
```

---

## 📦 Déploiement Scalingo

Scalingo compile automatiquement avec :
```bash
go build -o bin/readmii .
```

Vous n'avez rien à faire, juste :
```bash
git push scalingo main
```

---

## 🔧 Migration PostgreSQL

Pour migrer les données JSON vers PostgreSQL :

```bash
cd cmd/migrate
export DATABASE_URL="postgres://localhost/cs2_lineups?sslmode=disable"
go run migrate.go models.go

# Ou compiler d'abord :
go build -o migrate migrate.go models.go
./migrate
```

---

## 📁 Structure des Fichiers

```
main.go       → Code du serveur web (handlers, routes, DB)
models.go     → Structures de données (Map, Lineup, Library)
lineups.json  → Données JSON (fallback si pas de DB)
schema.sql    → Schéma PostgreSQL
cmd/migrate/  → Script de migration
```

**Important :** `models.go` contient les structures utilisées par **main.go** ET **cmd/migrate/migrate.go**.
