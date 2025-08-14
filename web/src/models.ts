export interface Chat {
    id: string;
    tg_chat_name: string;
}

export interface Role {
    id: string;
    name: string;
    chat_id: string;
    tg_chat_name: string;
}

export interface User {
    id: string;
    tg_user_id: string;
    tg_user_name: string;
    tg_chat_name: string;
    chat_id: string;
}