import { Component, OnInit } from '@angular/core';
import { User, UserCompressed } from 'src/models/user';
import { AuthService } from '../services/backend-api/auth.service';

@Component({
  selector: 'app-friends',
  templateUrl: './friends.component.html',
  styleUrls: ['./friends.component.scss']
})
export class FriendsComponent implements OnInit {
  friends: UserCompressed[] = [];

  constructor(private auth: AuthService) { }

  ngOnInit(): void {
    if (!this.auth.user)
      return

    this.auth.getFriends(this.auth.user.id).subscribe(response => {
      if (response != false)
        this.friends = response
    })
  }

}
