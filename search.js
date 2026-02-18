// Variables globales
let allData = [];
let filteredData = [];
let currentCategory = 'all';

// Chargement des données au démarrage
document.addEventListener('DOMContentLoaded', async () => {
    await loadData();
    displayAllResults();
    setupEventListeners();
});

// Charger les données depuis le fichier JSON
async function loadData() {
    try {
        const response = await fetch('data.json');
        const data = await response.json();
        allData = data.items;
        filteredData = allData;
    } catch (error) {
        console.error('Erreur lors du chargement des données:', error);
        showError('Impossible de charger les données');
    }
}

// Configuration des événements
function setupEventListeners() {
    const searchInput = document.getElementById('searchInput');
    const filterButtons = document.querySelectorAll('.filter-btn');

    // Recherche en temps réel
    searchInput.addEventListener('input', (e) => {
        performSearch(e.target.value);
    });

    // Filtres par catégorie
    filterButtons.forEach(btn => {
        btn.addEventListener('click', () => {
            // Mise à jour visuelle des boutons
            filterButtons.forEach(b => b.classList.remove('active'));
            btn.classList.add('active');
            
            // Filtrage
            currentCategory = btn.dataset.category;
            performSearch(searchInput.value);
        });
    });
}

// Fonction de recherche principale
function performSearch(query) {
    const searchTerm = query.toLowerCase().trim();
    
    // Filtrer par catégorie
    let results = currentCategory === 'all' 
        ? [...allData] 
        : allData.filter(item => item.category === currentCategory);
    
    // Filtrer par terme de recherche
    if (searchTerm) {
        results = results.filter(item => {
            const titleMatch = item.title.toLowerCase().includes(searchTerm);
            const descMatch = item.description.toLowerCase().includes(searchTerm);
            const tagsMatch = item.tags.some(tag => tag.toLowerCase().includes(searchTerm));
            
            return titleMatch || descMatch || tagsMatch;
        });
    }
    
    // Tri par pertinence
    if (searchTerm) {
        results = results.sort((a, b) => {
            const aScore = calculateRelevance(a, searchTerm);
            const bScore = calculateRelevance(b, searchTerm);
            return bScore - aScore;
        });
    }
    
    filteredData = results;
    displayResults(results, searchTerm);
}

// Calcul de la pertinence
function calculateRelevance(item, searchTerm) {
    let score = 0;
    
    const titleLower = item.title.toLowerCase();
    const descLower = item.description.toLowerCase();
    
    // Le titre a plus de poids
    if (titleLower.includes(searchTerm)) score += 10;
    if (titleLower.startsWith(searchTerm)) score += 5;
    
    // La description a un poids moyen
    if (descLower.includes(searchTerm)) score += 5;
    
    // Les tags ont un poids faible
    item.tags.forEach(tag => {
        if (tag.toLowerCase().includes(searchTerm)) score += 3;
    });
    
    return score;
}

// Afficher tous les résultats
function displayAllResults() {
    displayResults(allData, '');
}

// Afficher les résultats
function displayResults(results, searchTerm) {
    const resultsContainer = document.getElementById('results');
    const noResultsDiv = document.getElementById('noResults');
    const statsDiv = document.getElementById('searchStats');
    
    // Mise à jour des statistiques
    const categoryText = currentCategory === 'all' ? 'toutes catégories' : currentCategory + 's';
    statsDiv.textContent = `${results.length} résultat${results.length > 1 ? 's' : ''} trouvé${results.length > 1 ? 's' : ''} (${categoryText})`;
    
    // Affichage
    if (results.length === 0) {
        resultsContainer.innerHTML = '';
        noResultsDiv.style.display = 'block';
        return;
    }
    
    noResultsDiv.style.display = 'none';
    resultsContainer.innerHTML = results.map(item => createResultCard(item, searchTerm)).join('');
}

// Créer une carte de résultat
function createResultCard(item, searchTerm) {
    const categoryEmoji = {
        'article': '📝',
        'produit': '📦',
        'service': '⚙️'
    };
    
    const highlightedTitle = highlightText(item.title, searchTerm);
    const highlightedDesc = highlightText(item.description, searchTerm);
    
    const priceHTML = item.price 
        ? `<div class="price">${item.price}</div>` 
        : '';
    
    const dateHTML = item.date 
        ? `<div class="date">${formatDate(item.date)}</div>` 
        : '';
    
    const tagsHTML = item.tags
        .map(tag => `<span class="tag">${highlightText(tag, searchTerm)}</span>`)
        .join('');
    
    return `
        <div class="result-card" data-category="${item.category}">
            <div class="card-header">
                <span class="category-badge">${categoryEmoji[item.category]} ${item.category}</span>
                ${priceHTML}
                ${dateHTML}
            </div>
            <h3 class="result-title">
                <a href="${item.url}">${highlightedTitle}</a>
            </h3>
            <p class="result-description">${highlightedDesc}</p>
            <div class="tags">
                ${tagsHTML}
            </div>
        </div>
    `;
}

// Surligner le texte recherché
function highlightText(text, searchTerm) {
    if (!searchTerm) return text;
    
    const regex = new RegExp(`(${escapeRegex(searchTerm)})`, 'gi');
    return text.replace(regex, '<mark>$1</mark>');
}

// Échapper les caractères spéciaux regex
function escapeRegex(str) {
    return str.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
}

// Formater la date
function formatDate(dateString) {
    const date = new Date(dateString);
    const options = { year: 'numeric', month: 'long', day: 'numeric' };
    return date.toLocaleDateString('fr-FR', options);
}

// Afficher une erreur
function showError(message) {
    const resultsContainer = document.getElementById('results');
    resultsContainer.innerHTML = `
        <div class="error-message">
            <p>❌ ${message}</p>
        </div>
    `;
}
