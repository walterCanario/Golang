function mostrarMenu(menu) {
    document.getElementById('menuReporte').style.display = menu === 'reporte' ? 'block' : 'none';
}

function mostrarFiltros() {
    // Lógica para mostrar filtros según el panel seleccionado
}

function generarGraficos() {
    const panel = document.getElementById('selectorPanel').value;
    const ano = document.getElementById('ano').value;

    // Aquí puedes hacer una solicitud al backend en Go
    fetch('/api/graficos', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ panel, ano }),
    })
    .then(response => response.json())
    .then(data => {
        // Manejar la respuesta y mostrar gráficos
        console.log(data);
    })
    .catch(error => console.error('Error:', error));
}
