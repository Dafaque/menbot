import { Table } from "kibodo";
import ChatUsersRemoveView from "./chat-users-remove";
import { Chat } from "../models";

export default class ChatUsersView extends Table {
    chat: Chat;
    constructor(data: Chat) {
        super({
            url: `/api/chats/${data.id}/users`,
            attr: "UserListResponseJSONResponse"
        });
        this.title = `Chat Users of ${data.tg_chat_name}`;
        this.chat = data;
    }

    onSelected = (row: any) => {
        window.app.push(new ChatUsersRemoveView(row)).then(() => {
            this.fetchData();
        });
    }
}
