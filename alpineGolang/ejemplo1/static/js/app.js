document.addEventListener("alpine:init", () => {
    Alpine.data("mainApp", () => ({
        menu: "default",
        sidenavHtml: "",

        init() {
            this.fetchSideNav(this.menu);
        },

        fetchSideNav(menu) {
            fetch(`/sidenav?menu=${menu}`)
                .then(response => response.text())
                .then(html => {
                    this.sidenavHtml = html;
                })
                .catch(error => console.error("Error cargando el menú lateral:", error));
        }
    }));
});
