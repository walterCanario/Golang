// Archivo: static/js/app.js

document.addEventListener('alpine:init', () => {
    console.log('Alpine.js inicializado'); // Verifica que el evento se esté ejecutando
    Alpine.data('mainApp', () => ({
        menu: 'default',
        sidenavContent: [],

        init() {
            console.log('Componente mainApp inicializado'); // Verifica que el componente se esté inicializando
            this.fetchSideNav(this.menu);
        },

        fetchSideNav(menu) {
            console.log('Cargando menú lateral para:', menu); // Verifica que la función se esté ejecutando
            fetch(`/sidenav?menu=${menu}`)
                .then(response => {
                    if (!response.ok) {
                        throw new Error('Error en la respuesta del servidor');
                    }
                    return response.json();
                })
                .then(data => {
                    console.log('Datos recibidos:', data); // Verifica los datos recibidos
                    this.sidenavContent = data;
                })
                .catch(error => {
                    console.error('Error cargando el menú lateral:', error);
                    alert('Error cargando el menú lateral. Inténtalo de nuevo.');
                });
        }
    }));
});