import { Menu } from "kibodo";
import ChatRolesAddView from "./chat-roles-add";
import ChatRolesListView from "./chat-roles-list";
import { Chat } from "../models";

export default class ChatRolesView extends Menu {
    chat: Chat;
    constructor(data: Chat) {
        super([
            {
                label: "Add Role",
                action: () => this.goToAddRole()
            },
            {
                label: "List Roles",
                action: () => this.goToListRoles()
            }
        ]);
        this.title = `Chat Roles of ${data.tg_chat_name}`;
        this.chat = data;
    }

    goToAddRole = () => {
        window.app.push(new ChatRolesAddView(this.chat));
    }

    goToListRoles = () => {
        window.app.push(new ChatRolesListView(this.chat));
    }
}
