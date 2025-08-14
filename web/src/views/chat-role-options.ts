import { Menu } from "kibodo";
import ChatRoleRemoveView from "./chat-role-remove";
import ChatRoleUsersView from "./chat-role-users";
import { Role } from "../models";

export default class ChatRoleOptionsView extends Menu {
    role: Role;
    constructor(data: Role) {
        super(
            [
                {
                    label: "Edit users",
                    action: () => this.gotToEditUsers()
                },
                {
                    label: "Remove role",
                    action: () => this.gotToRemoveRole()
                }
            ]

        );
        this.role = data;
        this.title = `Role ${this.role.name} options`;
    }

    gotToEditUsers = () => {
        window.app.push(new ChatRoleUsersView(this.role));
    }

    gotToRemoveRole = () => {
        window.app.push(new ChatRoleRemoveView(this.role)).then((removed:boolean) => {
            if (removed) {
                window.app.pop();
            }
        });
    }
}
