import { Table } from "kibodo";
import ChatRoleOptionsView from "./chat-role-options";
import { Chat } from "../models";

export default class ChatRolesListView extends Table {
    chat: Chat;
    constructor(data: Chat) {
        super({
            url: `/api/chats/${data.id}/roles`,
            attr: "RoleListResponseJSONResponse"
        });
        this.title = `Chat ${data.tg_chat_name} Roles List`;
        this.chat = data;
    }

    onSelected(row: any) {
        window.app.push(new ChatRoleOptionsView(row)).then(() => {
            this.fetchData();
        });
    }

}