/**
    * Input Utilities Module
    * Core utility functions for input validation and formatting
    */

    // Date formatting functions
export const DateUtils = {
    format(date) {
        const ano = date.getFullYear();
        const mes = String(date.getMonth() + 1).padStart(2, "0");
        const dia = String(date.getDate()).padStart(2, "0");
        return `${dia}/${mes}/${ano}`;
    },

    increment(quantidade = 6, tipo = "meses") {
        const data = new Date();
        if (tipo === "meses") {
            data.setMonth(data.getMonth() + quantidade);
        } else {
            data.setFullYear(data.getFullYear() + quantidade);
        }
        return this.format(data);
    },

    today() {
        return this.format(new Date());
    }
};

// Number and currency formatting functions
export const CurrencyUtils = {
    removeZerosAtStart(str) {
        const regex = /^0*(\d+)$/;
        const match = str.match(regex);
        return match ? match[1] : str;
    },

    removeMask(value) {
        return value.replace(/\D/g, '') || '0';
    },

    toNumber(value) {
        return Number(Number(value.replaceAll('.', '').replaceAll(',', '.').replaceAll('R$', '').trim()).toFixed(2)) || 0;
    },

    toMonetaryValue(number) {
        return new Intl.NumberFormat('pt-BR', { style: 'currency', currency: 'BRL' }).format(number);
    },

    formatCurrency(valor) {
        valor = this.removeMask(valor);
        valor = this.removeZerosAtStart(valor);
        const valorSplit = valor.split('');

        // Handle special cases
        if (valorSplit.length === 0) return "0,00";
        if (valorSplit.length === 1) return `0,0${valorSplit[0]}`;
        if (valorSplit.length === 2) return `0,${valorSplit[0]}${valorSplit[1]}`;
        if (valorSplit.length === 3) return `${valorSplit[0]},${valorSplit[1]}${valorSplit[2]}`;

        // Handle longer numbers with thousand separators
        const centavos = [valorSplit.pop(), valorSplit.pop()].reverse().join("");
        const gruposDeTres = [[]];
        let contGrupo = 0;

        // Handle numbers that don't divide evenly into groups of 3
        const resto = valorSplit.length % 3;
        const primeiros = [];
        for (let i = 0; i < resto; i++) {
            primeiros.push(valorSplit.shift());
        }

        // Group remaining digits into groups of 3
        valorSplit.forEach((digit, index) => {
            if (gruposDeTres[contGrupo].length === 3) {
                gruposDeTres.push([]);
                contGrupo++;
            }
            gruposDeTres[contGrupo].push(digit);
        });

        const valorFormatado = gruposDeTres.map(grupo => grupo.join('')).join('.');

        // Assemble the final formatted string
        if (primeiros.length) {
            if (valorFormatado !== "") {
                return primeiros.join("") + "." + valorFormatado + "," + centavos;
            } else {
                return primeiros.join("") + "," + centavos;
            }
        } else {
            return valorFormatado + "," + centavos;
        }
    },

    handleCurrencyInput(e) {
        e.target.value = CurrencyUtils.formatCurrency(e.target.value);
    }
};

// Validation functions
export const Validators = {
    validateInputs(validarNull = false) {
        const erros = [];

        // Get DOM elements only when needed
        const valorInicial = document.getElementById('initial_value');
        const valorAporte = document.getElementById('contribution');
        const taxInflation = document.getElementById('tax_decimal_inflation');
        //const dataFinal = document.getElementById('data_final');

        if (!valorInicial || !valorAporte) {
            console.error('Required form elements not found');
            return [['error_general', 'Form elements not found']];
        }

        const valorInicialNum = CurrencyUtils.toNumber(valorInicial.value);
        const valorAporteNum = CurrencyUtils.toNumber(valorAporte.value);

        // Validate initial value
        if (valorInicialNum < 0 || (validarNull && ["", null, false].includes(valorInicial.value.trim()))) {
            erros.push(["error_initial_value", "Valor inicial inválido"]);
        }

        // Validate monthly contribution
        if (valorAporteNum > 1000000000) {
            erros.push(["error_contribution", "Aporte mensal muito alto"]);
        }

        // Validate that at least one value is provided
        if (!(valorAporteNum + valorInicialNum > 0)) {
            erros.push(
                ["error_nitial_value", "O valor inicial ou valor de aporte devem ser preenchidos!"],
                ["error_contribution", "O valor inicial ou valor de aporte devem ser preenchidos!"]
            );
        }

        // Validate end date
        /*if (!dataFinal.value) {
            erros.push(["error_data_final", "Data final inválida"]);
        }*/
        return erros.length ? erros : false;
    }
};

// Storage functions
export const StorageUtils = {
    saveInputValues(form) {
        if (!form) return;

        const inputs = form.querySelectorAll('input');
        inputs.forEach(input => {
            sessionStorage.setItem(input.name || input.id, input.value);
        });
    },

    loadInputValues(form) {
        if (!form) return;

        const inputs = form.querySelectorAll('input');
        inputs.forEach(input => {
            const storedValue = sessionStorage.getItem(input.name || input.id);
            if (input.name == "_csrf") return;
            if (storedValue) {
                input.value = storedValue;
            }
        });
    }
};

// Form processing functions
export const FormUtils = {
    processInputs(form) {
        if (!form) return;

        const inputsPossiveis = [...form.elements].filter(input => !input.dataset.ignore_input);

        // Set default values for empty number inputs
        inputsPossiveis.filter(input => input.type === "number").forEach(input => {
            if (input.value === "") {
                input.value = 0.0;
                console.log(`Empty input detected for ${input.name}, defaulting to 0.0`);
            }
        });

        // Get DOM elements
        const valorTaxaAnual = document.getElementById('tax_decimal');
        const valorTaxaAnualInput = document.getElementById('tax_decimal_input');
        const valorAporte = document.getElementById('contribution');
        const valorAporteInput = document.getElementById('contribution_input');
        const valorInicial = document.getElementById('initial_value');
        const valorInicialInput = document.getElementById('initial_value_input');
        const dataFinalOpcoes = document.getElementById('periods');
        const dataFinalEspecificoInput = document.getElementById('data_final');
        const taxDecimalInflation = document.getElementById('tax_decimal_inflation');
        const taxDecimalInflationInput = document.getElementById('tax_decimal_inflation_input');

        if (!valorTaxaAnual || !valorTaxaAnualInput || !valorAporte || !valorAporteInput ||
            !valorInicial || !valorInicialInput || !dataFinalOpcoes || !taxDecimalInflationInput) {
            console.error('Required form elements not found');
            return;
        }

        // Convert input values to numbers
        const taxaAnualValue = CurrencyUtils.toNumber(valorTaxaAnual.value);
        const valorAporteValue = CurrencyUtils.toNumber(valorAporte.value);
        const valorInicialValue = CurrencyUtils.toNumber(valorInicial.value);
        const taxaAnualInflationValue = CurrencyUtils.toNumber(taxDecimalInflation.value);

        // Update hidden inputs with calculated values
        valorAporteInput.value = valorAporteValue || 0;
        valorTaxaAnualInput.value = taxaAnualValue || 0;
        valorInicialInput.value = valorInicialValue || 0;
        taxDecimalInflationInput.value = taxaAnualInflationValue || 0;
    },

    handleErrorsEvent(e) {
        const errorSpan = document.getElementById(`error_${e.target.id}`);
        if (!errorSpan) return;

        errorSpan.classList.add('hidden');

        const validacao = Validators.validateInputs(false);
        if (!validacao) return;

        const validacaoInput = validacao.find(([errorSpanTargetName]) => errorSpanTargetName === errorSpan.id);
        if (!validacaoInput) return;

        errorSpan.innerText = validacaoInput[1];
        errorSpan.classList.remove('hidden');
    },

    validateRequest(event) {
        const validacoes = Validators.validateInputs(true);
        if (validacoes) {
            validacoes.forEach(validacao => {
                const errorSpan = document.getElementById(validacao[0]);
                if (errorSpan) {
                    errorSpan.innerText = validacao[1];
                    errorSpan.classList.remove('hidden');
                }
            });
            return false;
        }
        return true;
    }
};
