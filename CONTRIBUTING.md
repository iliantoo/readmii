# 🤝 Guide de contribution - CS2 Lineups

Merci de votre intérêt pour améliorer CS2 Lineups ! Toutes les contributions sont les bienvenues.

## 📋 Comment contribuer

### 1. Types de contributions acceptées

- 🎯 **Nouveaux lineups** - Smokes, flashs, molotovs, HE grenades
- 🐛 **Corrections de bugs** - Rapports et fixes
- ✨ **Nouvelles fonctionnalités** - Filtres, recherche, etc.
- 📝 **Documentation** - Guides, tutoriels, traductions
- 🎨 **Design** - Améliorations UI/UX
- 🌐 **Traductions** - Multi-langue (EN, ES, DE, etc.)

### 2. Soumettre un lineup

Pour soumettre un nouveau lineup:

#### Via GitHub Issue

1. Créez une [nouvelle issue](issues/new)
2. Titre: `[Lineup] Map - Type - Description courte`
3. Incluez:
   - **Map**: Mirage, Dust II, etc.
   - **Type de grenade**: Smoke, Flash, Molotov, ou HE
   - **Side**: T ou CT
   - **Description**: "Smoke A site depuis spawn"
   - **Screenshots**: Position + crosshair (min 2 images)
   - **GIF/Video**: Démo du lineup (optionnel mais recommandé)
   - **Bind console**: Si disponible
   - **Difficulté**: Easy, Medium, ou Hard
   - **Tags**: rush, retake, execute, etc.

#### Via Pull Request

1. Forkez le repository
2. Créez une branche: `git checkout -b lineup/mirage-smoke-ct`
3. Ajoutez les médias dans `public/assets/lineups/`
4. Éditez `seed_data.json` pour ajouter votre lineup
5. Testez localement: `go run cmd/seed/seed.go`
6. Committez: `git commit -m "Add: Mirage CT smoke"`
7. Push: `git push origin lineup/mirage-smoke-ct`
8. Ouvrez une Pull Request

### 3. Standards de qualité pour les lineups

#### Images

- **Format**: PNG ou JPG
- **Résolution minimale**: 1920x1080
- **Poids**: < 500 KB par image
- **Nommage**: `{map}_{type}_{id}_position.png` et `{map}_{type}_{id}_crosshair.png`
- **Pas de watermark** (sauf les vôtres)

#### GIFs

- **Format**: GIF ou MP4
- **Résolution**: 1280x720 minimum
- **Durée**: 3-10 secondes
- **Poids**: < 5 MB
- **FPS**: 30-60
- **Nommage**: `{map}_{type}_{id}_demo.gif`

#### Informations requises

```json
{
  "id": "mirage-smoke-ct-01",
  "title": "Smoke CT depuis spawn T",
  "grenade_type": "smoke",
  "side": "T",
  "difficulty": "easy",
  "position_image": "/assets/lineups/mirage_smoke_01_position.png",
  "crosshair_image": "/assets/lineups/mirage_smoke_01_crosshair.png",
  "demo_gif": "/assets/lineups/mirage_smoke_01_demo.gif",
  "bind_command": "bind v \"+jump; -attack; -jump\"",
  "action_details": "Depuis spawn T, viser le coin du toit",
  "tags": ["execute", "standard", "essential"]
}
```

### 4. Soumettre un bug

1. Vérifiez que le bug n'existe pas déjà dans [les issues](issues)
2. Créez une nouvelle issue avec le label `bug`
3. Incluez:
   - **Description claire** du problème
   - **Étapes pour reproduire** le bug
   - **Comportement attendu** vs **comportement actuel**
   - **Screenshots** si applicable
   - **Environnement**: OS, navigateur, version

### 5. Proposer une fonctionnalité

1. Créez une issue avec le label `enhancement`
2. Décrivez:
   - **Le problème** que ça résout
   - **La solution proposée**
   - **Alternatives considérées**
   - **Mockups/wireframes** si applicable

### 6. Guidelines de code

#### Go

- Suivez les [Effective Go](https://golang.org/doc/effective_go) guidelines
- Utilisez `gofmt` pour formater le code
- Ajoutez des commentaires pour les fonctions publiques
- Testez vos changements: `go test ./...`

#### HTML/CSS

- Indentation: 4 espaces
- Noms de classes: kebab-case (`lineup-card`, `filter-btn`)
- Variables CSS: `--nom-variable`
- Mobile-first responsive

#### JavaScript

- ES6+ moderne
- Pas de jQuery (vanilla JS uniquement)
- Commentaires pour la logique complexe

### 7. Process de Pull Request

1. **Avant de soumettre**:
   - [ ] Code formaté (`gofmt`)
   - [ ] Tests passent (`go test`)
   - [ ] Serveur démarre sans erreur
   - [ ] Testé dans navigateur (Chrome + Firefox minimum)
   - [ ] Documentation mise à jour si nécessaire

2. **Description de la PR**:
   - Résumé clair des changements
   - Issue liée (fixes #123)
   - Screenshots si changements visuels
   - Notes pour les reviewers

3. **Review**:
   - Un mainteneur reviewera votre PR
   - Répondez aux commentaires et suggestions
   - Faites les ajustements demandés

4. **Merge**:
   - Une fois approuvée, votre PR sera mergée
   - Elle apparaîtra dans la prochaine release

### 8. Commit Messages

Format:
```
Type: Description courte (50 caractères max)

Corps du message optionnel avec plus de détails.
Peut être sur plusieurs lignes.

Fixes #123
```

**Types**:
- `Add:` Nouveau lineup, fonctionnalité
- `Fix:` Correction de bug
- `Update:` Mise à jour de contenu existant
- `Refactor:` Refactorisation sans changement de fonctionnalité
- `Docs:` Documentation uniquement
- `Style:` CSS, formatage
- `Perf:` Amélioration de performance
- `Test:` Ajout/modification de tests

**Exemples**:
```
Add: Mirage CT smoke from stairs

Fix: Filter buttons not highlighting on mobile

Update: Dust2 lineup images with higher quality

Docs: Add deployment guide for Scalingo
```

### 9. Code of Conduct

- ✅ Soyez respectueux et professionnel
- ✅ Acceptez les critiques constructives
- ✅ Concentrez-vous sur ce qui est mieux pour la communauté
- ❌ Pas de harcèlement, discrimination, ou comportement toxique
- ❌ Pas de spam ou auto-promotion excessive

### 10. Licence

En contribuant, vous acceptez que vos contributions soient sous licence MIT.

Tout le contenu de gameplay (screenshots, GIFs) doit respecter les droits de Valve Corporation.

### 11. Questions ?

- 💬 Ouvrez une [Discussion](discussions)
- 📧 Email: [votre-email]
- 💭 Discord: [Lien serveur] (à venir)

---

## 🏆 Contributeurs

Merci à tous les contributeurs qui améliorent ce projet !

[Liste des contributeurs](contributors)

---

## 🚀 Roadmap

Consultez notre [roadmap](CHECKLIST.md) pour voir les fonctionnalités prévues.

Si vous voulez contribuer mais ne savez pas par où commencer, regardez les issues avec le label `good first issue`.

---

**Merci de contribuer à CS2 Lineups ! 🎮**
