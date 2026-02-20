# 🎮 CS2 Lineups - Guide Complet

## 📊 État actuel du projet

### ✅ Ce qui est fait

**Base de données PostgreSQL:**
- ✅ 7 maps du pool compétitif (Mirage, Dust II, Inferno, Nuke, Ancient, Anubis, Vertigo)
- ✅ 54 lineups professionnels avec métadonnées complètes
- ✅ Système de tags et catégorisation
- ✅ Schema optimisé avec index pour les performances

**Interface utilisateur moderne:**
- ✅ Page d'accueil avec grid de maps
- ✅ Page de détail par map avec filtres avancés
- ✅ Page de détail par lineup avec toutes les informations
- ✅ Design dark mode responsive
- ✅ Filtres en temps réel (grenade, side, difficulté)

**Backend Go:**
- ✅ API REST avec routes dynamiques
- ✅ Support PostgreSQL avec fallback JSON
- ✅ Templates HTML optimisés
- ✅ Gestion des médias

---

## 🚀 Démarrage rapide

### 1. Lancer le serveur localement

```bash
# Avec PostgreSQL
DATABASE_URL="postgresql://cs2user:cs2password@localhost:5432/cs2_lineups?sslmode=disable" go run main.go models.go

# Le serveur démarre sur http://localhost:8080
```

### 2. Accéder au site

- **Page d'accueil:** http://localhost:8080/
- **Exemple map:** http://localhost:8080/map/mirage
- **Exemple lineup:** http://localhost:8080/lineup/mirage-a-ct-smoke-tspawn

---

## 📝 Ajouter des lineups

### Méthode 1: Modifier seed_data.json et réimporter

1. Éditez `seed_data.json`
2. Ajoutez vos lineups dans le format suivant:

```json
{
  "id": "unique-lineup-id",
  "title": "Titre du lineup",
  "description": "Description détaillée",
  "grenade_type": "smoke|flash|molotov|he",
  "side": "T|CT",
  "difficulty": "easy|medium|hard",
  "throw_zone": "Zone de départ",
  "landing_zone": "Zone d'impact",
  "action_type": "jumpthrow|throw",
  "action_details": {
    "bind_required": true|false,
    "bind_command": "alias \"+jumpthrow\" \"+jump;-attack\"...",
    "movement": "standing|walking|running|crouching",
    "click_type": "left-click|right-click"
  },
  "media": {
    "position_image": "/assets/lineups/map/image.jpg",
    "position_thumbnail": "/assets/lineups/map/thumb.jpg",
    "crosshair_image": "/assets/lineups/map/crosshair.jpg",
    "crosshair_zoom_level": "default|zoomed",
    "demo_gif": "/assets/lineups/map/demo.gif",
    "demo_thumbnail": "/assets/lineups/map/demo-thumb.jpg"
  },
  "tags": ["Tag1", "Tag2"],
  "popularity": 8.5,
  "views": 1200
}
```

3. Réimportez les données:

```bash
echo "Y" | DATABASE_URL="postgresql://cs2user:cs2password@localhost:5432/cs2_lineups?sslmode=disable" go run cmd/seed/seed.go
```

### Méthode 2: Insertion SQL directe

```sql
-- Insérer un lineup
INSERT INTO lineups (
    id, map_id, title, description, grenade_type, side, difficulty,
    throw_zone, landing_zone, action_type,
    bind_required, bind_command, movement, click_type,
    position_image, position_thumbnail, crosshair_image, crosshair_zoom_level,
    demo_gif, demo_thumbnail,
    popularity, views
) VALUES (
    'votre-lineup-id',
    'mirage',
    'Titre',
    'Description',
    'smoke',
    'T',
    'easy',
    'Zone départ',
    'Zone impact',
    'jumpthrow',
    true,
    'alias "+jumpthrow" "+jump;-attack"',
    'standing',
    'left-click',
    '/assets/lineups/mirage/pos.jpg',
    '/assets/lineups/mirage/pos-thumb.jpg',
    '/assets/lineups/mirage/crosshair.jpg',
    'default',
    '/assets/lineups/mirage/demo.gif',
    '/assets/lineups/mirage/demo-thumb.jpg',
    8.5,
    1000
);
```

---

## 🖼️ Gestion des médias

### Structure des dossiers

```
public/
└── assets/
    ├── maps/
    │   ├── mirage.jpg
    │   ├── dust2.jpg
    │   └── ...
    └── lineups/
        ├── mirage/
        │   ├── a-ct-smoke-position.jpg
        │   ├── a-ct-smoke-thumb.jpg
        │   ├── a-ct-smoke-crosshair.jpg
        │   └── a-ct-smoke-demo.gif
        └── dust2/
            └── ...
```

### Options pour obtenir des images

**Option 1: Créer vos propres captures d'écran**
- Lancez CS2 en mode practice
- Utilisez `sv_cheats 1` et `noclip` pour positionner la caméra
- `F12` pour capturer (Steam) ou OBS pour enregistrer

**Option 2: Placeholder temporaires**
- Utilisez des emojis/icônes comme dans les templates actuels
- Remplacez progressivement par de vraies images

**Option 3: Utiliser des ressources existantes**
- ⚠️ Vérifiez les droits d'auteur avant utilisation
- Sites communautaires CS2 (avec permission)
- Créez votre propre contenu pour éviter les problèmes légaux

---

## 🔧 APIs et sources de données

### ❌ Ce qui n'existe PAS

- ❌ API publique officielle de Valve pour les lineups
- ❌ Base de données open-source exhaustive
- ❌ Web scraping légal des sites concurrents

### ✅ Solutions recommandées

**1. Crowdsourcing communautaire**
```
- Créez un formulaire Google Forms
- Laissez la communauté soumettre des lineups
- Validez et importez manuellement
```

**2. Partenariats**
```
- Contactez des créateurs de contenu CS2
- Proposez de référencer leurs vidéos YouTube
- Échangez visibilité contre contenu
```

**3. Création progressive**
```
- Ajoutez 2-3 lineups essentiels par map
- Concentrez-vous sur la qualité
- Complétez au fil du temps
```

**4. Import depuis des sources légales**
```
- Utilisez des vidéos YouTube (avec attribution)
- Serveurs Discord CS2 avec permissions
- Workshops Steam (avec crédits)
```

---

## 🎨 Personnalisation de l'UI

### Modifier les couleurs

Éditez les templates HTML (variables CSS):

```css
:root {
    --primary: #ff6b35;      /* Couleur principale */
    --secondary: #4a9eff;    /* Couleur secondaire */
    --success: #44dd88;      /* Succès/Validé */
}
```

### Ajouter des filtres

Dans `map-detail.html`, ajoutez un nouveau groupe de filtres:

```html
<div class="filter-group">
    <span class="filter-label">Votre filtre</span>
    <div class="filter-buttons">
        <button class="filter-btn" data-filter="votrefiltre" data-value="valeur">
            Label
        </button>
    </div>
</div>
```

---

## 📦 Déploiement sur Scalingo

### 1. Préparer le déploiement

```bash
# Vérifier que tout est commité
git add .
git commit -m "Site CS2 Lineups complet"
```

### 2. Configurer PostgreSQL sur Scalingo

```bash
# Ajouter l'addon
scalingo -a votre-app addons-add postgresql postgresql-starter-512

# Vérifier que DATABASE_URL est configurée
scalingo -a votre-app env
```

### 3. Importer les données en production

```bash
# Se connecter à Scalingo
scalingo -a votre-app run bash

# Lancer l'import
go run cmd/seed/seed.go
```

### 4. Déployer

```bash
git push scalingo main
```

---

## 🛠️ Interface d'administration (À développer)

**Fonctionnalités suggérées:**

```
/admin
  ├── /lineups          → Liste et CRUD des lineups
  ├── /maps             → Gestion des maps
  ├── /users            → Gestion des contributeurs
  └── /media            → Upload d'images/GIFs
```

**Stack recommandée:**
- Backend: Routes Go additionnelles
- Frontend: Admin panel simple en HTML/JS
- Auth: Basic Auth ou JWT
- Upload: Système de gestion de fichiers

---

## 📊 Statistiques actuelles

```
✅ 7 maps du pool compétitif
✅ 54 lineups détaillés
✅ 27 tags différents
✅ 4 types de grenades
✅ 2 sides (T/CT)
✅ 3 niveaux de difficulté
```

---

## 🎯 Prochaines étapes recommandées

### Court terme (1-2 semaines)

1. **Ajouter de vraies images**
   - Capturez 5-10 lineups essentiels par map
   - Position + Crosshair + GIF démo

2. **Créer un formulaire de contribution**
   - Google Forms ou Typeform
   - Permettre à la communauté de soumettre

3. **SEO et partage**
   - Meta tags OpenGraph
   - Descriptions optimisées
   - Sitemap XML

### Moyen terme (1 mois)

4. **Interface d'administration**
   - CRUD des lineups
   - Upload d'images
   - Modération

5. **Fonctionnalités sociales**
   - Système de votes
   - Commentaires
   - Notes de qualité

6. **Analytics**
   - Google Analytics
   - Tracking des lineups populaires
   - A/B testing

### Long terme (3+ mois)

7. **API publique**
   - Exposer vos données
   - Documentation Swagger
   - Rate limiting

8. **Mobile app**
   - React Native / Flutter
   - Overlay in-game (Steam Workshop)

9. **Monétisation**
   - Premium features
   - Ads non-intrusives
   - Patreon/Donations

---

## 🤝 Ressources légales

### ℹ️ Important sur les droits d'auteur

**Vous POUVEZ:**
- ✅ Créer vos propres captures d'écran in-game
- ✅ Utiliser des screenshots du jeu pour du contenu éducatif
- ✅ Référencer/lier des vidéos YouTube (avec attribution)
- ✅ Citer d'autres sites avec liens et crédits

**Vous NE POUVEZ PAS:**
- ❌ Copier directement le contenu d'autres sites
- ❌ Scraper automatiquement sans permission
- ❌ Réutiliser des images sans licence appropriée
- ❌ Prétendre que le contenu tiers est le vôtre

### 📜 Mentions légales recommandées

Ajoutez dans votre footer:

```html
<p>
    CS2 Lineups n'est pas affilié à Valve Corporation. 
    Counter-Strike et le logo CS sont des marques de Valve Corporation.
    Tout le contenu de gameplay appartient à Valve Corporation.
</p>
```

---

## 💡 Conseils pour la création de contenu

### Captures d'écran de qualité

```
1. Résolution: 1920x1080 minimum
2. Format: PNG pour statique, GIF pour démos
3. Poids: < 500KB par image (optimisation)
4. Nommage: descriptif (map-zone-grenade-vue.png)
```

### Vidéos démo

```
1. Durée: 3-5 secondes max
2. FPS: 30-60 fps
3. Codec: H.264 ou WebM
4. Montrer: Position → Visée → Lancer → Impact
```

### Descriptions efficaces

```
✅ "Smoke CT depuis T Spawn - Jumpthrow bind required"
✅ "Flash Palace Pop - Run + Left-click depuis Ramp"

❌ "Smoke A"
❌ "Grenade truc"
```

---

## 🆘 Support et aide

**Problèmes courants:**

1. **Port 8080 déjà utilisé**
   ```bash
   lsof -ti:8080 | xargs kill -9
   ```

2. **Erreur de connexion PostgreSQL**
   ```bash
   sudo systemctl start postgresql
   ```

3. **Templates non trouvés**
   ```bash
   # Vérifier que vous êtes dans le bon dossier
   pwd
   ls templates/
   ```

---

## 📚 Ressources utiles

**Documentation:**
- [PostgreSQL](https://www.postgresql.org/docs/)
- [Go Templates](https://pkg.go.dev/html/template)
- [Gorilla Mux](https://github.com/gorilla/mux)

**Communauté CS2:**
- [/r/CounterStrike](https://reddit.com/r/CounterStrike)
- [/r/GlobalOffensive](https://reddit.com/r/GlobalOffensive)
- Discord CS2 France

**Outils:**
- [TinyPNG](https://tinypng.com/) - Compression d'images
- [ezgif](https://ezgif.com/) - Optimisation GIF
- [Figma](https://figma.com/) - Design UI/UX

---

## 🎉 Conclusion

Vous disposez maintenant d'une **infrastructure complète et professionnelle** pour votre site de lineups CS2 :

✅ Base de données PostgreSQL optimisée  
✅ Interface utilisateur moderne et responsive  
✅ Système de filtres avancé  
✅ 54 lineups de départ sur 7 maps  
✅ Backend Go performant  
✅ Ready pour production sur Scalingo  

**La clé du succès:** Ajoutez du contenu **progressivement** et **de qualité**. 
Mieux vaut 10 excellents lineups que 100 médiocres.

Bon courage ! 🚀
