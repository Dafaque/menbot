import App from "./framework/app.js";
import MenuView from "./views/menu-view.js";
import SettingsView from "./views/settings-view.js";
import InfoView from "./views/info-view.js";
import ApiView from "./views/api-view.js";

// Создаем приложение
const app = new App();

// Регистрируем маршруты
app.registerRoute("/", MenuView);
app.registerRoute("/menu", MenuView);
app.registerRoute("/settings", SettingsView);
app.registerRoute("/info", InfoView);
app.registerRoute("/api", ApiView);

// Запускаем приложение
app.start("/");