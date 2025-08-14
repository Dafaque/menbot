import { Menu } from "kibodo";
import EditChatView from "./chat-edit";
import ChatRolesView from "./chat-roles";
import ChatUsersView from "./chat-users";
import { Chat } from "../models";

export default class ChatDetailsView extends Menu {
    chat: Chat;
    
    constructor(data: Chat) {
        super([
            {
                label: "Edit Chat",
                action: () => {
                    window.app.push(new EditChatView(this.chat));
                }
            },{
                label: "Roles",
                action: () => {
                    window.app.push(new ChatRolesView(this.chat));
                }
            },{
                label: "Users",
                action: () => {
                    window.app.push(new ChatUsersView(this.chat));
                }
            }
        ]);
        this.chat = data;
        this.title = `Chat Details of ${this.chat.tg_chat_name}`;
    }
}
