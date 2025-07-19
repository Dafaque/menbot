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
app.registerRoute("/menu/settings", SettingsView);
app.registerRoute("/menu/info", InfoView);
app.registerRoute("/menu/api", ApiView);

// Запускаем приложение
app.start("/");