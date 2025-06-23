import { CurrencyUtils, StorageUtils } from './inputUtils.js';

// Function to validate individual fields
function validateField(fieldId, isRequired = true) {
    const input = document.getElementById(fieldId);
    const errorElement = document.getElementById(`error_${fieldId}`);
    input.value = CurrencyUtils.formatCurrency(input.value);
    const value = input.value;

    if (isRequired && value == "0,00") {
        errorElement.textContent = 'Este campo é obrigatório';
        errorElement.classList.remove('hidden');
        return false;
    }

    if (value !== null && Number(value) < 0) {
        errorElement.textContent = 'O valor não pode ser negativo';
        errorElement.classList.remove('hidden');
        return false;
    }

    errorElement.classList.add('hidden');
    processInputsPredict();
    StorageUtils.saveInputValues(document.getElementById('formulario_prever'));
    return true;
}
function processInputsPredict() {
    const valorFuturoInput = document.getElementById('final_value_input');
    const taxDecimalInflationInput = document.getElementById('tax_decimal_inflation_input');
    const taxaJurosInput = document.getElementById('tax_decimal_input');
    const valorInicialInput = document.getElementById('initial_value_input');
    taxDecimalInflationInput.value = CurrencyUtils.toNumber(document.getElementById('tax_decimal_inflation').value);
    valorFuturoInput.value = CurrencyUtils.toNumber(document.getElementById('final_value').value);
    taxaJurosInput.value = CurrencyUtils.toNumber(document.getElementById('tax_decimal').value);
    valorInicialInput.value = CurrencyUtils.toNumber(document.getElementById('initial_value').value) || 0;
}

// Main validation function
function validateForm() {
    const isTaxDecimalInflationValid = validateField('tax_decimal_inflation', false);
    const isValorFuturoValid = validateField('final_value');
    const isTaxaJurosValid = validateField('tax_decimal');
    const isValorInicialValid = validateField('initial_value', false);
    return isValorFuturoValid && isTaxaJurosValid && isValorInicialValid && isTaxDecimalInflationValid;
}

// Set up event listeners for real-time validation
export default function setupFormValidation() {
    const form = document.getElementById('formulario_prever');
    if (!form) return;
    processInputsPredict();
    StorageUtils.loadInputValues(form);
    // Validate on input change
    document.getElementById('final_value').addEventListener('input', () => validateField('final_value'));
    document.getElementById('tax_decimal').addEventListener('input', () => validateField('tax_decimal'));
    document.getElementById('initial_value').addEventListener('input', () => validateField('initial_value', false));
    document.getElementById('tax_decimal_inflation').addEventListener('input', () => validateField('tax_decimal_inflation', false));

    // Validate before HTMX request
    form.addEventListener('htmx:beforeRequest', function(event) {
        if (!validateForm()) {
            event.preventDefault();
            // Scroll to first error
            const firstError = document.querySelector('[id^="error_"]:not(.hidden)');
            if (firstError) {
                firstError.scrollIntoView({ behavior: 'smooth', block: 'center' });
            }
        }
    });
}
