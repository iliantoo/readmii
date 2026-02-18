// Reader functionality
let currentPage = 1;
const pageCounter = document.getElementById('currentPage');
const pagesContainer = document.getElementById('pagesContainer');

// Track visible pages with Intersection Observer
if (pagesContainer) {
    const pageObserver = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting && entry.intersectionRatio > 0.5) {
                const pageNum = parseInt(entry.target.dataset.page);
                updateCurrentPage(pageNum);
            }
        });
    }, {
        threshold: 0.5,
        rootMargin: '-100px 0px -100px 0px'
    });

    // Observe all pages
    document.querySelectorAll('.manga-page').forEach(page => {
        pageObserver.observe(page);
    });
}

// Update current page counter
function updateCurrentPage(pageNum) {
    if (currentPage !== pageNum) {
        currentPage = pageNum;
        if (pageCounter) {
            pageCounter.textContent = pageNum;
        }
        
        // Save reading progress
        saveReadingProgress(pageNum);
    }
}

// Save reading progress to localStorage
function saveReadingProgress(pageNum) {
    const path = window.location.pathname;
    const match = path.match(/\/read\/([^\/]+)\/(\d+)/);
    
    if (match) {
        const mangaId = match[1];
        const chapterNum = match[2];
        const key = `progress_${mangaId}_${chapterNum}`;
        
        localStorage.setItem(key, pageNum.toString());
    }
}

// Load reading progress
function loadReadingProgress() {
    const path = window.location.pathname;
    const match = path.match(/\/read\/([^\/]+)\/(\d+)/);
    
    if (match) {
        const mangaId = match[1];
        const chapterNum = match[2];
        const key = `progress_${mangaId}_${chapterNum}`;
        const savedPage = localStorage.getItem(key);
        
        if (savedPage && parseInt(savedPage) > 1) {
            const pageElement = document.querySelector(`[data-page="${savedPage}"]`);
            if (pageElement) {
                // Scroll to saved page after a short delay
                setTimeout(() => {
                    pageElement.scrollIntoView({ behavior: 'smooth', block: 'start' });
                }, 500);
            }
        }
    }
}

// Keyboard navigation for reader
document.addEventListener('keydown', (e) => {
    // Don't trigger if user is typing in an input
    if (e.target.matches('input, textarea')) return;
    
    const pages = document.querySelectorAll('.manga-page');
    if (!pages.length) return;
    
    switch(e.key) {
        case 'ArrowDown':
        case 'PageDown':
        case 's':
        case 'S':
            e.preventDefault();
            scrollToNextPage('down');
            break;
            
        case 'ArrowUp':
        case 'PageUp':
        case 'w':
        case 'W':
            e.preventDefault();
            scrollToNextPage('up');
            break;
            
        case 'Home':
            e.preventDefault();
            pages[0].scrollIntoView({ behavior: 'smooth' });
            break;
            
        case 'End':
            e.preventDefault();
            pages[pages.length - 1].scrollIntoView({ behavior: 'smooth' });
            break;
            
        case 'ArrowLeft':
            e.preventDefault();
            navigateChapter('prev');
            break;
            
        case 'ArrowRight':
            e.preventDefault();
            navigateChapter('next');
            break;
    }
});

// Scroll to next/previous page
function scrollToNextPage(direction) {
    const pages = Array.from(document.querySelectorAll('.manga-page'));
    const currentPageElement = pages.find(p => parseInt(p.dataset.page) === currentPage);
    
    if (!currentPageElement) return;
    
    const currentIndex = pages.indexOf(currentPageElement);
    let targetIndex;
    
    if (direction === 'down') {
        targetIndex = Math.min(currentIndex + 1, pages.length - 1);
    } else {
        targetIndex = Math.max(currentIndex - 1, 0);
    }
    
    pages[targetIndex].scrollIntoView({ behavior: 'smooth', block: 'start' });
}

// Navigate to previous/next chapter
function navigateChapter(direction) {
    const link = document.querySelector(`.nav-chapter.${direction === 'prev' ? 'prev' : 'next'}`);
    if (link && !link.classList.contains('disabled')) {
        link.click();
    }
}

// Image preloading for better performance
function preloadImages() {
    const images = document.querySelectorAll('.manga-page img');
    let loadedCount = 0;
    
    images.forEach((img, index) => {
        // Preload next 2-3 images
        if (index <= currentPage + 2) {
            const preloadImg = new Image();
            preloadImg.onload = () => {
                loadedCount++;
                if (loadedCount === Math.min(3, images.length)) {
                    console.log('✅ Images préchargées');
                }
            };
            preloadImg.src = img.src;
        }
    });
}

// Initialize reader
if (pagesContainer) {
    loadReadingProgress();
    preloadImages();
    
    // Show keyboard shortcuts hint
    console.log('⌨️ Raccourcis clavier:');
    console.log('  ↑/↓ ou W/S - Page précédente/suivante');
    console.log('  ←/→ - Chapitre précédent/suivant');
    console.log('  Home/End - Première/dernière page');
    console.log('  Esc - Retour');
    console.log('  T - Changer le thème');
}

// Add swipe support for mobile
let touchStartY = 0;
let touchEndY = 0;

if (pagesContainer) {
    pagesContainer.addEventListener('touchstart', (e) => {
        touchStartY = e.touches[0].clientY;
    }, { passive: true });
    
    pagesContainer.addEventListener('touchend', (e) => {
        touchEndY = e.changedTouches[0].clientY;
        handleSwipe();
    }, { passive: true });
}

function handleSwipe() {
    const swipeThreshold = 50;
    const diff = touchStartY - touchEndY;
    
    if (Math.abs(diff) > swipeThreshold) {
        if (diff > 0) {
            // Swipe up - next page
            scrollToNextPage('down');
        } else {
            // Swipe down - previous page
            scrollToNextPage('up');
        }
    }
}

console.log('📖 Reader initialized!');
