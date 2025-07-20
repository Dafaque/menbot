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
import ChatUsersRemoveView from "./views/chat-users-remove.js";
import ChatRoleRemoveView from "./views/chat-role-remove.js";
import ChatRoleOptionsView from "./views/chat-role-options.js";
// Создаем приложение
new App();

const router = new Router();

// Регистрируем маршруты
router.register("/", IndexView);
router.register("/chats", ChatsView);
router.register("/chats/details", ChatDetailsView);
router.register("/chats/details/edit", EditChatView);
router.register("/chats/details/users", ChatUsersView);
router.register("/chats/details/users/remove", ChatUsersRemoveView);
router.register("/chats/details/roles", ChatRolesView);
router.register("/chats/details/roles/add", ChatRolesAddView);
router.register("/chats/details/roles/list", ChatRolesListView);
router.register("/chats/details/roles/list/options", ChatRoleOptionsView);
router.register("/chats/details/roles/list/options/remove", ChatRoleRemoveView);
router.register("/chats/details/roles/list/options/manage", ChatRoleUsersView);
window.app.setRouter(router);
router.navigate("/", window.app);
