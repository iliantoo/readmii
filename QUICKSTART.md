╔════════════════════════════════════════════════════════════════════════════╗
║                                                                            ║
║                   🎮 CS2 LINEUPS - QUICK START GUIDE 🎮                   ║
║                                                                            ║
║                    Démarrez votre site en 5 minutes !                     ║
║                                                                            ║
╚════════════════════════════════════════════════════════════════════════════╝


┌─────────────────────────────────────────────────────────────────────────┐
│ ✅ ÉTAPE 1 : VÉRIFIER LES PRÉREQUIS                                      │
└─────────────────────────────────────────────────────────────────────────┘

Vérifiez que vous avez :
  □ Go 1.21+ installé       → go version
  □ PostgreSQL 15+ installé → psql --version
  □ Git (optionnel)         → git --version


┌─────────────────────────────────────────────────────────────────────────┐
│ 🗄️ ÉTAPE 2 : CONFIGURER POSTGRESQL (si pas déjà fait)                  │
└─────────────────────────────────────────────────────────────────────────┘

Copiez-collez ces commandes dans votre terminal :

  # Créer la base de données
  sudo -u postgres psql -c "CREATE DATABASE cs2_lineups;"
  
  # Appliquer le schéma
  sudo -u postgres psql -d cs2_lineups -f schema.sql
  
  # Créer l'utilisateur
  sudo -u postgres psql -d cs2_lineups -c "
    CREATE USER cs2user WITH PASSWORD 'cs2password';
    GRANT ALL PRIVILEGES ON DATABASE cs2_lineups TO cs2user;
    GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO cs2user;
    GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO cs2user;
  "


┌─────────────────────────────────────────────────────────────────────────┐
│ 📦 ÉTAPE 3 : IMPORTER LES DONNÉES                                       │
└─────────────────────────────────────────────────────────────────────────┘

Une seule commande pour importer les 54 lineups :

  DATABASE_URL="postgresql://cs2user:cs2password@localhost:5432/cs2_lineups?sslmode=disable" \
    go run cmd/seed/seed.go

Répondez "Y" quand demandé pour confirmer l'import.


┌─────────────────────────────────────────────────────────────────────────┐
│ 🚀 ÉTAPE 4 : LANCER LE SERVEUR                                          │
└─────────────────────────────────────────────────────────────────────────┘

Lancez le serveur web :

  DATABASE_URL="postgresql://cs2user:cs2password@localhost:5432/cs2_lineups?sslmode=disable" \
    go run main.go models.go

Vous devriez voir :
  ✅ Connexion PostgreSQL réussie
  🚀 Serveur CS2 Lineups démarré sur http://localhost:8080


┌─────────────────────────────────────────────────────────────────────────┐
│ 🌐 ÉTAPE 5 : TESTER LE SITE                                             │
└─────────────────────────────────────────────────────────────────────────┘

Ouvrez votre navigateur sur : http://localhost:8080

Testez les fonctionnalités :
  □ Homepage affiche la grid des 7 maps
  □ Cliquez sur "Mirage" → affiche les lineups
  □ Utilisez les filtres (Smoke, T/CT, Easy)
  □ Cliquez sur un lineup → détails complets
  □ Copiez un bind console

✅ Félicitations ! Votre site fonctionne !


╔════════════════════════════════════════════════════════════════════════════╗
║                           📋 COMMANDES UTILES                              ║
╚════════════════════════════════════════════════════════════════════════════╝

# Démarrer le serveur
DATABASE_URL="postgresql://cs2user:cs2password@localhost:5432/cs2_lineups?sslmode=disable" \
  go run main.go models.go

# Réimporter les données (ATTENTION: écrase tout)
echo "Y" | DATABASE_URL="postgresql://cs2user:cs2password@localhost:5432/cs2_lineups?sslmode=disable" \
  go run cmd/seed/seed.go

# Vérifier la base de données
psql -U cs2user -d cs2_lineups -c "SELECT COUNT(*) FROM lineups;"

# Tester le serveur avec curl
curl http://localhost:8080/ | grep CS2


╔════════════════════════════════════════════════════════════════════════════╗
║                        🔧 DÉPANNAGE RAPIDE                                 ║
╚════════════════════════════════════════════════════════════════════════════╝

PROBLÈME : "psql: error: connection to server on socket"
SOLUTION : PostgreSQL n'est pas démarré
  → sudo systemctl start postgresql
  → brew services start postgresql (macOS)

PROBLÈME : "database cs2_lineups does not exist"
SOLUTION : Créer la base
  → sudo -u postgres psql -c "CREATE DATABASE cs2_lineups;"

PROBLÈME : "panic: pq: password authentication failed"
SOLUTION : Mauvais mot de passe utilisateur
  → Vérifier le mot de passe dans DATABASE_URL
  → Réinitialiser: ALTER USER cs2user WITH PASSWORD 'cs2password';

PROBLÈME : "bind: address already in use"
SOLUTION : Port 8080 déjà utilisé
  → lsof -ti:8080 | xargs kill -9
  → Ou changer le port: PORT=3000 go run main.go models.go

PROBLÈME : Page vide ou erreur 500
SOLUTION : Vérifier que les données sont importées
  → psql -U cs2user -d cs2_lineups -c "SELECT COUNT(*) FROM lineups;"
  → Devrait afficher "54"


╔════════════════════════════════════════════════════════════════════════════╗
║                          📚 PROCHAINES ÉTAPES                              ║
╚════════════════════════════════════════════════════════════════════════════╝

1. ⭐ AJOUTER DE VRAIES IMAGES
   → Capturer des screenshots en jeu
   → Mettre dans public/assets/lineups/
   → Éditer seed_data.json avec les vrais chemins
   → Réimporter les données

2. 🚢 DÉPLOYER EN PRODUCTION
   → Lire GUIDE_COMPLET.md section "Déploiement Scalingo"
   → scalingo create votre-app
   → git push scalingo main

3. 📊 SEO ET MARKETING
   → Ajouter Google Analytics
   → Partager sur Reddit et Discord
   → Créer sitemap.xml

4. 🛠️ DÉVELOPPER PLUS
   → Interface d'administration
   → Recherche globale
   → Système de votes


╔════════════════════════════════════════════════════════════════════════════╗
║                        📖 DOCUMENTATION COMPLÈTE                           ║
╚════════════════════════════════════════════════════════════════════════════╝

README.md              → Vue d'ensemble du projet
RESUME.md              → Résumé rapide (5 min)
GUIDE_COMPLET.md       → Guide exhaustif (tout ce qu'il faut savoir)
SOURCES_DONNEES.md     → Comment obtenir du contenu
CHECKLIST.md           → Planning 4 semaines
DATABASE_SETUP.md      → Configuration PostgreSQL avancée
CONTRIBUTING.md        → Comment contribuer
TODO.txt               → TODO quotidien simple


╔════════════════════════════════════════════════════════════════════════════╗
║                              🎯 OBJECTIFS                                  ║
╚════════════════════════════════════════════════════════════════════════════╝

Semaine 1  → 5-10 lineups avec vraies images
Semaine 2  → Déploiement production Scalingo
Semaine 3  → SEO + Partage sur réseaux sociaux
Semaine 4  → 30-40 lineups de qualité


╔════════════════════════════════════════════════════════════════════════════╗
║                            ℹ️ INFORMATIONS                                 ║
╚════════════════════════════════════════════════════════════════════════════╝

🔹 Database:  cs2_lineups
🔹 User:      cs2user
🔹 Password:  cs2password
🔹 Port:      5432 (PostgreSQL), 8080 (Web)
🔹 URL:       http://localhost:8080
🔹 Maps:      7 (Mirage, Dust2, Inferno, Nuke, Ancient, Anubis, Vertigo)
🔹 Lineups:   54
🔹 Tags:      27


╔════════════════════════════════════════════════════════════════════════════╗
║                Good luck et have fun ! 🎮                                  ║
╚════════════════════════════════════════════════════════════════════════════╝
