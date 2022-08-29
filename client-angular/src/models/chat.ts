import { Message } from "./message";
import { User } from "./user";

export class ChatDTO {
    id: number = 0;
    title: string = "";
    is_personal: boolean = false
    participants: User[] = []
    messages: Message[] = [];
    photo_url: string | null = null
}

export class Chat {
    id: number = 0;
    title: string = "";
    is_personal: boolean = false
    participants: User[] = []
    messages: Message[] = [];
    last_message?: Message;
    photo_url: string | null = null

    static fromDTO(dto: ChatDTO) {
        let chat = new Chat()

        chat.id = dto.id
        chat.title = dto.title;
        dto.participants.forEach(val => {
            chat.participants[val.id] = val
        })
        chat.messages = dto.messages;
        chat.photo_url = dto.photo_url

        return chat
    }
}