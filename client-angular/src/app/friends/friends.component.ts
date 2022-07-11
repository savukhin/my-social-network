import { Component, OnInit } from '@angular/core';
import { User } from 'src/models/user';

@Component({
  selector: 'app-friends',
  templateUrl: './friends.component.html',
  styleUrls: ['./friends.component.scss']
})
export class FriendsComponent implements OnInit {
  friends: User[] = [
    new User(1, "Kirill Klimonov"),
    new User(2, "Kirill Klimonov2", false)
  ];

  constructor() { }

  ngOnInit(): void {
  }

}
