import { Post } from "./post";
import { User, UserCompressed } from "./user";

export class UserPage {
    id: number = 0;
    name: string = "";
    username: string = "";
    is_online: boolean = true;
    status: string = "";
    birth_date: string = "";
    city: string = "";
    avatar_url: string | null = null;
    posts: Post[] = [];
    friends: UserCompressed[] = [];
    added_to_friends = false

    static FromUser(user: User) {
        const page = new UserPage()

        page.id = user.id;
        page.name = user.name;
        page.username = user.username;
        page.is_online = user.is_online;
        page.status = user.status;
        page.birth_date = user.birth_date;
        page.city = user.city;
        page.avatar_url = user.avatar_url;

        return page
    }
}