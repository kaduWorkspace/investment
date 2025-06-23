import { FormUtils, StorageUtils } from './inputUtils.js';

function setupPasswordValidation() {
    const password = document.getElementById('password');
    const passwordToogle = document.getElementById('toggle-password');
    passwordToogle.onclick = () => FormUtils.togglePassword('password');
    password.addEventListener('input', () => {
        if(!FormUtils.validatePassword(password.value)){
            password.classList.remove('border-green-500');
            password.classList.add('border-red-500');
        }else{
            password.classList.remove('border-red-500');
            password.classList.add('border-green-500');
        }
    });
}

export function setup() {
    setupPasswordValidation();
    const form = document.getElementById('signin-form');
    StorageUtils.loadInputValues(form);
    form.addEventListener("input", _ => {
        StorageUtils.saveInputValues(form);
    })
}
