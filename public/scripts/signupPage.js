import { FormUtils, StorageUtils } from './inputUtils.js';
function validatePasswords() {
    const password = document.getElementById('password');
    const passwordConfirm = document.getElementById('password_confirm');
    if (password.value && passwordConfirm.value) {
        if (password.value === passwordConfirm.value) {
            passwordConfirm.classList.remove('border-red-500');
            passwordConfirm.classList.add('border-green-500');
        } else {
            passwordConfirm.classList.remove('border-green-500');
            passwordConfirm.classList.add('border-red-500');
        }
    } else {
        passwordConfirm.classList.remove('border-red-500', 'border-green-500');
    }
    if(!FormUtils.validatePassword(password.value)){
        password.classList.remove('border-green-500');
        password.classList.add('border-red-500');
        return true
    }
    password.classList.remove('border-red-500');
    password.classList.add('border-green-500');
    return false
}
function setupPasswordValidation() {
    const password = document.getElementById('password');
    const passwordConfirm = document.getElementById('password_confirm');
    const passwordToogle = document.getElementById('toggle-password');
    const passwordConfirmToogle = document.getElementById('toggle-password-confirm');
    passwordToogle.onclick = () => FormUtils.togglePassword('password');
    passwordConfirmToogle.onclick = () => FormUtils.togglePassword('password_confirm');
    password.addEventListener('input', validatePasswords);
    passwordConfirm.addEventListener('input', validatePasswords);
}
function validateForm(e) {
    e.preventDefault();
    document.getElementById('email').value = document.getElementById('email').value.toLowerCase();
    if (validatePasswords()) {
        e.submit()
    }
}
export function setup() {
    setupPasswordValidation();
    const form = document.getElementById('signup-form');
    StorageUtils.loadInputValues(form);
    form.addEventListener("input", _ => {
        StorageUtils.saveInputValues(form);
    })
    form.addEventListener("submit", validateForm)
}
