const mover_para_resultados = () => {
    if(document.getElementById("resultados_container")) {
        document.getElementById("resultados_container").scrollIntoView({  behavior: 'smooth' });
    }
}











const main = () => {
    mover_para_resultados();
}
main();
