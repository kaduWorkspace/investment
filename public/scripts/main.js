/**
    * Main application entry point
    * Import necessary modules based on the current page
    */

    // Import core functionality that's needed on all pages
import './formInitializer.js';

// Conditionally load page-specific modules
document.addEventListener('DOMContentLoaded', () => {
    // Check which page we're on and load appropriate modules
    if (document.getElementById('form_container')) {
        // Dynamically import the financas page module only when needed
        import('./financasPage.js')
            .then(() => console.log('Financas page module loaded'))
            .catch(err => console.error('Error loading financas module:', err));
    }
});

// Handle htmx page transitions
document.body.addEventListener('htmx:afterSwap', (event) => {
    // Check if we've loaded the financas page through htmx
    if (event.detail.target.querySelector('#form_container')) {
        // Dynamically import the financas page module
        import('./financasPage.js')
            .then(() => console.log('Financas page module loaded after htmx swap'))
            .catch(err => console.error('Error loading financas module:', err));
    }
});
