import { Post } from "./post";
import { User } from "./user";

export class UserPage {
    id: number = 0;
    name: string = "";
    username: string = "";
    is_online: boolean = true;
    status: string = "";
    birthDate: string = "";
    city: string = "";
    avatar_url: string | null = null;
    posts: Post[] = [];

    static FromUser(user: User) {
        const page = new UserPage()

        page.id = user.id;
        page.name = user.name;
        page.username = user.username;
        page.is_online = user.is_online;
        page.status = user.status;
        page.birthDate = user.birthDate;
        page.city = user.city;
        page.avatar_url = user.avatar_url;

        return page
    }
}