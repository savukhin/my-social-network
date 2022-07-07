export class User {
    id: number = 0;
    name: string = "";
    isOnline: boolean = true;
    status: string = "";
    birthDate: string = "";
    city: string = "";

    constructor(id=0, name="", isOnline=true, status="", birthDate="", city="") {
        this.id = id;
        this.name = name;
        this.isOnline = isOnline;
        this.status = status;
        this.birthDate = birthDate;
        this.city = city;
    }
}