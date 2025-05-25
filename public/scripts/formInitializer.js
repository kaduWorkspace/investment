/**
 * Form Initializer Module
 * Handles form initialization, event binding, and dynamic form updates
 */

import { DateUtils, CurrencyUtils, FormUtils, StorageUtils } from './inputUtils.js';
import setupFormValidation from './predict_form.js';

/**
 * Setup form event listeners and initialize the form
 * @param {HTMLFormElement} form - The form element to initialize
 */
export function setupFormListeners(form) {
  if (!form) return;

  // Setup form input event listener
  form.addEventListener("input", e => {
    FormUtils.handleErrorsEvent(e);
    FormUtils.processInputs(form);
    StorageUtils.saveInputValues(form);
  });

  // Setup currency input masks
  const valorAporte = document.getElementById('valor_aporte');
  const valorInicial = document.getElementById('valor_inicial');
  const valorTaxaAnual = document.getElementById('valor_taxa_anual');

  if (valorAporte) valorAporte.addEventListener('input', CurrencyUtils.handleCurrencyInput);
  if (valorInicial) valorInicial.addEventListener('input', CurrencyUtils.handleCurrencyInput);
  if (valorTaxaAnual) valorTaxaAnual.addEventListener('input', CurrencyUtils.handleCurrencyInput);

  // Setup htmx form validation
  document.body.addEventListener("htmx:configRequest", event => {
    if (event.detail.elt.id === "formulario_calcular") {
      if (!FormUtils.validateRequest(event)) {
        event.preventDefault();
      }
    }
  });

  // Setup data_final_opcoes change event
  const dataFinalOpcoes = document.getElementById('data_final_opcao');
  const dataFinalEspecificoWrapper = document.getElementById('data_especifica_wrapper');

  if (dataFinalOpcoes && dataFinalEspecificoWrapper) {
    dataFinalOpcoes.addEventListener('change', (e) => {
      e.target.value === "data_especifica"
        ? dataFinalEspecificoWrapper.classList.remove('hidden')
        : dataFinalEspecificoWrapper.classList.add('hidden');
    });
  }
}

/**
 * Initialize form with stored values and default settings
 */
export function initForm() {
  const form = document.getElementById('formulario_calcular');
  if (!form) return;

  // Load stored values
  StorageUtils.loadInputValues(form);

  // Process inputs to ensure all values are properly set
  FormUtils.processInputs(form);

  // Setup date values
  const dataFinalOpcoes = document.getElementById('data_final_opcao');
  const dataFinalEspecificoInput = document.getElementById('data_final');
  const dataFinalEspecificoWrapper = document.getElementById('data_especifica_wrapper');
  const dataInicialInput = document.getElementById("data_inicial");

  if (dataFinalOpcoes && dataFinalEspecificoInput) {
    if (dataFinalOpcoes.value !== "data_especifica") {
      const tipo = dataFinalOpcoes.value === "6" ? "meses" : "anos";
      const dataResultado = DateUtils.increment(parseInt(dataFinalOpcoes.value), tipo);
      dataFinalEspecificoInput.value = dataResultado;
    } else if (dataFinalEspecificoWrapper) {
      dataFinalEspecificoWrapper.classList.remove('hidden');
    }
  }

  if (dataInicialInput) {
    dataInicialInput.value = DateUtils.today();
  }

  // Setup form listeners
  setupFormListeners(form);

  return form;
}

// Setup listeners for page load and htmx content swaps
export function setupGlobalListeners() {
    // Initialize on page load
    if (document.getElementById('formulario_calcular')) {
        initForm();
    }

    // Handle htmx content loaded events for dynamic page updates
    document.body.addEventListener('htmx:afterSwap', (event) => {
        // Check if our form is in the swapped content
        if (event.detail.target.querySelector('#formulario_calcular')) {
            initForm();
        }
        if (event.detail.target.querySelector('#formulario_prever')) {

            setupFormValidation();
        }
    });
}

// Auto-initialize when imported
setupGlobalListeners();
