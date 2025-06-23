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
    }else{
        password.classList.remove('border-red-500');
        password.classList.add('border-green-500');
    }
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

export function setup() {
    setupPasswordValidation();
    const form = document.getElementById('signup-form');
    StorageUtils.loadInputValues(form);
    form.addEventListener("input", _ => {
        StorageUtils.saveInputValues(form);
    })
}
