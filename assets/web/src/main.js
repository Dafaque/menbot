import App from "./framework/app.js";
import Router from "./framework/router.js";
import IndexView from "./views/index.js";
import ChatsView from "./views/chats.js";
import EditChatView from "./views/chat-edit.js";
import ChatDetailsView from "./views/chat-details.js";
import ChatUsersView from "./views/chat-users.js";
import ChatRolesView from "./views/chat-roles.js";
import ChatRolesAddView from "./views/chat-roles-add.js";
import ChatRolesListView from "./views/chat-roles-list.js";
import ChatRoleUsersView from "./views/chat-role-users.js";
// Создаем приложение
new App();

const router = new Router();

// Регистрируем маршруты
router.register("/", IndexView);
router.register("/chats", ChatsView);
router.register("/chats/details", ChatDetailsView);
router.register("/chats/details/edit", EditChatView);
router.register("/chats/details/users", ChatUsersView);
router.register("/chats/details/roles", ChatRolesView);
router.register("/chats/details/roles/add", ChatRolesAddView);
router.register("/chats/details/roles/list", ChatRolesListView);
router.register("/chats/details/roles/list/users", ChatRoleUsersView);

window.app.setRouter(router);
router.navigate("/", window.app);
