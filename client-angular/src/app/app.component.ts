import { AfterViewInit, Component } from '@angular/core';
import { User } from 'src/models/user';
import { AuthService } from './services/backend-api/auth.service';

interface Answer {
  title: string
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
  providers: [AuthService]
})
export class AppComponent implements AfterViewInit {
  user?: User;
  title = 'client-angular';
  token = localStorage.getItem("token");

  constructor(private auth: AuthService) {}

  ngAfterViewInit(): void {
    if (this.token != null)
      this.auth.authWithToken(this.token).subscribe((user) => {
        console.log(user);
        
        if (user != null)
          this.user = user;
      })
  }
}
