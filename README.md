# YBoost Scalingo - Système de Recherche

Application web avec système de recherche sans base de données, déployable sur Scalingo.

## 🚀 Fonctionnalités

- ✅ Recherche en temps réel
- ✅ Filtrage par catégorie (Articles, Produits, Services)
- ✅ Surlignage des résultats
- ✅ Tri par pertinence
- ✅ Interface responsive
- ✅ Pas de base de données nécessaire (utilise JSON)

## 📦 Installation locale

```bash
# Installer les dépendances
npm install

# Lancer le serveur en local
npm start
```

L'application sera accessible sur `http://localhost:3000`

## 🌐 Déploiement sur Scalingo

### Prérequis
- Un compte Scalingo
- Git installé sur votre machine
- Scalingo CLI (optionnel mais recommandé)

### Méthode 1: Via Git

```bash
# Initialiser un dépôt Git
git init

# Ajouter tous les fichiers
git add .

# Faire un commit
git commit -m "Initial commit"

# Ajouter le remote Scalingo (remplacez <app-name> par le nom de votre app)
git remote add scalingo git@ssh.osc-fr1.scalingo.com:<app-name>.git

# Déployer
git push scalingo master
```

### Méthode 2: Via l'interface Scalingo

1. Créez une nouvelle application sur https://dashboard.scalingo.com
2. Connectez votre dépôt GitHub/GitLab
3. Scalingo détectera automatiquement Node.js et déploiera l'application

### Variables d'environnement

Aucune variable d'environnement n'est requise pour le fonctionnement de base.

## 📝 Personnalisation

### Modifier les données de recherche

Éditez le fichier `data.json` pour ajouter/modifier/supprimer des éléments :

```json
{
  "items": [
    {
      "id": 1,
      "title": "Votre titre",
      "description": "Votre description",
      "category": "article|produit|service",
      "tags": ["tag1", "tag2"],
      "url": "#lien",
      "date": "2026-01-15",
      "price": "29.99€"
    }
  ]
}
```

### Modifier le style

Éditez le fichier `styles.css` pour personnaliser les couleurs et le design.

## 🔧 Structure du projet

```
yboost_scalingo/
├── index.html      # Page principale
├── search.js       # Logique de recherche
├── styles.css      # Styles CSS
├── data.json       # Données de recherche
├── server.js       # Serveur Express
├── package.json    # Dépendances Node.js
├── Procfile        # Configuration Scalingo
└── README.md       # Ce fichier
```

## 🎨 Fonctionnement de la recherche

La recherche fonctionne en :
1. Chargeant les données du fichier JSON au démarrage
2. Filtrant en temps réel selon les critères
3. Calculant un score de pertinence
4. Triant et affichant les résultats

## 📱 Compatibilité

- ✅ Chrome, Firefox, Safari, Edge
- ✅ Mobile et tablette
- ✅ Progressive Enhancement

## 🤝 Support

Pour toute question, consultez la documentation Scalingo : https://doc.scalingo.com/

## 📄 Licence

MIT
