class Menu {
    constructor(items = []) {
        this.items = items;
        this.selectedIndex = 0;
        this.container = document.createElement("div");
        this.container.classList.add("menu");
        this.app = null;
    }

    setApp(app) {
        this.app = app;
    }

    addItem(text, action) {
        this.items.push({ text, action });
    }

    handleKey(key) {
        switch (key) {
            case "ArrowUp":
                this.selectedIndex = this.selectedIndex > 0 ? this.selectedIndex - 1 : this.items.length - 1;
                this.render();
                break;
            case "ArrowDown":
                this.selectedIndex = this.selectedIndex < this.items.length - 1 ? this.selectedIndex + 1 : 0;
                this.render();
                break;
            case "Enter":
                if (this.items.length > 0) {
                    const item = this.items[this.selectedIndex];
                    if (item.action) {
                        item.action();
                    }
                }
                break;
        }
    }

    render() {
        this.container.innerHTML = "";
        
        this.items.forEach((item, index) => {
            const menuItem = document.createElement("div");
            menuItem.classList.add("menu-item");
            menuItem.textContent = item.text;
            
            if (index === this.selectedIndex) {
                menuItem.classList.add("selected");
            }
            
            this.container.appendChild(menuItem);
        });
        
        return this.container;
    }
}

export default Menu; 