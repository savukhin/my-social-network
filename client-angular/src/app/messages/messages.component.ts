import { Component, OnInit } from '@angular/core';
import { Message } from 'src/models/message';
import { User } from 'src/models/user';

class Item {
  message: Message = new Message();
  author: User = new User();

  constructor(message:Message, author:User) {
    this.message = message;
    this.author = author;
  }
}

@Component({
  selector: 'app-messages',
  templateUrl: './messages.component.html',
  styleUrls: ['./messages.component.scss']
})
export class MessagesComponent implements OnInit {
  items: Item[] = [
    new Item(
      new Message("Hi! How are you doing?", new Date()),
      new User(1, "Kirill Klimonov", false),
    )
  ];

  constructor() { }

  ngOnInit(): void {
  }

}
