document.addEventListener("alpine:init", () => {
    Alpine.data("menuController", () => ({
        selectedMenu: "inicio", // Opción por defecto

        get sideNavComponent() {
            switch (this.selectedMenu) {
                case "comparativas":
                    return "sideNavComparativas";
                case "reportes":
                    return "sideNavReportes";
                case "georeferencia":
                    return "sideNavGeoreferencia";
                default:
                    return "sideNavDefault";
            }
        },

        changeMenu(menu) {
            this.selectedMenu = menu;
        }
    }));
});
