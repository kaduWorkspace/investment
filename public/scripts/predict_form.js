import { CurrencyUtils } from './inputUtils.js';

// Function to validate individual fields
function validateField(fieldId, isRequired = true) {
    const input = document.getElementById(fieldId);
    const errorElement = document.getElementById(`error_${fieldId}`);
    input.value = CurrencyUtils.formatCurrency(input.value);
    const {value} = input;

    if (isRequired && (!value && value !== 0)) {
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
    return true;
}
function processInputsPredict() {
    const valorFuturoInput = document.getElementById('valor_futuro_input');
    const taxaJurosInput = document.getElementById('taxa_juros_anual_input');
    const valorInicialInput = document.getElementById('valor_inicial_input');
    valorFuturoInput.value = CurrencyUtils.toNumber(document.getElementById('valor_futuro').value);
    taxaJurosInput.value = CurrencyUtils.toNumber(document.getElementById('taxa_juros_anual').value);
    valorInicialInput.value = CurrencyUtils.toNumber(document.getElementById('valor_inicial').value) || 0;
}

// Main validation function
function validateForm() {
    const isValorFuturoValid = validateField('valor_futuro');
    const isTaxaJurosValid = validateField('taxa_juros_anual');
    const isValorInicialValid = validateField('valor_inicial', false);

    return isValorFuturoValid && isTaxaJurosValid && isValorInicialValid;
}

// Set up event listeners for real-time validation
export default function setupFormValidation() {
    const form = document.getElementById('formulario_prever');
    if (!form) return;
    processInputsPredict();
    // Validate on input change
    document.getElementById('valor_futuro').addEventListener('input', () => validateField('valor_futuro'));
    document.getElementById('taxa_juros_anual').addEventListener('input', () => validateField('taxa_juros_anual'));
    document.getElementById('valor_inicial').addEventListener('input', () => validateField('valor_inicial', false));

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
