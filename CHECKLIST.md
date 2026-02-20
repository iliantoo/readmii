# ✅ Checklist - Prochaines étapes

## 🎯 Étapes immédiates (Aujourd'hui)

- [ ] **Tester le site localement**
  - Aller sur http://localhost:8080
  - Naviguer entre les pages
  - Tester les filtres
  - Vérifier que tout fonctionne

- [ ] **Comprendre la structure**
  - Lire [GUIDE_COMPLET.md](GUIDE_COMPLET.md)
  - Explorer seed_data.json
  - Regarder les templates HTML

## 📸 Contenu prioritaire (Cette semaine)

- [ ] **Capturer 3-5 lineups essentiels pour Mirage**
  - [ ] Xbox Smoke T Spawn
  - [ ] CT Smoke T Spawn
  - [ ] Window Smoke T Spawn
  - [ ] Jungle Smoke
  - [ ] Stairs Smoke

  Pour chaque lineup:
  - [ ] Position du joueur (screenshot)
  - [ ] Placement du viseur (screenshot zoomé)
  - [ ] Démo du lancer (GIF 3-5 sec)

- [ ] **Optimiser les images**
  - [ ] Compresser avec TinyPNG
  - [ ] Renommer proprement (map-zone-grenade-type.jpg)
  - [ ] Placer dans public/assets/lineups/mirage/

- [ ] **Mettre à jour seed_data.json**
  - [ ] Remplacer les chemins placeholder par les vrais chemins
  - [ ] Réimporter dans PostgreSQL

## 🚀 Déploiement (Cette semaine)

- [ ] **Préparer pour Scalingo**
  - [ ] Vérifier que tous les fichiers sont commités
  - [ ] Tester en local une dernière fois
  - [ ] Ajouter un Procfile si nécessaire

- [ ] **Déployer sur Scalingo**
  ```bash
  git add .
  git commit -m "Site CS2 Lineups - Version 1.0"
  git push scalingo main
  ```

- [ ] **Configurer PostgreSQL en production**
  ```bash
  scalingo -a votre-app addons-add postgresql postgresql-starter-512
  scalingo -a votre-app run bash
  # Dans le shell Scalingo
  go run cmd/seed/seed.go
  ```

- [ ] **Tester en production**
  - [ ] Ouvrir votre-app.osc-fr1.scalingo.io
  - [ ] Vérifier toutes les pages
  - [ ] Tester les filtres

## 📝 Contenu additionnel (Semaine 2)

- [ ] **Compléter Dust II**
  - [ ] 5-10 lineups avec images
  - [ ] Mettre à jour la base

- [ ] **Compléter Inferno**
  - [ ] 5-10 lineups avec images
  - [ ] Mettre à jour la base

- [ ] **Améliorer les descriptions**
  - [ ] Relire chaque lineup
  - [ ] Ajouter des détails tactiques
  - [ ] Vérifier l'orthographe

## 🎨 Améliorations UI (Semaine 3)

- [ ] **Ajouter de vraies images de maps**
  - [ ] Créer ou trouver des minimap HD
  - [ ] Remplacer les emojis 🗺️ par de vraies images
  - [ ] public/assets/maps/mirage.jpg, etc.

- [ ] **Améliorer les badges de grenades**
  - [ ] Icônes custom au lieu d'emojis
  - [ ] Couleurs plus précises
  - [ ] Animations au hover

- [ ] **Mode clair optionnel**
  - [ ] Implémenter le toggle theme
  - [ ] CSS pour light mode
  - [ ] LocalStorage pour préférence

## 🔧 Fonctionnalités avancées (Semaine 4+)

- [ ] **Système de recherche**
  - [ ] Barre de recherche globale
  - [ ] Recherche par zone (A, B, Mid)
  - [ ] Recherche par texte

- [ ] **Statistiques**
  - [ ] Google Analytics
  - [ ] Tracking des lineups populaires
  - [ ] Heatmap des clics

- [ ] **Interface d'administration**
  - [ ] Route /admin protégée
  - [ ] CRUD des lineups
  - [ ] Upload d'images

- [ ] **Fonctionnalités sociales**
  ```
  - [ ] Système de votes (👍/👎)
  - [ ] Commentaires
  - [ ] Notes de difficulté par la communauté
  - [ ] Partage sur réseaux sociaux
  ```

## 💼 Marketing et croissance

- [ ] **SEO**
  - [ ] Meta tags optimisés
  - [ ] Descriptions uniques par page
  - [ ] Sitemap.xml
  - [ ] robots.txt

- [ ] **Réseaux sociaux**
  - [ ] Créer un Twitter/X du projet
  - [ ] Poster sur /r/GlobalOffensive (avec permission mods)
  - [ ] Partager dans Discord CS2 France
  - [ ] TikTok/YouTube Shorts avec démos

- [ ] **Partenariats**
  - [ ] Contacter des streamers CS2 FR
  - [ ] Proposer des échanges visibilité
  - [ ] Collaborations avec créateurs de contenu

## 📊 Métriques de succès

**Objectif Mois 1:**
- [ ] 100 visiteurs uniques
- [ ] 10 lineups complets (images + GIF)
- [ ] 3 maps complètement documentées
- [ ] Temps de chargement < 2 secondes

**Objectif Mois 3:**
- [ ] 1000 visiteurs uniques
- [ ] 50 lineups complets
- [ ] 7 maps complètes
- [ ] 1er retour communautaire positif

**Objectif Mois 6:**
- [ ] 5000 visiteurs uniques
- [ ] 100+ lineups
- [ ] Top 3 Google pour "lineups CS2 [mapname]"
- [ ] Interface admin fonctionnelle

## 🐛 Bugs connus à corriger

- [ ] Vérifier compatibilité mobile
- [ ] Tester sur Safari/Firefox/Edge
- [ ] Optimiser temps de chargement
- [ ] Ajouter loading states

## 💡 Idées futures

**Ne pas oublier:**
- Mini-map interactive avec marqueurs de position
- Mode "entraînement" avec timer pour mémoriser
- Export PDF des lineups favoris
- API publique pour apps tierces
- Widget embeddable pour intégrer sur d'autres sites
- Mode "dark/light" auto basé sur heure de la journée
- Traduction EN/FR/ES
- Mobile app companion

---

## 🎯 Focus de la semaine

**Semaine actuelle:** _______________

**Priorité TOP 3:**
1. [ ] __________________________________
2. [ ] __________________________________
3. [ ] __________________________________

**Temps estimé:** ___ heures

**Bloqueurs:** 
- _______________________________________
- _______________________________________

**Notes:**
```
_______________________________________________
_______________________________________________
_______________________________________________
```

---

## 📞 Besoin d'aide?

Si vous bloquez sur quelque chose:

1. **Relire le [GUIDE_COMPLET.md](GUIDE_COMPLET.md)**
2. **Vérifier les logs:** `tail -f logs/server.log`
3. **Tester en local avant de déployer**
4. **Google est votre ami** (Stack Overflow, Reddit, etc.)

Bon courage! Vous avez une base solide, maintenant il faut peupler avec du contenu de qualité! 🚀
