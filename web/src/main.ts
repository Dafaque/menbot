import { App, Router } from "kibodo";
import IndexView from "./views/index";
import ChatsView from "./views/chats";
import EditChatView from "./views/chat-edit";
import ChatDetailsView from "./views/chat-details";
import ChatUsersView from "./views/chat-users";
import ChatRolesView from "./views/chat-roles";
import ChatRolesAddView from "./views/chat-roles-add";
import ChatRolesListView from "./views/chat-roles-list";
import ChatRoleUsersView from "./views/chat-role-users";
import ChatUsersRemoveView from "./views/chat-users-remove";
import ChatRoleRemoveView from "./views/chat-role-remove";
import ChatRoleOptionsView from "./views/chat-role-options";
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
