if (typeof usuario_acessou_via_mobile !== 'function') {
  function usuario_acessou_via_mobile() {
    return /Mobi|Android|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(
      navigator.userAgent,
    );
  }
}

if (typeof desabilitarScroll !== 'function') {
  function desabilitarScroll() {
    document.body.style.overflow = "hidden";
  }
}

if (typeof habilitarScroll !== 'function') {
  function habilitarScroll() {
    document.body.style.overflow = "auto";
  }
}

if (typeof moverParaTopo !== 'function') {
  function moverParaTopo() {
    window.scrollTo({
      top: 0,
      behavior: "smooth", // Adiciona animação suave
    });
  }
}
if(typeof montar_grafico_v2 === 'undefined') {
    function montar_grafico_v2 (idElemento, dados) {
        const ctx = document.getElementById(idElemento);
        if (!ctx) throw `${idElemento} não encontrado!`;
        else if (window.grafico_canva) window.grafico_canva.destroy();

        const labels = dados.y;
        const data = dados.x;
        const datasets = [
            { label: "Juros", data: data[0] },
            { label: "Acumulado", data: data[1] },
        ];
        window.grafico_canva = new Chart(ctx, {
            type: "line",
            data: { labels, datasets },
            options: {
                scales: {
                    y: { beginAtZero: true },
                },
                responsive: true, // Torna o gráfico responsivo
                maintainAspectRatio: false, // Permite alterar a altura
            },
        });
    };
}
if (typeof main_local !== "function"){
    function main_local() {
        const dados_tabela = window.dados_tabela;
        if (!dados_tabela) return;
        const meses = dados_tabela.map((item) => item.data_mes_ano);
        const juros = dados_tabela.map((item) => item.juros);
        const acumulado = dados_tabela.map((item) => item.acumulado);
        const grafico_container = document.getElementById("grafico_container");
        montar_grafico_v2("chartjs", { y: meses, x: [juros, acumulado] });

        document
            .getElementById("botao_ativar_grafico")
            .addEventListener("click", () => {
                grafico_container.classList.remove("hidden");
                desabilitarScroll();
                moverParaTopo();
            });
        document.getElementById("grafico_fechar").addEventListener("click", () => {
            grafico_container.classList.add("hidden");
            habilitarScroll();
        });
    };
}
main_local();
//document.addEventListener("DOMContentLoaded", () => main())
