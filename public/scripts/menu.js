const hamburguerMenu = document.getElementById('hamburguer-menu');
const hamburguerClose = document.getElementById('hamburguer-menu-close');
const mobileMenu = document.getElementById('menu-container-mobile');
const searchMobileInput = document.getElementById('search-mobile-input');
hamburguerMenu.addEventListener('click', () => {
    mobileMenu.classList.remove('hidden');
    hamburguerMenu.classList.add('hidden');
    hamburguerClose.classList.remove('hidden');
});
hamburguerClose.addEventListener('click', () => {
    mobileMenu.classList.add('hidden');
    hamburguerMenu.classList.remove('hidden');
    hamburguerClose.classList.add('hidden');
});
