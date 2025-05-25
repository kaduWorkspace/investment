/**
    * Financas Page Specific Functionality
    * Contains code specific to the financas page
    */

    import { initForm } from './formInitializer.js';

/**
    * Creates a map of form elements by name for easier access
    * @param {HTMLFormElement} form - The form to map
    * @returns {Object} Map of form elements keyed by name
    */
    function mapFormInputs(form) {
        if (!form) return {};

        const inputsPossiveis = [...form.elements].filter(input => !input.dataset.ignore_input);
        return inputsPossiveis.reduce((acc, curr) => {
            acc[curr.name] = curr;
            return acc;
        }, {});
    }

/**
    * Initialize financas-specific functionality
    */
    function initFinancasPage() {
        // Initialize the form using the shared initializer
        const form = initForm();
        if (!form) return;

        // Map the form inputs by name for easy access in financas-specific code
        const inputsPorNome = mapFormInputs(form);

        // Add any financas-specific initialization here
        // ...

            console.log('Financas page initialized');

        return {
            form,
            inputsPorNome
        };
    }

// Initialize the financas page
// This only runs if this module is imported on the financas page
initFinancasPage();
